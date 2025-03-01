use axum::{extract::State, http::StatusCode, response::IntoResponse, Json};

use crate::{models::{requests::create_request::CreateRequest, responses::{self, http_response}}, repositories::mysql_repository::MysqlRepository, services::mysql_service::MysqlService, states::app_state::AppState};

pub async fn create(State(
    state): State<AppState>,
    Json(create_request): Json<CreateRequest>
) -> impl IntoResponse {
    let http_response = match state.mysql_service.create(create_request).await {
        
    };
    
}