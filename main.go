package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-todos/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	DOB           string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Phone_number  string             `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Email_address string             `json:"email_add,omitempty" bson:"email_ add,omitempty"`
}

var client *mongo.Database

func CreateUserEndpoint(response http.ResponseWriter, request http.Request) {
	response.Header().Add("content-type", "User/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	collection := database.Database("ContactTracing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func main() {
	conf := config.GetConfiguration()
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(conf)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	database, _ = mongo.Connect(ctx, "mongodb://localhost:8080", r)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Meathods("Post")
	http.ListenAndServe(":123456", router)

}
