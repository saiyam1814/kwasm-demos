use anyhow::{Error, Result};
use spin_sdk::config;
use serde::Deserialize;
use std::collections::HashMap;
use http::{header::HeaderMap, HeaderName};

#[derive(Deserialize)]
pub struct RatingResponse {
    ratings: HashMap<String, i8>,
}

pub fn get_rating(resp: &RatingResponse, k: String) -> i8 {
    resp.ratings.get(&k).and_then(|r| Some(*r)).unwrap_or(-1)
}

pub fn get_ratings(product_id: &String, headers: &HeaderMap) -> Result<RatingResponse> {
    let url = format!("{}/{}", ratings_service(), product_id);
    let mut req = http::Request::builder()
            .method("GET")
            .uri(url);
    if let Some(req_headers) = req.headers_mut() {
        for (h, v) in headers.iter().filter(filter_headers) {
            req_headers.insert(h, v.into());
        }
    }

    let res = spin_sdk::http::send( req.body(None)?)?;
    
    let body = res.body().as_ref().unwrap().as_ref();
    serde_json::from_slice(body).map_err(Error::from)
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

fn filter_headers<T>(h: &(&HeaderName, &T)) -> bool {
    match h.0.as_str() {
        "x-request-id" |
        "x-ot-span-context" |
        "x-datadog-trace-id" |
        "x-datadog-parent-id" |
        "x-datadog-sampling-priority" |
        "traceparent" |
        "tracestate" |
        "x-cloud-trace-context" |
        "grpc-trace-bin" |
        "x-b3-traceid" |
        "x-b3-spanid" |
        "x-b3-parentspanid" |
        "x-b3-sampled" |
        "x-b3-flags" |
        "sw8" |
        "end-user" |
        "user-agent" |
        "cookie" |
        "authorization" |
        "jwt" => true,
        _ => false
    }
}