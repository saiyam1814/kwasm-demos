use anyhow::{Error, Result};
use route_recognizer::{Params, Router};
use spin_sdk::{
    http::{Request, Response},
    http_component,
};

mod health;
mod reviews;

fn error_handler(_req: Request, _p: &Params) -> Result<Response> {
    http::Response::builder()
        .status(404)
        .body(None)
        .map_err(Error::from)
}

/// A simple Spin HTTP component.
#[http_component]
fn reviews(req: Request) -> Result<Response> {
    let mut router: Router<fn(Request, &Params) -> Result<Response>> = Router::new();

    router.add("/health", crate::health::handler);
    router.add("/reviews/:productId", crate::reviews::handler);
    router.add("/", error_handler);
    router.add("/*", error_handler);

    println!("{:?}", req);

    let m = router.recognize(req.uri().path()).map_err(Error::msg)?;
    let handler = m.handler();
    handler(req, m.params())
}
