using Microsoft.AspNetCore.Mvc;

namespace net9.Controllers;

[ApiController]
[Route("api/v1/stacktrace")]
public class StacktraceController: ControllerBase
{
    public StacktraceController()
    {

    }

    [HttpGet]
    public async Task<IActionResult> CheckPanic()
    {
        return Ok();
    }
}