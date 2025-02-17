package response

import (
	"encoding/json"
	"net/http"
)

func RespondWithHeader(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	_ = json.NewEncoder(w).Encode(&data)
}

func RespondWithError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
