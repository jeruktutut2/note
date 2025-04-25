using Microsoft.AspNetCore.Mvc;
using net9.Services;

namespace net9.Controllers;

[ApiController]
// [Route("api/v1/pdf")]
[Route("pdf")]
public class PdfController: ControllerBase
{
    private IPdfService _pdfService;
    public PdfController(IPdfService pdfService)
    {
        _pdfService = pdfService;
    }

    // [HttpGet("{id}")]
    // public async Task<IActionResult> GetTest1([FromRoute] int id)
    // {
    //     var response = await _test1Service.GetById(id);
    //     return StatusCode(response.HttpStatuscode, response.BodyResponse);
    // }

    // [HttpPost]
    // public async Task<IActionResult> PostTest1([FromBody] Test1CreateRequest test1CreateRequest)
    // {
    //     var response = await _test1Service.Create(test1CreateRequest);
    //     return StatusCode(response.HttpStatuscode, response.BodyResponse);
    // }
    [HttpGet]
    public IActionResult GeneratePdf()
    {
        var pdfBytes = _pdfService.GeneratePdf();
        return File(pdfBytes, "application/pdf", "generated.pdf");
    }
}