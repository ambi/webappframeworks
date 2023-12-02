using Microsoft.AspNetCore.Mvc;
using aspnetcoremvc.Database;
using aspnetcoremvc.Models;

namespace aspnetcoremvc.Controllers;

public class UpdatesController : Controller
{
    [Route("updates")]
    [Produces("application/json")]
    public Task<World[]> Index([FromServices] Db db, [FromQuery(Name = "n")] int n = 1)
    {
        return db.LoadMultipleUpdatesRows(n);
    }
}
