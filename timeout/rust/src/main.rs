use std::time::Duration;

use axum::{http::StatusCode, routing::get, Router};
use chrono::Local;
use controllers::timeout_controller::timeout_handler;
use tokio::{net::TcpListener, signal};
use tower::{layer, timeout::TimeoutLayer, ServiceBuilder};

mod utils;
mod repositories;
mod models;
mod controllers;

#[tokio::main]
async fn main() {
    // println!("Hello, world!");
    let app = Router::new()
        .route("/", get(timeout_handler))
        .layer(
            ServiceBuilder::new()
                .layer(TimeoutLayer::new(Duration::from_secs(3)))
                .layer(HandleError tower::layer::util::HandleErrorLayer::new(|_: Box<dyn std::error::Error + Send + Sync>| async {
                    (StatusCode::REQUEST_TIMEOUT, "Request Timeout")
                })),
        );
        // .layer(
        //     ServiceBuilder::new()
        //         .layer(TimeoutLayer::new(Duration::from_secs(5)))
        //         .layer(HandleErrorLayer::new(|_: Box<dyn std::error::Error + Send + Sync>| async {
        //             (StatusCode::REQUEST_TIMEOUT, "Request Timeout")
        //         })),
        // );
        // .layer(TimeoutLayer::new(Duration::from_secs(5)));
        // .layer(TimeoutLayer::new(Duration::from_secs(3)));
    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    println!("{} axum: connected to {} {}", Local::now().format("%Y-%m-%d %H:%M:%S%.3f").to_string(), "0.0.0.0", "8080");
    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

async fn shutdown_signal() {
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {
            println!("in ctrl c");
        },
        _ = terminate => {
            println!("in terminate");
        },
    }
}