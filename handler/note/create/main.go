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
	svc service.NoteService
)

func noteCreate(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	newBoard, err := svc.CreateDefault(ctx, event.PathParameters["boardID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse().
		WithBody(newBoard).
		WithStatus(http.StatusCreated)

	return res, nil
}

func main() {
	noteRepo := repository.NewNoteRepo()
	svc = service.NewNoteService(noteRepo)

	util.Start(noteCreate)
}
