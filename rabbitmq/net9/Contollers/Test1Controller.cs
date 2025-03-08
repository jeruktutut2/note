using Microsoft.AspNetCore.Mvc;

namespace net9.Controllers {

    [ApiController]
    [Route("api/v1/message")]
    public class Test1Controller: ControllerBase
    {

        public Test1Controller() {

        }

        [HttpGet("send-message")]
        public async Task<IActionResult> GetTest1() {

        }

        [HttpGet]
        public async Task<IActionResult> PostTest1() {

        }
    }
}