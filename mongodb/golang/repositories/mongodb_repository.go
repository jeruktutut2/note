package repositories

import (
	"context"
	modelentities "note-golang-mongodb/models/entitites"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository interface {
	Create(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (err error)
	Get(db *mongo.Database, ctx context.Context, test string) (test1s []modelentities.Test1, err error)
	GetById(db *mongo.Database, ctx context.Context, id primitive.ObjectID) (test1 modelentities.Test1, err error)
	UpdateOne(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error)
	UpdateById(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error)
	DeleteOne(db *mongo.Database, ctx context.Context, id primitive.ObjectID) (rowsAffected int64, err error)
}

type mongodbRepository struct {
}

func NewMongodbRepository() MongodbRepository {
	return &mongodbRepository{}
}

func (repository *mongodbRepository) Create(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (err error) {
	_, err = db.Collection("test1").InsertOne(ctx, test1, nil)
	return
}

func (repository *mongodbRepository) Get(db *mongo.Database, ctx context.Context, test string) (test1s []modelentities.Test1, err error) {
	cursor, err := db.Collection("test1").Find(ctx, bson.M{"test": test}, nil)
	if err != nil {
		return
	}

	defer func() {
		errCursorClose := cursor.Close(ctx)
		if errCursorClose != nil {
			test1s = []modelentities.Test1{}
			err = errCursorClose
		}
	}()

	err = cursor.All(ctx, &test1s)
	if err != nil {
		test1s = []modelentities.Test1{}
		return
	}

	if cursor.Err() != nil {
		test1s = []modelentities.Test1{}
		err = cursor.Err()
		return
	}

	return
}

func (repository *mongodbRepository) GetById(db *mongo.Database, ctx context.Context, id primitive.ObjectID) (test1 modelentities.Test1, err error) {
	err = db.Collection("test1").FindOne(ctx, bson.M{"_id": id}, nil).Decode(&test1)
	return
}

func (repository *mongodbRepository) UpdateOne(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error) {
	result, err := db.Collection("test1").UpdateOne(ctx, bson.M{"_id": test1.Id}, bson.M{"$set": bson.M{"test": test1.Test}})
	if err != nil {
		return
	}
	rowsAffected = result.ModifiedCount
	return
}

func (repository *mongodbRepository) UpdateById(db *mongo.Database, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error) {
	result, err := db.Collection("test1").UpdateByID(ctx, test1.Id, bson.M{"$set": bson.M{"test": test1.Test}}, nil)
	if err != nil {
		return
	}
	rowsAffected = result.ModifiedCount
	return
}

func (repository *mongodbRepository) DeleteOne(db *mongo.Database, ctx context.Context, id primitive.ObjectID) (rowsAffected int64, err error) {
	result, err := db.Collection("test1").DeleteOne(ctx, bson.M{"_id": id}, nil)
	if err != nil {
		return
	}
	rowsAffected = result.DeletedCount
	return
}
