using Microsoft.AspNetCore.Mvc;
using aspnetcoremvc.Database;
using aspnetcoremvc.Models;

namespace aspnetcoremvc.Controllers;

public class SingleQueryController : Controller
{
    [Route("db")]
    [Produces("application/json")]
    public Task<World> Index([FromServices] Db db)
    {
        return db.LoadSingleQueryRow();
    }
}
