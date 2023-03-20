use anyhow::{Result, Error};
use serde::Serialize;
use route_recognizer::Params;
use spin_sdk::http::{Request, Response};

#[derive(Serialize)]
struct Health{
    status: String,
}

pub fn handler(_req: Request, _p: &Params) -> Result<Response> {
    let r = Health{status: "Reviews is healthy".to_string()};
    let json = serde_json::to_string(&r).unwrap();

    http::Response::builder()
        .status(200)
        .body(Some(json.into()))
        .map_err(Error::from)
}