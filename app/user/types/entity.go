package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		CustomerId string             `bson:"customerId"`
		Age        int                `bson:"age"`
		Name       string             `bson:"name"`
		LastName   string             `bson:"lastName"`
	}
)
