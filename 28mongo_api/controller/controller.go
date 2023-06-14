package controller

import (
	"context"
	"fmt"
	"log"

	"./28mongo_api/model"
	"command-line-argumentsC:\\Users\\Dell\\Documents\\go_Workspace\\28mongo_api\\model\\models.go"
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://shrutieddula:<Pranav@2020>@cluster0.5lyldpq.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colectionName = "watchList"

var collection *mongo.Collection

// connect with mongo DB
// initializatin method whihc runs first ime and onky 1 times
func init() {
	//client options
	//it basically tell the uri we are gonna use
	clientOptions := options.client().ApplyURI(connectionString)

	//connect to mongodb
	//we have  types of context background and TODO.toDO can be used for any tyoe of operation
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	//connect inside the database
	collection = client.Database(dbname).Collection(colectionName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//mongoDB helpers -seperate file

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func update1movie(movieId string) {
	//first we will take the id
	//this will gives us the id
	id, _ := primitive.ObjectIDFromHex(movieId)

	//in mongoDB we have _id
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}
