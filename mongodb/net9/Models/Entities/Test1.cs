using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace net9.Models.Entities;

public class Test1
{
    [BsonId]
    // [BsonRepresentation(BsonType.String)]
    public ObjectId Id { set; get; } = ObjectId.GenerateNewId();

    [BsonElement("test")]
    public string Test { set; get; } = string.Empty;
}