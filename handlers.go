package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

var system string

func statusHandler(w http.ResponseWriter, r *http.Request) {

	system = systems[rand.Intn(len(systems))]

	json.NewEncoder(w).Encode(map[string]string{
		"damaged_system": system,
	})
	w.WriteHeader(http.StatusOK)
}

func repairHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
	<title>Repair</title>
</head>
<body>
<div class="anchor-point">` + systemsCatalog[system] + `</div>
</body>
</html>`))
}

func teapotHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
