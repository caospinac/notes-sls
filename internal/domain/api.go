package domain

type PutBoardRequest struct {
	Name string `json:"name,omitempty"`
}

type PutNoteRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
