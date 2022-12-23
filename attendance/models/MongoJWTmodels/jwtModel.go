package MongoJWTmodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCollection struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TeacherID string             `bson:"TeacherID,omitempty"`
	JWT       []Jwt              `bson:"JWT,omitempty"`
}

type Jwt struct {
	Token     string `bson:"omitempty"`
	Timestamp int    `bson:"omitempty"`
}
