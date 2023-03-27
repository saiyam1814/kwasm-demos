use anyhow::{Result, Error};
use spin_sdk::{
    http::{Request, Response},
    config,
};
use serde::Serialize;
use route_recognizer::Params;
use regex::Regex;

mod ratings;

#[derive(Serialize)]
struct Rating {
    #[serde(skip_serializing_if = "Option::is_none")]
    stars: Option<i8>,
    #[serde(skip_serializing_if = "Option::is_none")]
    color: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    error: Option<String>
}

#[derive(Serialize)]
struct Review {
    reviewer: String,
    text: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    rating: Option<Rating>
}

#[derive(Serialize)]
struct Reviews {
    id: String,
    podname: String,
    clustername: String,
    reviews: Vec<Review>,
}

fn ratings_enabled() -> bool {
    match config::get("enable_ratings") {
        Ok(s) => {
            let re = Regex::new("(?i)true").unwrap();
            re.find(&s).is_some()
        },
        Err(_) => false
    }
}

fn rating(r: i8) -> Option<Rating> {
    let star_color = config::get("star_color").unwrap_or("black".to_string());

    if ratings_enabled() {
        if r == -1 {
            Some(Rating { error: Some("Ratings service is currently unavailable".to_string()), stars: None, color: None })
        } else {
            Some(Rating { stars: Some(r), color: Some(star_color.clone()), error: None })
        }
    } else {
        None
    }
}

fn reviews_response(product_id: &str, stars_review1: i8, stars_review2: i8) -> Reviews {
    let podname = config::get("hostname").unwrap_or("".to_string());
    let clustername = config::get("cluster_name").unwrap_or("".to_string());
    
    let rating1 = rating(stars_review1); 
    let rating2 = rating(stars_review2); 

    Reviews{
        id: product_id.to_string(),
        podname: podname,
        clustername: clustername,
        reviews: vec![
            Review{
                reviewer: "Reviewer1".to_string(),
                text: "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!".to_string(),
                rating: rating1,
            },
            Review{
                reviewer: "Reviewer2".to_string(),
                text: "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare.".to_string(),
                rating: rating2,
            },
        ],
    }
}

pub fn handler(req: Request, p: &Params) -> Result<Response> {
    let product_id = p.find("productId").unwrap().to_string();
    
    let mut stars_review1 = -1;
    let mut stars_review2 = -1;
    
    if ratings_enabled() {
        let ratings = ratings::get_ratings(&product_id, req.headers());
        if let Ok(rating) = ratings {
            stars_review1 = rating.get("Reviewer1");
            stars_review2 = rating.get("Reviewer2");
        }
    }

    let r = reviews_response(&product_id, stars_review1, stars_review2);
    let json = serde_json::to_string(&r).unwrap();

    http::Response::builder()
        .status(200)
        .header("foo", "bar")
        .body(Some(json.into()))
        .map_err(Error::from)
}