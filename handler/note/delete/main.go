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

func noteDelete(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	err := svc.Delete(ctx, event.PathParameters["boardID"], event.PathParameters["noteID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse()

	return res, nil
}

func main() {
	noteRepo := repository.NewNoteRepo()
	svc = service.NewNoteService(noteRepo)

	util.Start(noteDelete)
}
