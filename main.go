package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type User struct {
	Name  string `json: "Field Str"`
	Id  int    `json: "Field Int"`
	Email  string `json: "Field Str"`
	Password  string `json: "Field Str"`
}

type Post struct {
	Caption  string `json: "Field Str"`
	Id  int    `json: "Field Int"`
	Url  string `json: "Field Str"`
	Timestamp  string `json: "Field Str"`
}

func (h *createUser ) post(w http.ResponseWriter, r *http.Request){
	
	h.Lock()
	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}

	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func main() {
	//initialize connection 
	clientOptions:= options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	//initialize database 
	col:= client.Database("Appointy").Collection("Users")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := User{
		Name:  "Camboda Sun",
		Id:  2,
		Email: "sun@gmail.com",
		Password: "456",
	}

	http.HandleFunc("/users",createUser(oneDoc))
	err:= http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	
	
}