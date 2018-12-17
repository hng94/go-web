package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//Board Model
type Board struct {
	ID    objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string            `bson:"title" json:"title"`
}

var client *mongo.Client
var err error
var boards *mongo.Collection

//Controllers
func getBoards(w http.ResponseWriter, r *http.Request) {
	var data []Board

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := boards.Find(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Board
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result)
		data = append(data, result)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getBoard(w http.ResponseWriter, r *http.Request) {

}

func createBoard(w http.ResponseWriter, r *http.Request) {

}

func updateBoard(w http.ResponseWriter, r *http.Request) {

}

func deleteBoard(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//mongodb connection
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	boards = client.Database("taskmanager").Collection("boards")
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// cur, err := boards.Find(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result Board
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// do something with result....
	// 	fmt.Println(result, result.ID.Hex())
	// }
	defer client.Disconnect(nil)

	//Declare router
	router := mux.NewRouter()
	router.HandleFunc("/api/boards", getBoards).Methods("GET")
	router.HandleFunc("/api/boards/{id}", getBoard).Methods("GET")
	router.HandleFunc("/api/boards", createBoard).Methods("POST")
	router.HandleFunc("/api/boards/{id}", updateBoard).Methods("PUT")
	router.HandleFunc("/api/boards/{id}", deleteBoard).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
