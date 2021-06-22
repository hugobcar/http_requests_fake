package router

import (
	"encoding/json"
	"net/http"
)

func UpdateRequests(w http.ResponseWriter, r *http.Request) {
	// var req Req
	var resp Resp

	// decoder := json.NewDecoder(r.Body)

	// err := decoder.Decode(&req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer r.Body.Close()

	// *Value = req.Value

	resp = Resp{Message: "Updated"}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
