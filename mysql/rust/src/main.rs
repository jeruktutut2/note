use std::sync::Arc;

use axum::Router;
use chrono::Local;
use tokio::{net::TcpListener, signal};

mod utils;
mod services;
mod models;

#[tokio::main]
async fn main() {
    // println!("Hello, world!");

    let mysql_util = Arc::new(utils::mysql_util::MysqlUtilImpl::new().await);
    println!("mysql_util: {:?}", mysql_util);

    let app = Router::new();
    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    
    println!("{} axum: connecting to {} {}", Local::now().format("%Y-%m-%d %H:%M:%S%.3f").to_string(), "0.0.0.0", "8080");
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