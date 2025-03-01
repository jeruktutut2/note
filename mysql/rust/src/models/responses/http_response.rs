use std::i32;

use serde::{Deserialize, Serialize};
use sqlx::FromRow;

use std::collections::HashMap;


#[derive(Debug, Deserialize, Serialize)]
pub struct MessageResponse {
    #[serde(rename = "message")]
    pub message: String
}

#[derive(Debug)]
pub enum ErrorType {
    Message(MessageResponse),
    Map(HashMap<String, String>),
}

#[derive(Debug, Deserialize, Serialize, Default)]
pub struct Response<T> {
    #[serde(rename = "data")]
    pub data: Option<T>,
    #[serde(rename = "errors")]
    pub errors: Option<HashMap<String, String>>,
}

#[derive(FromRow, Debug, Default, Serialize, Deserialize)]
pub struct HttpResponse<T> {
    pub http_status_code: i32,
    pub response: Response<T>
}

impl <T>HttpResponse<T> {
    pub fn set_data_http_response(http_status_code: i32, data: T) -> Self {
        Self {
            http_status_code,
            response: Response { 
                data: Some(data), 
                errors: None 
            }
        }
    }

    pub fn set_created_http_response(data: T) -> Self {
        Self {
            http_status_code: 201,
            response: Response { 
                data: Some(data), 
                errors: None 
            }
        }
    }

//     // pub fn set_message_http_response(http_status_code: i32, data: T) -> Self {
//     //     Self {
//     //         http_status_code,
//     //         response: Response { 
//     //             data: Some(data), 
//     //             errors: None 
//     //         }
//     //     }
//     // }
    
    pub fn set_bad_request_http_response(message: String) -> Self {
        // let error_message = HashMap::new()
        let mut error_message = HashMap::new();
        error_message.insert("message".to_string(), message);
        Self {
            http_status_code: 400,
            response: Response { 
                data: None, 
                // errors: Some(ErrorType::Message(MessageResponse { message }))
                errors: Some(error_message)
            }
        }
    }

    pub fn set_internal_server_error_http_response() -> Self {
        let mut error_message = HashMap::new();
        error_message.insert("message".to_string(), "internal server error".to_string());
        HttpResponse {
            http_status_code: 500,
            response: Response { 
                data: None, 
                // errors: Some(ErrorType::Message(MessageResponse { message: "internal server errir".to_string() }))
                errors: Some(error_message)
            }
        }
    }
}

// #[derive(Debug)]
// pub enum GenericResponse {
//     BadRequest(HttpResponse<(), MessageResponse>),
//     ValidationError(HttpResponse<(), HashMap<String, String>>),
//     InternalServerError(HttpResponse<(), MessageResponse>),
// }

// impl GenericResponse {
//     pub fn set_bad_request_http_response(message: String) -> Self {
//         GenericResponse::BadRequest(HttpResponse{
//             http_status_code: 400,
//             response: Response { 
//                 data: None, 
//                 errors: Some(MessageResponse { message }) 
//             }
//         }) 
//         // {
//         //     http_status_code: 400,
//         //     response: Response { 
//         //         data: None, 
//         //         errors: Some(MessageResponse { message }),
//         //     }
//         // }
//     }

//     pub fn set_validation_error_http_response(errors: HashMap<String, String>) -> Self {
//         GenericResponse::ValidationError(HttpResponse { 
//             http_status_code: 400, 
//             response: Response { 
//                 data: None, 
//                 errors: Some(errors) 
//             } 
//         }) 
//         // {
//         //     http_status_code: 400,
//         //     response: Response { 
//         //         data: None, 
//         //         errors: Some(errors),
//         //     }
//         // }
//     }

//     pub fn set_internal_server_error_http_response() -> Self {
//         GenericResponse::InternalServerError(HttpResponse { 
//             http_status_code: 500, 
//             response: Response { 
//                 data: None, 
//                 errors: Some(MessageResponse { message: "internal server error".to_string() }) 
//             } 
//         }) 
//         // {
//         //     http_status_code: 400,
//         //     response: Response { 
//         //         data: None, 
//         //         errors: Some(MessageResponse { message: "internal server error".to_string() }),
//         //     }
//         // }
//     }
// }

// impl <T> HttpResponse<T, HashMap<String, String>> {
//     pub fn set_validation_error_http_response(errors: HashMap<String, String>) -> Self {
//         Self {
//             http_status_code: 400,
//             response: Response { 
//                 data: None, 
//                 errors: Some(errors),
//             }
//         }
//     }
// }