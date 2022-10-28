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

func boardGet(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	data, err := svc.Get(ctx, event.PathParameters["id"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse()
	if data == nil {
		res.WithStatus(http.StatusNotFound)
	} else {
		res.WithBody(data)
	}

	return res, nil
}

func main() {
	boardRepo := repository.NewBoardRepo()
	svc = service.NewBoardService(boardRepo)

	util.Start(boardGet)
}
