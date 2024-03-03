#![deny(warnings)]

use std::convert::Infallible;
use std::env;
use std::net::SocketAddr;
use std::path::Path;

use bytes::Bytes;
use http_body_util::Full;
use hyper::server::conn::http1;
use hyper::service::service_fn;
use hyper::{body,Method, Request, Response};
use hyper_util::rt::TokioIo;
use tokio::net::TcpListener;

mod assets;
use assets::Assets;

async fn serve_req(req: Request<impl hyper::body::Body>) -> Result<Response<Full<Bytes>>, Infallible> {
    if req.method() == &Method::GET && req.uri().path().starts_with("/images") {
        //TODO: Make desired art configurable
        let desired_part_number = "01";
        let path = Path::new(req.uri().path());
        let filestem = match path.file_stem() {
            Some(filestem) => filestem.to_str().unwrap(),
            None => "",
        };

        let path = format!(
            "images/{}/{}-{}.svg",
            filestem, filestem, desired_part_number
        );
        // TODO: Handle file not found
        let content = Assets::get(&path).unwrap();
        let content_data = content.data.to_vec();
        return Ok(Response::new(Full::new(Bytes::from(content_data))));
    } else if req.method() == &Method::GET && req.uri().path() == "/metrics" {
        return Ok(Response::new(Full::new(Bytes::from("METRICS!"))));
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
