using Microsoft.AspNetCore.Mvc;
using aspnetcoremvc.Database;
using aspnetcoremvc.Models;

namespace aspnetcoremvc.Controllers;

public class MultipleQueriesController : Controller
{
    [Route("queries")]
    [Produces("application/json")]
    public Task<World[]> Index([FromServices] Db db, [FromQuery(Name = "n")] int n = 1)
    {
        return db.LoadMultipleQueriesRows(n);
    }
}
