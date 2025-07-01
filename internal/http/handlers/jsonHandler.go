package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type JsonHandler func(*http.Request) (any, error)

func Wrap(handler JsonHandler) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		fmt.Println("Wrap: handler started")
		res.Header().Set("Content-Type", "application/json")

		rsBody, err := handler(req)
		if err != nil {
			fmt.Println("Wrap: handler error", err)
			res.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(res).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Wrap: encoding response")
		if err := json.NewEncoder(res).Encode(rsBody); err != nil {
			fmt.Println("Wrap: failed to encode response", err)
			http.Error(res, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}
