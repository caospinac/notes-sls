package domain

type Board struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"title" dynamodbav:"title,omitempty"`
}

type Note struct {
	BoardID     string `json:"board_id" dynamodbav:"board_id"`
	NoteID      string `json:"note_id" dynamodbav:"note_id"`
	Title       string `json:"title" dynamodbav:"title,omitempty"`
	Description string `json:"description" dynamodbav:"description,omitempty"`
}
