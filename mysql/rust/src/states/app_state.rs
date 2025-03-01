use std::sync::Arc;

use crate::{repositories::mysql_repository::MysqlRepositoryImpl, services::mysql_service::{MysqlService, MysqlServiceImpl}};

#[derive(Clone)]
pub struct AppState {
    pub mysql_service: Arc<MysqlServiceImpl>,
}