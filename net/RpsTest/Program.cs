using RpsTest;

var arrays = AppData.Init().Arrays;

var builder = WebApplication.CreateSlimBuilder(args);
var app = builder.Build();
app.MapGet("/", () => RunCode.MethodRun(arrays));

app.Run();