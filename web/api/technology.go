package api

import (
	"encoding/json"
	"net/http"

	"github.com/sensepost/gowitness/pkg/log"
	"github.com/sensepost/gowitness/pkg/models"
)

type technologyListResponse struct {
	Value []string `json:"technologies"`
}

func (h *ApiHandler) TechnologyListHandler(w http.ResponseWriter, r *http.Request) {
	var results = &technologyListResponse{}

	if err := h.DB.Model(&models.Technology{}).Distinct("value").
		Find(&results.Value).Error; err != nil {

		log.Error("could not find distinct technologies", "err", err)
		return
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
