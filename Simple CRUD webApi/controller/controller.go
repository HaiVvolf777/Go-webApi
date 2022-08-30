package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haider-star/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://haider:wolf@cluster0.ckdyquq.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netfilx"
const colName = "watchlist"

// most important

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo Db is Connect Sucessfully")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection insstance are Ready")
}

// insert One Movie
func insertOneMovie(movie models.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("One Movie inserted Sucessfully", inserted.InsertedID)

}

// Update one movie

func updateOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Value is ", result.ModifiedCount)
}

// Delete One Record

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movie count is ", deleteCount)
}

// delete All recourds from DB

func deleteAllMovies() int64 {
	// filter := bson.D{{}}
	// collection.DeleteMany(context.Background(), filter)
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies := append(movies, movie)
		fmt.Println(movies)
	}
	defer cursor.Close(context.Background())

	return movies

}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)

}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleleOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleleAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
