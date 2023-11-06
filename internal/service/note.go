package service

import (
	"context"

	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/repository"
	"github.com/caospinac/notes-sls/internal/service/helper"
	"github.com/caospinac/notes-sls/pkg/util"
)

type NoteService interface {
	CreateDefault(context.Context, string) (*string, util.ApiError)
	GetAll(context.Context, string) ([]domain.Note, util.ApiError)
	Update(context.Context, string, *domain.PutNoteRequest) util.ApiError
	Delete(context.Context, string, string) util.ApiError
}

type noteService struct {
	repo repository.NotesRepo
}

func NewNoteService(repo repository.NotesRepo) NoteService {
	return &noteService{
		repo,
	}
}

func (s *noteService) CreateDefault(ctx context.Context, boardID string) (*string, util.ApiError) {
	newNote := domain.Note{}
	insertedID, err := s.repo.Insert(ctx, boardID, &newNote)
	if err != nil {
		return nil, err
	}

	return &insertedID, nil
}

func (s *noteService) GetAll(ctx context.Context, boardID string) ([]domain.Note, util.ApiError) {
	boards, err := s.repo.FindByBoardID(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s *noteService) Update(ctx context.Context, noteID string, newData *domain.PutNoteRequest) util.ApiError {
	err := s.repo.Update(ctx, noteID, helper.MapToNote(newData))
	if err != nil {
		return err
	}

	return nil
}

func (s *noteService) Delete(ctx context.Context, boardID, noteID string) util.ApiError {
	if err := s.repo.Delete(ctx, noteID); err != nil {
		return err
	}

	return nil
}
