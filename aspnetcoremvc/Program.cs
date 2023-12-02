using aspnetcoremvc.Database;
using Microsoft.EntityFrameworkCore;

var connectionString = "server=localhost;user=root;database=hello_world";
var serverVersion = new MySqlServerVersion(new Version(8, 0, 29));

var builder = WebApplication.CreateBuilder(args);

builder.Logging.ClearProviders();

// Add services to the container.
builder.Services.AddControllersWithViews();
builder.Services.AddTransient<Db>();
builder.Services.AddDbContextPool<ApplicationDbContext>(
    options => options.UseMySql(connectionString, serverVersion));

var app = builder.Build();

// app.UseHttpsRedirection();
// app.UseStaticFiles();

app.UseRouting();

// app.UseAuthorization();

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.Lifetime.ApplicationStarted.Register(() => Console.WriteLine("Application started. Press Ctrl+C to shut down."));
app.Lifetime.ApplicationStopping.Register(() => Console.WriteLine("Application is shutting down..."));

app.Run();
