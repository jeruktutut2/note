use serde::{Deserialize, Serialize};
use sqlx::FromRow;

pub struct Response {
    pub data:
}

#[derive(FromRow, Debug, Default, Serialize, Deserialize)]
pub struct HttpResponse {
    pub http_status_code: i32;
    
}