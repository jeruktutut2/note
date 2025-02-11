use std::sync::Arc;

use axum::{extract::Request, middleware::Next, response::Response};
use uuid::Uuid;

use crate::middlewares::model_middleware;

pub async fn set_request_id(
    mut request: Request,
    next: Next,
) -> Response {
    // do something with `request`...
    println!("request middleware");
    // let uuidv4 = Uuid::new_v4();
    // println!("uuidv4{}", uuidv4);
    let data = model_middleware::ModelMiddleware {
        request_id: Uuid::new_v4().to_string(),
    };
    request.extensions_mut().insert(Arc::new(data));
    let response = next.run(request).await;
    println!("response middleware");
    // do something with `response`...

    response
}