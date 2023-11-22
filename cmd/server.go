package main

import (
	"context"
	"log"
	"net/http"

	"github.com/alangadiel/stori-challenge/pkg/handler"
	"github.com/alangadiel/stori-challenge/pkg/repo"
	"github.com/alangadiel/stori-challenge/pkg/srv"
)

const (
	ServerPort           = "8080"
	TransactionsEndpoint = "/transactions"
)

func main() {
	ctx := context.Background()

	var r repo.Repository
	{
		var err error
		r, err = repo.CreateRepository(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close(ctx)
	}

	h := handler.Handler{
		Service: srv.Service{
			Repository: r,
		},
	}

	// Configure endpoint
	http.HandleFunc(TransactionsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		h.PostTransactions(w, r)
	})

	// Start server
	log.Println("Server started at port " + ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, nil))
}
