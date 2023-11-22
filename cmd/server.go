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
	ServerPort      = "8080"
	BalanceEndpoint = "/balance"
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

	var emailSrv srv.EmailService
	{
		var err error
		if emailSrv, err = srv.CreateEmailService(); err != nil {
			log.Fatal(err)
		}
	}

	h := handler.Handler{
		BalanceService: srv.BalanceService{
			Repository:   r,
			EmailService: emailSrv,
		},
	}

	// Configure endpoint
	http.HandleFunc(BalanceEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		h.PostBalance(w, r)
	})

	// Start server
	log.Println("Server started at port " + ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, nil))
}
