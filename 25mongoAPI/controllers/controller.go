package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swapnil/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Golang:golang@cluster0.2pndy.mongodb.net/"

const dbName = "Netflix"
const collName = "watchlist"

var collection *mongo.Collection


//connect to MongoDB
func init() {

	// set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Collection instance created!")
}

// mongo db helper function

// insert one document

func insertOneMovie(movie models.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}


// update one document

func updateOneMovie(movieId string) {
id , _ := primitive.ObjectIDFromHex(movieId)

filter := bson.M{"_id": id}
update := bson.M{"$set": bson.M{"watched": true}}
result, err:=collection.UpdateOne(context.Background(), filter, update)
if err != nil {
	log.Fatal(err)
}
	fmt.Println("Updated a Single Record ", result.ModifiedCount)
}


// delete one document

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted a Single Record ", result.DeletedCount)
}

// delete all documents

func deleteAllMovies() {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted All Records ", result.DeletedCount)
}

// get all documents

func getAllMovies() []primitive.M{
  cur , err := collection.Find(context.Background(), bson.D{{}})
  if err != nil {
	  log.Fatal(err)
  }
  var movies []primitive.M

  for cur.Next(context.Background()) {
	  var movie bson.M
	  err := cur.Decode(&movie)
	  if err != nil {
		  log.Fatal(err)
	  }
	  movies = append(movies, movie)
  }
  defer cur.Close(context.Background())	
  return movies
}


// actual controller functions 

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	var movies = getAllMovies()
	json.NewEncoder(w).Encode(movies)


}