package repositorys

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type eventsRepository struct {
	client *mongo.Client
}

func AdpterEventsRepository(client *mongo.Client) EventsRepo {
	return eventsRepository{client: client}
}

func (events eventsRepository) FindOne(ctx context.Context, id string) (*Events, error) {
	var event Events
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	collection := events.client.Database("project").Collection("events")
	collection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&event)
	return &event, nil
}
func (events eventsRepository) FindAll(ctx context.Context, limit int64, skip int64) ([]*Events, error) {
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	var event []*Events
	collection := events.client.Database("project").Collection("events")
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var e Events
		if err = cur.Decode(&e); err != nil {
			return nil, err
		}
		event = append(event, &e)
	}
	fmt.Println(cur)
	cur.Close(ctx)
	return event, nil
}
func (events eventsRepository) Created(ctx context.Context, name string, start_time int64) error {
	datetime := time.UnixMilli(start_time)
	datenow := time.Now()
	fmt.Println(datetime)
	type Action struct {
		Code    string `bson:"code" json:"code"`
		Message string `bson:"message" json:"message" `
	}
	// action := new(Action)
	// action.Code = "1"
	// action.Message = "OPEN"

	newEvent := Events{
		Name:       name,
		Start_time: datetime,
		Action:     Action{Code: "1", Message: "OPEN"},
		Created_at: datenow,
		Updated_at: datenow,
	}
	collection := events.client.Database("project").Collection("events")
	_, err := collection.InsertOne(ctx, newEvent)
	if err != nil {
		return err
	}
	return nil
}
