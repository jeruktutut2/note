use std::sync::Arc;

use crate::{models::{entities::test1::{self, Test1}, requests::{create_request::CreateRequest, delete_request::DeleteRequest, update_request::UpdateRequest}, responses::{create_response::CreateResponse, get_by_id_response::GetByIdResponse, http_response::{ErrorType, HttpResponse, MessageResponse}, update_response::UpdateResponse}}, repositories::mysql_repository::{self, MysqlRepository, MysqlRepositoryImpl}, utils::mysql_util::{MysqlUtil, MysqlUtilImpl}};

pub trait MysqlService {
    async fn create(&self, create_request: CreateRequest) -> HttpResponse<CreateResponse>;
    async fn get_by_id(&self, id: i32) -> HttpResponse<GetByIdResponse>;
    async fn update(&self, update_request: UpdateRequest) -> HttpResponse<UpdateResponse>;
    async fn delete(&self, delete_request: DeleteRequest) -> HttpResponse<()>;
}

pub struct MysqlServiceImpl {
    mysql_util: Arc<MysqlUtilImpl>,
    mysql_repository: Arc<MysqlRepositoryImpl>
}

impl MysqlServiceImpl {
    pub fn new(mysql_util: Arc<MysqlUtilImpl>, mysql_repository: Arc<MysqlRepositoryImpl>) -> MysqlServiceImpl {
        return MysqlServiceImpl {
            mysql_util,
            mysql_repository
        };
    }
}

impl MysqlService for MysqlServiceImpl {
    async fn create(&self, create_request: CreateRequest) -> HttpResponse<CreateResponse> {
        let mut tx = match self.mysql_util.begin().await {
            Ok(tx) => tx,
            Err(_) => {
                // return GenericResponse::set_internal_server_error_http_response()
                return HttpResponse::set_internal_server_error_http_response();
            }
        };

        let mut test1 = Test1 { id: 0, test: create_request.test };
        let (rows_affected, last_inserted_id) = match self.mysql_repository.create(&mut tx, &test1).await {
            Ok((rows_affected, last_inserted_id)) => (rows_affected, last_inserted_id),
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response()
            }
        };
        if rows_affected != 1 {
            return HttpResponse::set_internal_server_error_http_response()
        }
        test1.id = last_inserted_id as i32;
        
        // return GenericResponse::set_internal_server_error_http_response()
        // return HttpResponse::set_internal_server_error_http_response();
        let create_response = CreateResponse{id: test1.id, test: test1.test};
        return HttpResponse::set_created_http_response(create_response);
    }

    async fn get_by_id(&self, id: i32) -> HttpResponse<GetByIdResponse> {
        let test1 = match self.mysql_repository.get_by_id(self.mysql_util.get_pool().await, id).await {
            Ok(test1) => test1,
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response();
            }
        };
        let get_by_id_response = GetByIdResponse{id: test1.id, test: test1.test};
        return HttpResponse::set_data_http_response(200, get_by_id_response);
    }
    async fn update(&self, update_request: UpdateRequest) -> HttpResponse<UpdateResponse> {
        let mut tx = match self.mysql_util.begin().await {
            Ok(tx) => tx,
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response();
            }
        };
        let test1 = Test1{id: update_request.id, test: update_request.test};
        let rows_affected = match self.mysql_repository.update(&mut tx, &test1).await {
            Ok(rows_affected) => rows_affected,
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response();
            }
        };
        if rows_affected != 1 {
            return HttpResponse::set_internal_server_error_http_response();
        }
        let update_response = UpdateResponse{id: test1.id, test: test1.test};
        HttpResponse::set_data_http_response(200, update_response)
    }
    async fn delete(&self, delete_request: DeleteRequest) -> HttpResponse<()> {
        let mut tx = match self.mysql_util.begin().await {
            Ok(tx) => tx,
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response();
            }
        };
        let rows_affected = match self.mysql_repository.delete(&mut tx, delete_request.id).await {
            Ok(rows_affected) => rows_affected,
            Err(_) => {
                return HttpResponse::set_internal_server_error_http_response();
            }
        };
        if rows_affected != 1 {
            return HttpResponse::set_internal_server_error_http_response()
        }
        HttpResponse::set_internal_server_error_http_response()
    }
}