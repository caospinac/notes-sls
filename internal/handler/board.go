package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/service"
	"github.com/caospinac/notes-sls/pkg/util"
)

type BoardHandler interface {
	Create(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	Get(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	GetAll(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	Update(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
	Delete(ctx context.Context, event util.EventRequest) (util.Response, util.ApiError)
}

type boardHandler struct {
	service service.BoardService
}

func NewBoardHandler(service service.BoardService) BoardHandler {
	return &boardHandler{
		service,
	}
}

func (handler boardHandler) Create(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	result, err := handler.service.CreateDefault(ctx)
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusCreated).
		WithBody(result)

	return res, nil
}

func (handler boardHandler) Get(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	data, err := handler.service.Get(ctx, event.PathParameters["id"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusOK).WithBody(data)

	return res, nil
}

func (handler boardHandler) GetAll(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	data, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusOK).WithBody(data)

	return res, nil
}

func (handler boardHandler) Update(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	newData := new(domain.PutBoardRequest)

	if err := json.Unmarshal([]byte(event.Body), newData); err != nil {
		return nil, util.NewApiError(http.StatusBadRequest)
	}

	err := handler.service.Update(ctx, event.PathParameters["id"], newData)
	if err != nil {
		return nil, err
	}

	return util.NewResponse(http.StatusOK), nil
}

func (handler boardHandler) Delete(
	ctx context.Context, event util.EventRequest,
) (util.Response, util.ApiError) {
	err := handler.service.Delete(ctx, event.PathParameters["id"])
	if err != nil {
		return nil, err
	}

	res := util.NewResponse(http.StatusOK)

	return res, nil
}
