package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Product_name string             `bson:"product_name"json:"name"  validate:"required"`
	Price        int                `bson:"price"json:"price"  validate:"required"`
	Type         string             `bson:"type" json:"type" validate:"required"`
	Size         string             `bson:"size" json:"size" validate:"required"`
}

type ProductRepo interface {
	FindOne(ctx context.Context, id string) (*Products, error)
	FindALL(ctx context.Context, limit int64, skip int64) ([]*Products, error)
	Created(ctx context.Context, name string, price int, t string, s string) error
	SizeFilm(ctx context.Context) ([]bson.M, error)
	TypeFilm(ctx context.Context, t string) ([]bson.M, error)
	UpdatePrice(ctx context.Context, id string, price int) error
}
