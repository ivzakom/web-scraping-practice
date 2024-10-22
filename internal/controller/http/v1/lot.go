package v1

import (
	"context"
	"encoding/json"
	"github.com/ivzakom/web-scraping-practice/internal/controller/http/dto"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	lotsURL    = "/lots"
	updateLots = "/lots/update"
	pkkInfo    = "/lots/pkkInfo"
)

type LotUseCase interface {
	GetAllLots(ctx context.Context) ([]dto.LotViewDto, error)
	UpdateLots(ctx context.Context) error
}

type lotHandler struct {
	lotUseCase LotUseCase
}

func (h *lotHandler) Register(r *httprouter.Router) {
	r.GET(lotsURL, h.GetAllLots)
	r.GET(updateLots, h.UpdateLots)
	r.GET(pkkInfo, h.UpdateLots)
}

func NewLotHandler(lotUseCase LotUseCase) *lotHandler {
	return &lotHandler{lotUseCase: lotUseCase}
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
	_, err = w.Write([]byte(marshal))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *lotHandler) UpdateLots(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := h.lotUseCase.UpdateLots(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *lotHandler) GetPkkInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//scrap := pkkRosreestr.NewGurievskGovScraper()
	//scrap.Scrap(entity.Lot{})
}
