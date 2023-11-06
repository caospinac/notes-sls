package service

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/repository"
	"github.com/caospinac/notes-sls/internal/service/helper"
	"github.com/caospinac/notes-sls/pkg/util"
)

type BoardService interface {
	CreateDefault(context.Context) (*string, util.ApiError)
	GetAll(context.Context) ([]domain.Board, util.ApiError)
	Get(context.Context, string) (*domain.Board, util.ApiError)
	Update(context.Context, string, *domain.PutBoardRequest) util.ApiError
	Delete(context.Context, string) util.ApiError
}

type boardService struct {
	repo repository.BoardsRepo
}

func NewBoardService(repo repository.BoardsRepo) BoardService {
	return &boardService{
		repo,
	}
}

func (s *boardService) CreateDefault(ctx context.Context) (*string, util.ApiError) {
	newBoard := &domain.Board{
		Name:    "Untitled",
		NoteIDs: []primitive.ObjectID{},
	}
	insertedID, err := s.repo.Insert(ctx, newBoard)
	if err != nil {
		return nil, err
	}

	return &insertedID, nil
}

func (s *boardService) GetAll(ctx context.Context) ([]domain.Board, util.ApiError) {
	return []domain.Board{}, nil
}

func (s *boardService) Get(ctx context.Context, id string) (*domain.Board, util.ApiError) {
	board, err := s.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if board == nil {
		return nil, util.NewApiError(http.StatusNotFound)
	}

	return board, nil
}

func (s *boardService) Update(ctx context.Context, id string, newData *domain.PutBoardRequest) util.ApiError {
	err := s.repo.Update(ctx, id, helper.MapToBoard(newData))
	if err != nil {
		return err
	}

	return nil
}

func (s *boardService) Delete(ctx context.Context, id string) util.ApiError {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
