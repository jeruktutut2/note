using Microsoft.AspNetCore.Mvc;

namespace net9.Controllers;

[ApiController]
[Route("pdf")]
public class FileController: ControllerBase
{

    public FileController()
    {

    }

    [HttpGet]
    public IActionResult GeneratePdf()
    {
        // var pdfBytes = _pdfService.GeneratePdf();
        // return File(pdfBytes, "application/pdf", "generated.pdf");
    }

}