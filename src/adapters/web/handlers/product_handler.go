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
	r.Handle("/product", n.With(negroni.Wrap(createProduct(s)))).Methods("POST", "OPTIONS")
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

func createProduct(s interfaces.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var pDto dtos.Product
		err := json.NewDecoder(r.Body).Decode(&pDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		p, err := s.Create(pDto.Name, pDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
