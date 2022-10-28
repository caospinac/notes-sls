package service

import (
	"context"
	"net/http"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/repository"
	"github.com/caospinac/notes-sls/util"
	"github.com/google/uuid"
)

type BoardService interface {
	CreateDefault(context.Context) (*domain.Board, util.ApiError)
	GetAll(context.Context) ([]domain.Board, util.ApiError)
	Get(context.Context, string) (*domain.Board, util.ApiError)
	Update(context.Context, string, domain.UpdateBoardRequest) util.ApiError
	Delete(context.Context, string) util.ApiError
}

type boardService struct {
	repo repository.BoardRepo
}

func NewBoardService(repo repository.BoardRepo) BoardService {
	return &boardService{
		repo: repo,
	}
}

func (s *boardService) CreateDefault(ctx context.Context) (*domain.Board, util.ApiError) {
	newBoard := domain.Board{
		ID:   uuid.NewString(),
		Name: "Untitled",
	}
	err := s.repo.Create(ctx, newBoard)
	if err != nil {
		return nil, util.NewApiError()
	}

	return &newBoard, nil
}

func (s *boardService) GetAll(ctx context.Context) ([]domain.Board, util.ApiError) {
	boards, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, util.ToApiError(err)
	}

	return boards, nil
}

func (s *boardService) Get(ctx context.Context, id string) (*domain.Board, util.ApiError) {
	board, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, util.NewApiError()
	}

	if board == nil {
		return nil, util.NewApiError().WithStatus(http.StatusNotFound)
	}

	return board, nil
}

func (s *boardService) Update(ctx context.Context, id string, newData domain.UpdateBoardRequest) util.ApiError {
	err := s.repo.Update(ctx, id, newData)
	if err != nil {
		return util.NewApiError().WithMessage(err.Error())
	}

	return nil
}

func (s *boardService) Delete(ctx context.Context, id string) util.ApiError {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return util.NewApiError()
	}

	return nil
}
