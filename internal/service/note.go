package service

import (
	"context"

	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/repository"
	"github.com/caospinac/notes-sls/pkg/util"
)

type NoteService interface {
	CreateDefault(context.Context, string) (*domain.Note, util.ApiError)
	GetAll(context.Context, string) ([]domain.Note, util.ApiError)
	Update(context.Context, string, string, domain.UpdateNoteRequest) util.ApiError
	Delete(context.Context, string, string) util.ApiError
}

type noteService struct {
	repo repository.NoteRepo
}

func NewNoteService(repo repository.NoteRepo) NoteService {
	return &noteService{
		repo,
	}
}

func (s *noteService) CreateDefault(ctx context.Context, boardID string) (*domain.Note, util.ApiError) {
	newNote := domain.Note{}
	err := s.repo.Create(ctx, boardID, &newNote)
	if err != nil {
		return nil, util.ToApiError(err)
	}

	return &newNote, nil
}

func (s *noteService) GetAll(ctx context.Context, boardID string) ([]domain.Note, util.ApiError) {
	boards, err := s.repo.GetAll(ctx, boardID)
	if err != nil {
		return nil, util.ToApiError(err)
	}

	return boards, nil
}

func (s *noteService) Update(ctx context.Context, boardID, noteID string, newData domain.UpdateNoteRequest) util.ApiError {
	err := s.repo.Update(ctx, boardID, noteID, newData)
	if err != nil {
		return util.ToApiError(err)
	}

	return nil
}

func (s *noteService) Delete(ctx context.Context, boardID, noteID string) util.ApiError {
	if err := s.repo.Delete(ctx, boardID, noteID); err != nil {
		return util.ToApiError(err)
	}

	return nil
}
