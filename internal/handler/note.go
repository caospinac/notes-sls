package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/service"
	"github.com/caospinac/notes-sls/pkg/util"
)

type NoteHandler interface {
	Create(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	GetAll(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	Update(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	Delete(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
}

type noteHandler struct {
	service service.NoteService
}

func NewNoteHandler(service service.NoteService) NoteHandler {
	return &noteHandler{
		service,
	}
}

func (handler noteHandler) Create(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	newBoard, err := handler.service.CreateDefault(ctx, event.PathParameters["boardID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusCreated).WithBody(newBoard)

	return res, nil
}

func (handler noteHandler) GetAll(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	data, err := handler.service.GetAll(ctx, event.PathParameters["boardID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusOK).WithBody(data)

	return res, nil
}

func (handler noteHandler) Update(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	newData := new(domain.PutNoteRequest)

	if err := json.Unmarshal([]byte(event.Body), newData); err != nil {
		return nil, util.NewApiError(http.StatusBadRequest)
	}

	err := handler.service.Update(ctx, event.PathParameters["noteID"], newData)
	if err != nil {
		return nil, err
	}

	return util.NewResponse(http.StatusOK), nil
}

func (handler noteHandler) Delete(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	err := handler.service.Delete(ctx, event.PathParameters["boardID"], event.PathParameters["noteID"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusOK)

	return res, nil
}
