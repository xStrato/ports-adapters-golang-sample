package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, s interfaces.IProductService) {
	r.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(s)))).Methods("GET", "OPTIONS")
}

func getProduct(s interfaces.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		p, err := s.Get(vars["id"])
		product := struct {
			Id     string  `json:"id"`
			Name   string  `json:"name"`
			Price  float32 `json:"price"`
			Status string  `json:"status"`
		}{p.GetID(), p.GetName(), p.GetPrice(), string(p.GetStatus())}

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
