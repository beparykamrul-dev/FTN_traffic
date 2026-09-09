package httpapi

import (
 "encoding/json"
 "net/http"
 "github.com/beparykamrul-dev/FTN_traffic/internal/latency"
)

type PathView struct { Paths []latency.Path `json:"paths"`; Selected string `json:"selected"` }

func PathHandler(view func() PathView) http.HandlerFunc { return func(w http.ResponseWriter,r *http.Request){ if r.Method!="GET" { w.WriteHeader(http.StatusMethodNotAllowed); return }; w.Header().Set("Content-Type","application/json"); _=json.NewEncoder(w).Encode(view()) } }
