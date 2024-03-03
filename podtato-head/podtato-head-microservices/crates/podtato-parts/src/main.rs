#![deny(warnings)]

use std::convert::Infallible;
use std::env;
use std::net::SocketAddr;
use std::path::Path;
use std::sync::Arc;
use std::time::SystemTime;

use bytes::Bytes;
use http_body_util::Full;
use hyper::server::conn::http1;
use hyper::service::service_fn;
use hyper::{Method, Request, Response};
use hyper_util::rt::TokioIo;
use once_cell::sync::Lazy;
use opentelemetry::{
    metrics::{Counter, Histogram, MeterProvider as _, Unit},
    KeyValue,
};
use opentelemetry_sdk::metrics::SdkMeterProvider;
use prometheus::{Encoder, Registry, TextEncoder};
use tokio::net::TcpListener;

mod assets;
use assets::Assets;

static HANDLER_ALL: Lazy<[KeyValue; 1]> = Lazy::new(|| [KeyValue::new("handler", "all")]);

async fn serve_req(req: Request<impl hyper::body::Body>, state: Arc<AppState>) -> Result<Response<Full<Bytes>>, Infallible> {
    let request_start = SystemTime::now();
    state.http_counter.add(1, HANDLER_ALL.as_ref());
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
        state.http_req_histogram.record(
            request_start.elapsed().map_or(0.0, |d| d.as_secs_f64()),
            &[],
        );
        return Ok(Response::new(Full::new(Bytes::from(content_data))));
    } else if req.method() == &Method::GET && req.uri().path() == "/metrics" {
        let mut buffer = vec![];
        let encoder = TextEncoder::new();
        let metric_families = state.registry.gather();
        encoder.encode(&metric_families, &mut buffer).unwrap();
        state
            .http_body_gauge
            .record(buffer.len() as u64, HANDLER_ALL.as_ref());

        // return Ok(Response::builder()
        //     .status(200)
        //     .header(CONTENT_TYPE, encoder.format_type())
        //     .body(Body::from(buffer)));
        return Ok(Response::new(Full::new(Bytes::from(buffer))));
    } else {
        return Ok(Response::new(Full::new(Bytes::from("Hello World!"))));
    }
}

struct AppState {
    registry: Registry,
    http_counter: Counter<u64>,
    http_body_gauge: Histogram<u64>,
    http_req_histogram: Histogram<f64>,
}

#[tokio::main(flavor = "current_thread")]
pub async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let port = env::var("PODTATO_PORT").unwrap_or("9000".to_string());
    let port = port
        .parse::<u16>()
        .expect(format!("Failed to parse port {}", port).as_str());

    println!("going to serve on port {}", port);

    let registry = Registry::new();
    let exporter = opentelemetry_prometheus::exporter()
        .with_registry(registry.clone())
        .build()?;
    let provider = SdkMeterProvider::builder().with_reader(exporter).build();

    let meter = provider.meter("hyper-example");
    let state = Arc::new(AppState {
        registry,
        http_counter: meter
            .u64_counter("http_requests_total")
            .with_description("Total number of HTTP requests made.")
            .init(),
        http_body_gauge: meter
            .u64_histogram("example.http_response_size")
            .with_unit(Unit::new("By"))
            .with_description("The metrics HTTP response sizes in bytes.")
            .init(),
        http_req_histogram: meter
            .f64_histogram("example.http_request_duration")
            .with_unit(Unit::new("ms"))
            .with_description("The HTTP request latencies in milliseconds.")
            .init(),
    });

    let addr: SocketAddr = ([0, 0, 0, 0], port).into();

    // Bind to the port and listen for incoming TCP connections
    let listener = TcpListener::bind(addr).await?;
    loop {
        let (tcp, _) = listener.accept().await?;
        let io = TokioIo::new(tcp);
        let state = state.clone();
        tokio::task::spawn(async move {
            if let Err(err) = http1::Builder::new()
                //                .timer(TokioTimer)
                .serve_connection(io, service_fn(move |req| serve_req(req, state.clone())))
                .await
            {
                println!("Error serving connection: {:?}", err);
            }
        });
    }
}
