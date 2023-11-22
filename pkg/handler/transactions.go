package handler

import (
	"log"
	"net/http"
)

func (h Handler) PostTransactions(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.PostTransactions(r.Context()); err != nil {
		log.Default().Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
