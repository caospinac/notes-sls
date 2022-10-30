package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/caospinac/notes-sls/domain"
	"github.com/caospinac/notes-sls/repository"
	"github.com/caospinac/notes-sls/service"
	"github.com/caospinac/notes-sls/util"
)

var (
	svc service.NoteService
)

func noteUpdate(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	newData := new(domain.UpdateNoteRequest)
	res := util.NewResponse()

	if err := json.Unmarshal([]byte(event.Body), newData); err != nil {
		return nil, util.NewApiError().WithStatus(http.StatusBadRequest)
	}

	err := svc.Update(ctx, event.PathParameters["boardID"], event.PathParameters["noteID"], *newData)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	noteRepo := repository.NewNoteRepo()
	svc = service.NewNoteService(noteRepo)

	util.Start(noteUpdate)
}
