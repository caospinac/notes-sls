package helper

import "github.com/caospinac/notes-sls/internal/domain"

func MapToBoard(boardRequest *domain.PutBoardRequest) *domain.Board {
	result := &domain.Board{
		Name: boardRequest.Name,
	}

	return result
}

func MapToNote(noteRequest *domain.PutNoteRequest) *domain.Note {
	result := &domain.Note{
		Title:       noteRequest.Title,
		Description: noteRequest.Description,
	}

	return result
}
