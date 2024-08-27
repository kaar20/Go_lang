package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaar20/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://codewithkar:tkWtuIT3jFSBXack@cluster0.exvfl.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchList"

var collection *mongo.Collection

func init() {
	// clientOPtion
	clientOption := options.Client().ApplyURI(ConnectionString)

	// connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb Connection Sucess")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Instance is Ready")
}

// Insert 1 Movie

func insertOneMovie(movie model.NetFlix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Interested Movie with ID: ", inserted.InsertedID)
}

// UPdated 1 record

func updateOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count : ", result.ModifiedCount)
}

// Delete 1 Movie

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deltedValue, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Deleted Record :", deltedValue.DeletedCount)
}

// Delete All

func deleteAll() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Records Count :", deleted.DeletedCount)

	return deleted.DeletedCount
}

// get all Movies

func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer curr.Close(context.Background())
	return movies
}

//

func GetAllMoviesList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allmovies := getAllMovies()

	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movies model.NetFlix
	_ = json.NewDecoder(r.Body).Decode(&movies)

	insertOneMovie(movies)

	json.NewEncoder(w).Encode(movies)

}

//

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneRecord(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Successfully Deleted Movie")
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAll()

	json.NewEncoder(w).Encode("Successfully Deleted All")

}
