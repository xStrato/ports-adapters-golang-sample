package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xStrato/ports-adapters-go-sample/src/adapters/dtos"
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
		newProduct := dtos.NewProduct(p.GetID(), p.GetName(), string(p.GetStatus()), p.GetPrice())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
