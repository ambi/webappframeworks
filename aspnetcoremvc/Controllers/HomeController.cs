using Microsoft.AspNetCore.Mvc;

namespace aspnetcoremvc.Controllers;

public class HomeController : Controller
{
    [HttpGet("plaintext")]
    public string Plaintext()
    {
        return "Hello, World!";
    }

    [HttpGet("json")]
    [Produces("application/json")]
    public object Json()
    {
        return new { message = "Hello, World!" };
    }
}
