use axum::{http::StatusCode, response::IntoResponse, Json};
use serde_json::json;

pub async fn timeout_handler() -> impl IntoResponse {
    (StatusCode::OK, Json(vec!["foo".to_owned(), "bar".to_owned()]))
}