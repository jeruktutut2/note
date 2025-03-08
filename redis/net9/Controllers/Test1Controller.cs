using Microsoft.AspNetCore.Mvc;

namespace net9.Controllers {

    [ApiController]
    [Route("api/v1/redis")]
    public class Test1Controller: ControllerBase
    {

        public Test1Controller() {

        }

        [HttpGet("{id}")]
        public async Task<IActionResult> GetTest1() {

        }

        [HttpPost]
        public async Task<IActionResult> PostTest1() {

        }
    }
}