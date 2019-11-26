package api

import (
	"encoding/json"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	encode, _ := json.Marshal("Hello!")
	w.Write(encode)
}
