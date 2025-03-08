using System.Text.Json.Serialization;

public class Test1
{
    [JsonPropertyName("id")]
    public int Id { set; get; }
    [JsonPropertyName("test")]
    public required string Test { set; get; }
}