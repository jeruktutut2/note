pub trait MysqlService {
    async fn create() 
}

pub struct MysqlServiceImpl {

}

impl MysqlServiceImpl {
    pub fn new() -> MysqlServiceImpl {
        return MysqlServiceImpl {

        };
    }
}

impl MysqlService for MysqlServiceImpl {
    
}