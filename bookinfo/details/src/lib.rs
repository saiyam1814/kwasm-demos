use anyhow::Result;
use spin_sdk::{
    http::{Request, Response},
    http_component,
};
use serde::Serialize;

#[derive(Serialize)]
struct Book {
    id: u32,
    title: String,
    author: String,
}


#[http_component]
fn handle_details(req: Request) -> Result<Response> {
    let book =  Book {
        id: 1,
        title: "Sample Book".to_string(),
        author: "John Doe".to_string(),
    };
    let book_json = serde_json::to_string(&book)?;

    Ok(http::Response::builder()
        .status(200)
        .header("Content-Type", "application/json")
        .body(Some(book_json.into()))?)
}