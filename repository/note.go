package repository

import (
	"context"
	"os"

	"github.com/caospinac/notes-sls/domain"
)

var (
	noteTableName = os.Getenv("NOTES_TABLE")
)

type NoteRepo interface {
	Create(context.Context, domain.Note) error
	Get(context.Context, string, string) (*domain.Note, error)
	GetAll(context.Context, string) ([]domain.Note, error)
	Update(context.Context, string, string, domain.UpdateNoteRequest) error
	Delete(context.Context, string, string) error
}
