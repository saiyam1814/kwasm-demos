use anyhow::{Error, Result};
use spin_sdk::config;
use serde::Deserialize;
use std::collections::HashMap;

#[derive(Deserialize)]
pub struct RatingResponse {
    ratings: HashMap<String, i8>,
}

pub fn get_rating(resp: &RatingResponse, k: String) -> i8 {
    resp.ratings.get(&k).and_then(|r| Some(*r)).unwrap_or(-1)
}

pub fn get_ratings(product_id: &String) -> Result<RatingResponse> {
    let url = format!("{}/{}", ratings_service(), product_id);
    let mut res = spin_sdk::http::send(
        http::Request::builder()
            .method("GET")
            .uri(url)
            .body(None)?,
    )?;

    res.headers_mut()
        .insert(http::header::SERVER, "spin/0.1.0".try_into()?);

    serde_json::from_slice(res.body().as_ref().unwrap().as_ref()).map_err(Error::from)
}

fn ratings_service() -> String {
    let services_domain = match config::get("services_domain") {
        Ok(s) => format!(".{}", &s),
        Err(_) => "".to_string()
    };
    let ratings_hostname = config::get("ratings_hostname").unwrap_or("ratings".to_string());
    let ratings_port = config::get("ratings_service_port").unwrap_or("9080".to_string());
    format!("http://{}{}:{}/ratings", ratings_hostname, services_domain, ratings_port)
}
