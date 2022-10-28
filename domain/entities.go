package domain

type Board struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"title" dynamodbav:"title"`
}

type Note struct {
	ID          string `json:"id" dynamodbav:"id"`
	BoardID     string `json:"board_id" dynamodbav:"board_id"`
	Title       string `json:"title" dynamodbav:"title"`
	Description string `json:"description" dynamodbav:"description"`
}
