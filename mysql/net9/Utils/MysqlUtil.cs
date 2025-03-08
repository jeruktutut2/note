using System.Data;
using System.Linq.Expressions;
using Dapper;
using MySql.Data.MySqlClient;

namespace net9.Utils
{
//     public class DatabaseUtil() : IDisposable
//     {
//         // private readonly string _connectionString = connectionString;
// //         private MySqlConnection? _connection;
//         private IDbConnection? _connection;

//         public async Task<IDbConnection> Connect()
//         {
//             try 
//             {
//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: connecting to localhost:3309");
//                 _connection = new MySqlConnection("server=localhost;database=test1;user=root;password=12345;port=3309;Pooling=true;Min Pool Size=5;Max Pool Size=100;Connection Timeout=30;Connection Lifetime=60;");
//                 _connection.Open();
//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: connected to localhost:3309");

//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: pinging to localhost:3309");
//                 // if (!_connection.Ping()) {
//                 var ping = await Ping();
//                 if (ping) {
//                     Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: failed to ping to localhost:3309");
//                     Environment.Exit(1);
//                 }
//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: pinged to localhost:3309");
//                 return _connection;
//             }
//             catch(Exception e)
//             {
//                 Console.WriteLine($"error when connectiong to mysql localhost:3009: {e}");
//                 Environment.Exit(1);
//                 return null;
//             }

//         }

// //         public IDbConnection GetConnection() => _connection;

//         private async Task<bool> Ping()
//         {
//             try
//             {
//                 var result = await _connection.ExecuteScalarAsync<int>("SELECT 1;");
//                 return result == 1;
//             }
//             catch(Exception)
//             {
//                 return false;
//             }
//         }

//         public void Dispose()
//         {
//             try
//             {
//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: closing to localhost:3309");
//                 _connection.Close();
//                 _connection.Dispose();
//                 Console.WriteLine($"{DateTime.UtcNow.ToString("yyyy-MM-dd HH:mm:ss")} mysql: closed to localhost:3309");
//             }
//             catch(Exception e)
//             {
//                 Console.WriteLine($"error when closing mysql connection localhost:3309: {e}");
//             }
//         }
//     }
}