package httpServer

import "net/http"

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func init() {
	AddNewFunction(CreateHandlerCommand("/update", "POST", UpdateHandler))
}
