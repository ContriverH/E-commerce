package models

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	CategoryID  int                `json:"categoryId" bson:"categoryId"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	Stock       decimal.Decimal    `json:"stock" bson:"stock"`
	Price       decimal.Decimal    `json:"price" bson:"price"`
	Version     int                `json:"version" bson:"version"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Products []Product

// NewProduct function for initialise Product model
func NewProduct() *Product {
	// now := time.Now()
	return &Product{
		Version:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
