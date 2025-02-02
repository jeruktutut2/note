package repositories

import (
	"context"
	modelentities "note-golang-mongodb/models/entitites"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository interface {
	InsertOne(db *mongo.Database, ctx context.Context, user modelentities.User) (err error)
	InsertMany(db *mongo.Database, ctx context.Context, users []modelentities.User) (err error)
	FindOne(db *mongo.Database, ctx context.Context, email string) (user modelentities.User, err error)
	Find(db *mongo.Database, ctx context.Context, email string) (users []modelentities.User, err error)
	UpdateOne(db *mongo.Database, ctx context.Context, user modelentities.User) (rowsAffected int64, err error)
	UpdateById(db *mongo.Database, ctx context.Context, id string) (rowsAffected int64, err error)
	DeleteOne(db *mongo.Database, ctx context.Context, id string) (deletedCount int64, err error)
	DeleteMany(db *mongo.Database, ctx context.Context) (deletedCount int64, err error)
}

type MongodbRepositoryImplementation struct {
}

func NewMongodbRepository() MongodbRepository {
	return &MongodbRepositoryImplementation{}
}

func (repository *MongodbRepositoryImplementation) InsertOne(db *mongo.Database, ctx context.Context, user modelentities.User) (err error) {
	_, err = db.Collection("users").InsertOne(ctx, user, nil)
	return
}

func (repository *MongodbRepositoryImplementation) InsertMany(db *mongo.Database, ctx context.Context, users []modelentities.User) (err error) {
	var documents []interface{}
	for _, user := range users {
		documents = append(documents, user)
	}
	_, err = db.Collection("users").InsertMany(ctx, documents, nil)
	return
}

func (repository *MongodbRepositoryImplementation) FindOne(db *mongo.Database, ctx context.Context, email string) (user modelentities.User, err error) {
	err = db.Collection("users").FindOne(ctx, bson.M{"email": email}, nil).Decode(&user)
	return
}

func (repository *MongodbRepositoryImplementation) Find(db *mongo.Database, ctx context.Context, email string) (users []modelentities.User, err error) {
	cursor, err := db.Collection("users").Find(ctx, bson.M{"email": email}, nil)
	if err != nil {
		return
	}
	defer func() {
		errCursor := cursor.Close(ctx)
		if errCursor != nil {
			users = []modelentities.User{}
			err = errCursor
		}
	}()

	err = cursor.All(ctx, &users)
	if err != nil {
		users = []modelentities.User{}
		return
	}

	if cursor.Err() != nil {
		users = []modelentities.User{}
		err = cursor.Err()
		return
	}
	return
}

func (repository *MongodbRepositoryImplementation) UpdateOne(db *mongo.Database, ctx context.Context, user modelentities.User) (rowsAffected int64, err error) {
	result, err := db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"email": user.Email, "password": user.Password}}, nil)
	if err != nil {
		return
	}
	rowsAffected = result.ModifiedCount
	return
}

func (repository *MongodbRepositoryImplementation) UpdateById(db *mongo.Database, ctx context.Context, id string) (rowsAffected int64, err error) {
	result, err := db.Collection("users").UpdateByID(ctx, id, bson.M{"$set": bson.M{"email": "email1@email.com", "password": "password@A2"}}, nil)
	if err != nil {
		return
	}
	rowsAffected = result.ModifiedCount
	return
}

func (repository *MongodbRepositoryImplementation) DeleteOne(db *mongo.Database, ctx context.Context, id string) (deletedCount int64, err error) {
	result, err := db.Collection("users").DeleteOne(ctx, bson.M{"_id": id}, nil)
	if err != nil {
		return
	}
	deletedCount = result.DeletedCount
	return
}

func (repository *MongodbRepositoryImplementation) DeleteMany(db *mongo.Database, ctx context.Context) (deletedCount int64, err error) {
	result, err := db.Collection("users").DeleteMany(ctx, bson.M{}, nil)
	if err != nil {
		return
	}
	deletedCount = result.DeletedCount
	return
}
