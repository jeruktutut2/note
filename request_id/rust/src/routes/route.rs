use std::sync::Arc;

use axum::{middleware, routing::get, Router};

use crate::{controllers::request_id_controller, middlewares};

// use crate::controllers::{self, request_id_controller::{self}};

// #[axum::debug_handler]
pub fn set_test_route() -> Router {
    // let request_id_controller = Arc::new(controllers::request_id_controller::RequestIdControllerImpl::new());
    // let request_id_controller = Arc::new(controllers::request_id_controller::RequestIdControllerImpl::new());
    // Router::new()
    //     .route("/test", get(move || async move  {
    //         let request_id_controller = Arc::clone(&request_id_controller);
    //         request_id_controller.get_test().await
    //     }))
    Router::new()
        .route("/test", get(request_id_controller::get_test))
        .layer(middleware::from_fn(middlewares::request_id_middleware::set_request_id))
}