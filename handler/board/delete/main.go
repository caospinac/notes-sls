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

func boardDelete(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	err := svc.Delete(ctx, event.PathParameters["id"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse()

	return res, nil
}

func main() {
	boardRepo := repository.NewBoardRepo()
	svc = service.NewBoardService(boardRepo)

	util.Start(boardDelete)
}
