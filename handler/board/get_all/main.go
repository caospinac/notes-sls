package main

import (
	"context"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/repository"
	"github.com/caospinac/notes-sls/service"
	"github.com/caospinac/notes-sls/util"
)

var (
	svc service.BoardService
)

func boardGetAll(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	data, err := svc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := util.NewResponse().
		WithBody(data)

	return res, nil
}

func main() {
	boardRepo := repository.NewBoardRepo()
	svc = service.NewBoardService(boardRepo)

	util.Start(boardGetAll)
}
