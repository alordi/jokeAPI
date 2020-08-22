package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Joke struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	JokeId   int32            `json:"jokeId,omitempty" bson:"jokeId,omitempty"`
	Setup  string             `json:"setup" bson:"setup,omitempty"`
	Punchline string           `json:"punchline" bson:"punchline,omitempty"`
	Type string              `json:"type" bson:"type,omitempty"`
}