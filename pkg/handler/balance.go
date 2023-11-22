package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/alangadiel/stori-challenge/pkg/model"
)

func (h Handler) PostBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read the request body
	var body []byte
	{
		var err error
		if body, err = io.ReadAll(r.Body); err != nil {
			log.Default().Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	// read fileName and email from request json
	var br model.BalanceRequest
	if err := json.Unmarshal(body, &br); err != nil {
		log.Default().Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.Service.PostBalance(ctx, br.FileName, br.Email); err != nil {
		log.Default().Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
