using Microsoft.AspNetCore.Mvc;

namespace net9.Controllers;

[ApiController]
[Route("api/v1/pdf")]
public class PdfController: ControllerBase
{
    public PdfController()
    {

    }

    [HttpGet("{id}")]
    public async Task<IActionResult> GetTest1([FromRoute] int id)
    {
        var response = await _test1Service.GetById(id);
        return StatusCode(response.HttpStatuscode, response.BodyResponse);
    }

    [HttpPost]
    public async Task<IActionResult> PostTest1([FromBody] Test1CreateRequest test1CreateRequest)
    {
        var response = await _test1Service.Create(test1CreateRequest);
        return StatusCode(response.HttpStatuscode, response.BodyResponse);
    }
}