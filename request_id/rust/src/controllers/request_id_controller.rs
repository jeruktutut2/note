use std::sync::Arc;

use axum::{http::StatusCode, response::IntoResponse, Extension, Json};
use serde_json::{json, Value};

use crate::middlewares;

// pub trait RequestIdController {
//     async fn get_test(&self) -> impl IntoResponse;
// }

// pub struct RequestIdControllerImpl {

// }

// impl RequestIdControllerImpl {
//     pub fn new() -> RequestIdControllerImpl {
//         RequestIdControllerImpl{

//         }
//     }
// }

// impl RequestIdController for RequestIdControllerImpl {
//     async fn get_test(&self) -> impl IntoResponse {
//         let body: Value = json!({"foo": "bar"});
//         (StatusCode::OK, Json(body))
//     }
// }

pub async fn get_test(Extension(data): Extension<Arc<middlewares::model_middleware::ModelMiddleware>>) -> impl IntoResponse {
    let body: Value = json!({"foo": "bar", "requestId": data.request_id});
    (StatusCode::OK, Json(body))
}