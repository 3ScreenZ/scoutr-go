package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/MichaelPalmer1/scoutr-go/pkg/helpers"
// 	"github.com/MichaelPalmer1/scoutr-go/pkg/types"
// 	"github.com/julienschmidt/httprouter"
// )

// func create(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	requestUser := helpers.GetUserFromOIDC(req, api)

// 	// Parse the request body
// 	var body map[string]string
// 	err := json.NewDecoder(req.Body).Decode(&body)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	// Build the request model
// 	request := types.Request{
// 		User:      requestUser,
// 		Method:    req.Method,
// 		Path:      req.URL.Path,
// 		Body:      body,
// 		SourceIP:  req.RemoteAddr,
// 		UserAgent: req.UserAgent(),
// 	}

// 	// Create the item
// 	err = api.Create(request, body, validation)

// 	// Check for errors in the response
// 	if helpers.HTTPErrorHandler(err, w) {
// 		return
// 	}

// 	// Marshal the response and write it to output
// 	out, _ := json.Marshal(map[string]bool{
// 		"created": true,
// 	})
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(out)
// }
