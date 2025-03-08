using System.Data;
using net9.Models.Requests;
using net9.Models.Responses;
using net9.Repositories;

namespace net9.Services {
    public class Test1Service(IDbConnection db, ITest1Repository test1Repository) : ITest1Service
    {
        private readonly IDbConnection? _db = db;
        private readonly ITest1Repository _test1Repository = test1Repository
;

        public async Task<Response<Test1?>> Create(Test1CreateRequest test1CreateRequest)
        {
            using var transaction = _db.BeginTransaction();
            try
            {
                var test1 = new Test1 { Test = test1CreateRequest.Test, Id = 0};
                var lastInserteId = await _test1Repository.Create(_db, test1, transaction);
                test1.Id = (int)lastInserteId;
                transaction.Commit();
                return ResponseHelper.SetCreatedResponse<Test1?>(test1);
            } catch(Exception) {
                transaction.Rollback();
                return ResponseHelper.SetInternalServerErrorResponse<Test1?>();
            }
        }

        public async Task<Response<Test1?>> GetById(int id)
        {
            try
            {
                var test1 = await _test1Repository.GetById(_db, id);
                return ResponseHelper.SetOkResponse<Test1>(test1);
            }
            catch(Exception)
            {
                return ResponseHelper.SetInternalServerErrorResponse<Test1?>();
            }
            // return 
        }

        public async Task<Response<Test1?>> Update(Test1UpdateRequest test1UpdateRequest)
        {
            // using var transaction = _db.BeginTransaction();
            try
            {
                var test1 = new Test1 { Id = test1UpdateRequest.Id, Test = test1UpdateRequest.Test};
                var rowsAffected = await _test1Repository.Update(_db, test1);
                if (rowsAffected != 1)
                {
                    return ResponseHelper.SetInternalServerErrorResponse<Test1?>();
                }
                // transaction.Commit();
                return ResponseHelper.SetOkResponse<Test1>(test1);
            } catch(Exception)
            {
                // transaction.Rollback();
                return ResponseHelper.SetInternalServerErrorResponse<Test1?>();
            }
        }

        public async Task<Response<MessageResponse?>> Delete(Test1DeleteRequest test1DeleteRequest)
        {
            // using var transaction = _db.BeginTransaction();
            try{
                var rowsAffected = await _test1Repository.Delete(_db, test1DeleteRequest.Id);
                if (rowsAffected != 1) {
                    return ResponseHelper.SetInternalServerErrorResponse<MessageResponse?>();
                }
                var messageResponse = new MessageResponse { Message = "successfully delete test" };
                // transaction.Commit();
                return ResponseHelper.SetNoContentResponse<MessageResponse?>(messageResponse);
            }
            catch(Exception)
            {
                // transaction.Rollback();
                return ResponseHelper.SetInternalServerErrorResponse<MessageResponse?>();
            }
        }
    }
}