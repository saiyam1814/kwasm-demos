#![deny(warnings)]

use std::convert::Infallible;
use std::env;
use std::net::SocketAddr;
use std::path::Path;

use bytes::Bytes;
use http_body_util::Full;
use hyper::header::CONTENT_TYPE;
use hyper::server::conn::http1;
use hyper::service::service_fn;
use hyper::{body,Method, Request, Response};
use hyper_util::rt::TokioIo;
use tokio::net::TcpListener;

mod assets;
use assets::Assets;

async fn serve_req(req: Request<impl hyper::body::Body>) -> Result<Response<Full<Bytes>>, Infallible> {
    if req.method() == &Method::GET && req.uri().path().starts_with("/parts") {
        let path = Path::new(req.uri().path());
        let url = match path.file_name().unwrap().to_str() {
            Some("hat.svg") =>       "http://podtato-head-hat:9001/images/hat.svg",
            Some("left-leg.svg") =>  "http://podtato-head-left-leg:9002/images/left-leg.svg",
            Some("left-arm.svg") =>  "http://podtato-head-left-arm:9003/images/left-arm.svg",
            Some("right-leg.svg") => "http://podtato-head-right-leg:9004/images/right-leg.svg",
            Some("right-arm.svg") => "http://podtato-head-right-arm:9005/images/right-arm.svg",
            Some(&_) => todo!(),
            None => todo!(),
        };
        let res = reqwest::get(url).await.unwrap();
        let content_data = res.text().await.unwrap();
        let mut response = Response::new(Full::new(Bytes::from(content_data)));
        if url.ends_with(".svg"){
            response.headers_mut().insert(CONTENT_TYPE, "image/svg+xml".parse().unwrap());
        }
        return Ok(response)
} else if req.method() == &Method::GET && req.uri().path().starts_with("/assets/") {
        let path = req.uri().path().strip_prefix("/assets/").unwrap();
        let content = Assets::get(&path);
        return match content {
            Some(content) => {
                println!("FOUND: {}", path);
            let content_data = content.data.to_vec();
            let mut response = Response::new(Full::new(Bytes::from(content_data)));
            if path.ends_with(".svg"){
                response.headers_mut().insert(CONTENT_TYPE, "image/svg+xml".parse().unwrap());
            }
            Ok(response)
            }
            None => Ok(Response::new(Full::new(Bytes::from("Not Found!")))),
        }
    } else if req.method() == &Method::GET && req.uri().path() == "/metrics" {
        return Ok(Response::new(Full::new(Bytes::from("METRICS!"))));
    } else if req.method() == &Method::GET && req.uri().path() == "/" {
        // TODO: Handle file not found
        let content = Assets::get("html/podtato-home.html").unwrap();
        let content_data = String::from_utf8(content.data.to_vec()).unwrap().replace("{{ . }}", "0.1.0-WasmEdge");
        return Ok(Response::new(Full::new(Bytes::from(content_data))));
    } else {
        return Ok(Response::new(Full::new(Bytes::from("Hello World!"))));
    }
}

#[tokio::main(flavor = "current_thread")]
pub async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let port = env::var("PODTATO_PORT").unwrap_or("9000".to_string());
    let port = port
        .parse::<u16>()
        .expect(format!("Failed to parse port {}", port).as_str());

    println!("going to serve on port {}", port);

    let addr: SocketAddr = ([0, 0, 0, 0], port).into();

    // Bind to the port and listen for incoming TCP connections
    let listener = TcpListener::bind(addr).await?;
    loop {
        let (tcp, _) = listener.accept().await?;
        let io = TokioIo::new(tcp);

        tokio::task::spawn(async move {
            if let Err(err) = http1::Builder::new()
                //                .timer(TokioTimer)
                .serve_connection(io, service_fn(|req: Request<body::Incoming>|serve_req(req)))
                .await
            {
                println!("Error serving connection: {:?}", err);
            }
        });
    }
}
