package v1

import (
	"context"
	"encoding/json"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	lot_usecase "github.com/ivzakom/web-scraping-practice/internal/domain/usecase/lot"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	lotURL  = "/lots/:lot_id"
	lotsURL = "/lots"
)

type LotUseCase interface {
	CreateLot(ctx context.Context, dto lot_usecase.CreateLotDto) (string, error)
	GetAllLots(ctx context.Context) ([]entity.LotView, error)
	UpdateLots(ctx context.Context) error
}

type lotHandler struct {
	lotUseCase LotUseCase
}

func (h *lotHandler) Register(r *httprouter.Router) {
	r.GET(lotsURL, h.GetAllLots)
}

func NewLotHandler(lotUseCase LotUseCase) *lotHandler {
	return &lotHandler{lotUseCase: lotUseCase}
}

func (h *lotHandler) CreateLot(w http.ResponseWriter, r *http.Request, params httprouter.Param) {
	//TODO
}

func (h *lotHandler) GetAllLots(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	lots, err := h.lotUseCase.GetAllLots(context.Background())
	if err != nil {
		return
	}
	marshal, err := json.Marshal(lots)
	if err != nil {
		return
	}
	w.Write([]byte(marshal))
	w.WriteHeader(http.StatusOK)
}

func (h *lotHandler) UpdateLots(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := h.lotUseCase.UpdateLots(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
