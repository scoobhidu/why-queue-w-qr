package MongoJWTmodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type JwtCollection struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TeacherID string             `bson:"TeacherID,omitempty"`
	JWT       []Jwt              `bson:"JWT,omitempty"`
}

type Jwt struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}
