package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/alangadiel/stori-challenge/pkg/model"
	"github.com/alangadiel/stori-challenge/pkg/srv"
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

	if err := h.BalanceService.PostBalance(ctx, br.FileName, br.Email); err != nil {
		if errors.Is(err, srv.ErrFileNotFound) {
			http.Error(w, srv.ErrFileNotFound.Error(), http.StatusNotFound)
			return
		}
		log.Default().Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
