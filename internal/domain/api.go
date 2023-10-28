package domain

type UpdateBoardRequest struct {
	Title string `json:"title"`
}

type UpdateNoteRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
