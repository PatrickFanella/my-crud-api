package handlers 

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patrickfanella/my-crud-api/config"
	"github.com/patrickfanella/my-crud-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"
)

//Create a new User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("ConnectToMongoDB() failed: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("json.NewDecoder().Decode() failed: %v", err), http.StatusBadRequest)
		return
	}

	collection := client.Database("my-crud-api").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("InsertOne() failed: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}