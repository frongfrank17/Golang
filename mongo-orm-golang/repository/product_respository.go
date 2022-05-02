package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepository struct {
	client *mongo.Client
}

func NewProductRepository(client *mongo.Client) ProductRepo {
	return productRepository{client: client}
}

type project struct{}

func (pr productRepository) FindOne(ctx context.Context, id string) (*Products, error) {
	//findOptions := options.Find()
	//findOptions.SetLimit(100)
	var product Products

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	collection := pr.client.Database("product").Collection("products")
	collection.FindOne(ctx, bson.M{"_id": ObjId}).Decode(&product)
	return &product, nil

}

func (pr productRepository) FindALL(ctx context.Context, limit int64, skip int64) ([]*Products, error) {
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	var products []*Products
	collection := pr.client.Database("product").Collection("products")
	cur, er := collection.Find(ctx, bson.D{}, findOptions)
	if er != nil {
		log.Fatal(er)
		return nil, er
	}

	for cur.Next(context.TODO()) {
		var prod Products
		if er = cur.Decode(&prod); er != nil {
			return nil, er
		}
		products = append(products, &prod)
	}
	fmt.Println(cur)
	cur.Close(ctx)
	return products, nil
}
func (pr productRepository) Created(ctx context.Context, name string, price int, t string, s string) error {
	newProduct := Products{
		Product_name: name,
		Price:        price,
		Size:         s,
		Type:         t,
	}
	collection := pr.client.Database("product").Collection("products")
	_, err := collection.InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (pr productRepository) SizeFilm(ctx context.Context) ([]bson.M, error) {
	fmt.Println("Repo")
	collection := pr.client.Database("product").Collection("products")

	data, err := collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{

			{
				"$group", bson.D{
					{"_id", "$size"},
					{"list", bson.D{{"$push", "$product_name"}}},
				},
			},

			/*	{
				"$project", bson.D{{"_id", 1}, {"size", "$_id"}, {"list", "$product"}},
			}, */
		},
	})

	if err != nil {
		return nil, err
	}

	var result []bson.M //:= new(Res)
	err = data.All(ctx, &result)
	//fmt.Println(result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
func (pr productRepository) TypeFilm(ctx context.Context, t string) ([]bson.M, error) {
	fmt.Println("Repo")
	collection := pr.client.Database("product").Collection("products")

	data, err := collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{
			{
				"$match", bson.D{
					{"type", t},
				},
			},
		},
		bson.D{
			{
				"$group", bson.D{
					{"_id", "$type"},
					{"list", bson.D{{"$push", "$product_name"}}},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var result []bson.M //:= new(Res)
	err = data.All(ctx, &result)
	//fmt.Println(result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
func (pr productRepository) UpdatePrice(ctx context.Context, id string, price int) error {

	ObjId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	collection := pr.client.Database("product").Collection("products")
	fmt.Print(" Repo PRice ")
	fmt.Print(price)
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": ObjId},
		bson.D{
			{"$set", bson.D{{"price", price}}},
		},
	)
	fmt.Println(result)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
