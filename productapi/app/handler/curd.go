package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// db "github.com/codernishchay/productapi/app"
	"github.com/codernishchay/productapi/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var limit int64 = 10

func Createproduct(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	product := new(models.Product)
	err := json.NewDecoder(req.Body).Decode(product)
	if err != nil {
		ResponseWriter(res, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	result, err := db.Collection("products").InsertOne(context.TODO(), product)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			ResponseWriter(res, http.StatusNotAcceptable, "username or email already exists in database.", nil)
		default:
			ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	ResponseWriter(res, http.StatusCreated, "", product)
}

func GetProducts(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var ProductList []models.Product
	pageString := req.FormValue("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		page = 0
	}
	page = page * limit
	findOptions := options.FindOptions{
		Skip:  &page,
		Limit: &limit,
		Sort: bson.M{
			"_id": -1, // -1 for descending and 1 for ascending
		},
	}
	curser, err := db.Collection("products").Find(context.TODO(), bson.D{}, &findOptions)
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &ProductList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	fmt.Printf("Products  are here")
	ResponseWriter(res, http.StatusOK, "", ProductList)
}
