package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler func(*http.Request) (any, error)

func Wrap(handler Handler) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		res.Header().Set("Content-Type", "application/json")

		rsBody, err := handler(req)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(res).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := json.NewEncoder(res).Encode(rsBody); err != nil {
			http.Error(res, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func Unwrap(h func(http.ResponseWriter, *http.Request, httprouter.Params)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r, nil)
	}
}
