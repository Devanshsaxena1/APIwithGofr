package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

type Car struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Description string `json:"description"`
}

var collection *mongo.Collection

func addCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	_ = json.NewDecoder(r.Body).Decode(&newCar)

	result, err := collection.InsertOne(context.Background(), newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newCar.ID = int(result.InsertedID.(int64))
	writeResponse(w, http.StatusCreated, newCar)
}

func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	var cars []Car

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var car Car
		err := cur.Decode(&car)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	writeResponse(w, http.StatusOK, cars)
}

func getCarHandler(w http.ResponseWriter, r *http.Request) {
	params := getRouteParams(r)
	carID, _ := strconv.Atoi(params["id"])

	var car Car
	err := collection.FindOne(context.Background(), bson.M{"id": carID}).Decode(&car)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	writeResponse(w, http.StatusOK, car)
}

func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	params := getRouteParams(r)
	carID, _ := strconv.Atoi(params["id"])

	var updatedCar Car
	_ = json.NewDecoder(r.Body).Decode(&updatedCar)
	updatedCar.ID = carID

	_, err := collection.ReplaceOne(context.Background(), bson.M{"id": carID}, updatedCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, http.StatusOK, updatedCar)
}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	params := getRouteParams(r)
	carID, _ := strconv.Atoi(params["id"])

	_, err := collection.DeleteOne(context.Background(), bson.M{"id": carID})
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func getRouteParams(r *http.Request) map[string]string {
	return map[string]string{"id": "1"} // Replace with your route parameter logic
}

func setupDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("carDB")
	collection = db.Collection("cars")
}

func TestHandlers(t *testing.T) {
	setupDatabase()

	// AddCarHandler test
	jsonStr := []byte(`{"brand":"Toyota","model":"Corolla","description":"Test car"}`)
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	testHandler(t, addCarHandler, req, http.StatusCreated)

	// GetCarsHandler test
	req, err = http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}
	testHandler(t, getCarsHandler, req, http.StatusOK)

	// GetCarHandler test
	req, err = http.NewRequest("GET", "/cars/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	testHandler(t, getCarHandler, req, http.StatusOK)

	// UpdateCarHandler test
	updateData := []byte(`{"brand":"UpdatedBrand","model":"UpdatedModel","description":"UpdatedDescription"}`)
	req, err = http.NewRequest("PUT", "/cars/1", bytes.NewBuffer(updateData))
	if err != nil {
		t.Fatal(err)
	}
	testHandler(t, updateCarHandler, req, http.StatusOK)

	// DeleteCarHandler test
	req, err = http.NewRequest("DELETE", "/cars/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	testHandler(t, deleteCarHandler, req, http.StatusOK)
}

func testHandler(t *testing.T, h http.HandlerFunc, req *http.Request, expectedStatus int) {
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	assert.Equal(t, expectedStatus, rr.Code)
}

func main() {
	setupDatabase()

	app := gofr.NewApp()
	app.Post("/cars", addCarHandler)
	app.Get("/cars", getCarsHandler)
	app.Get("/cars/{id}", getCarHandler)
	app.Put("/cars/{id}", updateCarHandler)
	app.Delete("/cars/{id}", deleteCarHandler)

	log.Fatal(http.ListenAndServe(":8080", app.Handler()))
}
