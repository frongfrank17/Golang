package repositorys

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Events struct {
	_id    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name   string             `bson:"name" json:"name" validate:"required"`
	Action struct {
		Code    string `bson:"code" json:"code"`
		Message string `bson:"message" json:"message" `
	} `bson:"action"json:"action"`

	Start_time time.Time `bson:"start_time" json:"start_time" validate:"required"`

	Created_at time.Time `bson:"created_at" json:"created_at , omitempty"`
	Updated_at time.Time `bson:"updated_at" json:"updated_at ,omitempty"`
}

type EventsRepo interface {
	FindOne(ctx context.Context, _id string) (*Events, error)
	FindAll(ctx context.Context, limit int64, skip int64) ([]*Events, error)
	Created(ctx context.Context, name string, start_time int64) error
	//Updated_action(ctx context.Context, code string, message string) error
}
