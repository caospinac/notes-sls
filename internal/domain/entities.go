package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Board struct {
	ID      primitive.ObjectID   `json:"_id" bson:"_id"`
	NoteIDs []primitive.ObjectID `json:"note_ids" bson:"note_ids"`
	Name    string               `json:"name" bson:"name"`
}

type Note struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	BoardID     primitive.ObjectID `json:"board_id" bson:"board_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}
