use anyhow::{Error, Result};
use http::{header::HeaderMap, HeaderName};
use serde::Deserialize;
use std::collections::HashMap;
use std::env;

#[derive(Deserialize)]
pub struct RatingResponse {
    ratings: HashMap<String, i8>,
}

impl RatingResponse {
    pub fn get(&self, id: &str) -> i8 {
        self.ratings.get(id).copied().unwrap_or(-1)
    }
}

pub fn get_ratings(product_id: &str, headers: &HeaderMap) -> Result<RatingResponse> {
    let url = format!("{}/{}", ratings_service_url(), product_id);
    println!("getting ratings from {url}");
    let mut req = http::Request::builder().method("GET").uri(url);
    if let Some(req_headers) = req.headers_mut() {
        for (h, v) in headers.iter().filter(filter_headers) {
            req_headers.insert(h, v.into());
        }
    }

    let res = spin_sdk::http::send(req.body(None)?)?;

    let body = res.body().as_ref().unwrap().as_ref();
    serde_json::from_slice(body).map_err(Error::from)
}

fn ratings_service_url() -> String {
    let services_domain = match env::var("SERVICES_DOMAIN") {
        Ok(s) => format!(".{s}"),
        Err(_) => "".to_string(),
    };
    let ratings_hostname = env::var("RATINGS_HOSTNAME").unwrap_or("ratings".to_string());
    let ratings_port = env::var("RATINGS_SERVICE_PORT").unwrap_or("9080".to_string());
    format!("http://{ratings_hostname}{services_domain}:{ratings_port}/ratings")
}

fn filter_headers<T>(h: &(&HeaderName, &T)) -> bool {
    matches!(
        h.0.as_str(),
        "x-request-id"
            | "x-ot-span-context"
            | "x-datadog-trace-id"
            | "x-datadog-parent-id"
            | "x-datadog-sampling-priority"
            | "traceparent"
            | "tracestate"
            | "x-cloud-trace-context"
            | "grpc-trace-bin"
            | "x-b3-traceid"
            | "x-b3-spanid"
            | "x-b3-parentspanid"
            | "x-b3-sampled"
            | "x-b3-flags"
            | "sw8"
            | "end-user"
            | "user-agent"
            | "cookie"
            | "authorization"
            | "jwt"
    )
}
