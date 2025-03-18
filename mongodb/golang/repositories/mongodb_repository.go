package repositories

import (
	"context"
	modelentities "note-golang-mongodb/models/entitites"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository interface {
	// InsertOne(db *mongo.Database, ctx context.Context, user modelentities.User) (err error)
	// InsertMany(db *mongo.Database, ctx context.Context, users []modelentities.User) (err error)
	// FindOne(db *mongo.Database, ctx context.Context, email string) (user modelentities.User, err error)
	// Find(db *mongo.Database, ctx context.Context, email string) (users []modelentities.User, err error)
	// UpdateOne(db *mongo.Database, ctx context.Context, user modelentities.User) (rowsAffected int64, err error)
	// UpdateById(db *mongo.Database, ctx context.Context, id string) (rowsAffected int64, err error)
	// DeleteOne(db *mongo.Database, ctx context.Context, id string) (deletedCount int64, err error)
	// DeleteMany(db *mongo.Database, ctx context.Context) (deletedCount int64, err error)
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

// func (repository *mongodbRepository) InsertOne(db *mongo.Database, ctx context.Context, user modelentities.User) (err error) {
// 	_, err = db.Collection("users").InsertOne(ctx, user, nil)
// 	return
// }

// func (repository *mongodbRepository) InsertMany(db *mongo.Database, ctx context.Context, users []modelentities.User) (err error) {
// 	var documents []interface{}
// 	for _, user := range users {
// 		documents = append(documents, user)
// 	}
// 	_, err = db.Collection("users").InsertMany(ctx, documents, nil)
// 	return
// }

// func (repository *mongodbRepository) FindOne(db *mongo.Database, ctx context.Context, email string) (user modelentities.User, err error) {
// 	err = db.Collection("users").FindOne(ctx, bson.M{"email": email}, nil).Decode(&user)
// 	return
// }

// func (repository *mongodbRepository) Find(db *mongo.Database, ctx context.Context, email string) (users []modelentities.User, err error) {
// 	cursor, err := db.Collection("users").Find(ctx, bson.M{"email": email}, nil)
// 	if err != nil {
// 		return
// 	}
// 	defer func() {
// 		errCursor := cursor.Close(ctx)
// 		if errCursor != nil {
// 			users = []modelentities.User{}
// 			err = errCursor
// 		}
// 	}()

// 	err = cursor.All(ctx, &users)
// 	if err != nil {
// 		users = []modelentities.User{}
// 		return
// 	}

// 	if cursor.Err() != nil {
// 		users = []modelentities.User{}
// 		err = cursor.Err()
// 		return
// 	}
// 	return
// }

// func (repository *mongodbRepository) UpdateOne(db *mongo.Database, ctx context.Context, user modelentities.User) (rowsAffected int64, err error) {
// 	result, err := db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"email": user.Email, "password": user.Password}}, nil)
// 	if err != nil {
// 		return
// 	}
// 	rowsAffected = result.ModifiedCount
// 	return
// }

// func (repository *mongodbRepository) UpdateById(db *mongo.Database, ctx context.Context, id string) (rowsAffected int64, err error) {
// 	result, err := db.Collection("users").UpdateByID(ctx, id, bson.M{"$set": bson.M{"email": "email1@email.com", "password": "password@A2"}}, nil)
// 	if err != nil {
// 		return
// 	}
// 	rowsAffected = result.ModifiedCount
// 	return
// }

// func (repository *mongodbRepository) DeleteOne(db *mongo.Database, ctx context.Context, id string) (deletedCount int64, err error) {
// 	result, err := db.Collection("users").DeleteOne(ctx, bson.M{"_id": id}, nil)
// 	if err != nil {
// 		return
// 	}
// 	deletedCount = result.DeletedCount
// 	return
// }

// func (repository *mongodbRepository) DeleteMany(db *mongo.Database, ctx context.Context) (deletedCount int64, err error) {
// 	result, err := db.Collection("users").DeleteMany(ctx, bson.M{}, nil)
// 	if err != nil {
// 		return
// 	}
// 	deletedCount = result.DeletedCount
// 	return
// }

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
