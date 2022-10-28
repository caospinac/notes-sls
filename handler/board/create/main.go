package main

import (
	"context"
	"net/http"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/repository"
	"github.com/caospinac/notes-sls/service"
	"github.com/caospinac/notes-sls/util"
)

var (
	svc service.BoardService
)

func boardCreate(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	newBoard, err := svc.CreateDefault(ctx)
	if err != nil {
		return nil, err
	}

	res := util.NewResponse().
		WithBody(newBoard).
		WithStatus(http.StatusCreated)

	return res, nil
}

func main() {
	boardRepo := repository.NewBoardRepo()
	svc = service.NewBoardService(boardRepo)

	util.Start(boardCreate)
}
