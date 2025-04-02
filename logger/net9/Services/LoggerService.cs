using Microsoft.Extensions.Logging;

namespace net9.Services;

public class LoggerService: ILoggerService
{
    private readonly ILogger<LoggerService> _logger;

    public LoggerService(ILogger<LoggerService> logger)
    {
        _logger = logger;
    }

    public string CheckLogger()
    {
        // var logData = new Dictionary<string, object>
        // {
        //     ["UserId"] = "id",
        //     ["RequestId"] = Guid.NewGuid().ToString()
        // };
        // _logger.LogDebug("Ini log Debug");
        // _logger.LogInformation("Ini log Information");
        // _logger.LogWarning("Ini log Warning");
        // _logger.LogError("Ini log Error");
        _logger.LogCritical("Ini log Critical {T}", "test1");
        return "ok";
    }
}