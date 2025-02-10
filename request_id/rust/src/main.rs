use std::time::Duration;

use tokio::net::TcpListener;
use tokio::signal;
use axum::{routing::get, Router};
use tokio::time::sleep;

#[tokio::main]
async fn main() {
    // println!("Hello, world!");
    let app = Router::new()
        .route("/slow", get(|| sleep(Duration::from_secs(5))));
    // Create a `TcpListener` using tokio.
    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();

    // Run the server with graceful shutdown
    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

async fn shutdown_signal() {
    // println!("1");
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    // println!("2");
    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    // println!("3");
    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    // println!("4");
    tokio::select! {
        _ = ctrl_c => {
            println!("in ctrl c");
        },
        _ = terminate => {
            println!("in terminate");
        },
    }
    // println!("5");
}