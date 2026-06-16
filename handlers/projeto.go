package handlers

import (
	"encoding/json"
	"http-server-projeto-korp/services"
	"net/http"
)

func ProjetoKorp(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	dados := services.GerarRespostaJson()

	err := json.NewEncoder(w).Encode(dados)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
