package main

import (
	"context"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/repository"
	"github.com/caospinac/notes-sls/service"
	"github.com/caospinac/notes-sls/util"
)

var (
	svc service.NoteService
)

func noteGetAll(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	data, err := svc.GetAll(ctx, event.PathParameters["boardID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse().
		WithBody(data)

	return res, nil
}

func main() {
	noteRepo := repository.NewNoteRepo()
	svc = service.NewNoteService(noteRepo)

	util.Start(noteGetAll)
}
