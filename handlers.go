package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	result := make([]Payload, 0)
	db.Range(func(k, v interface{}) bool {
		result = append(result, Payload{k.(string), v.(string)})
		return true
	})
	jsonResp(w, 200, result)
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if len(key) == 0 {
		jsonResp(w, 400, ErrorResponse{"error", "key is empty"})
		return
	}
	if v, ok := db.Load(key); ok {
		jsonResp(w, 200, Payload{key, v.(string)})
		return
	}
	jsonResp(w, 400, ErrorResponse{"error", "key/value pair not found"})
	return
}

func DeleteByKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if len(key) == 0 {
		jsonResp(w, 400, ErrorResponse{"error", "key is empty"})
		return
	}
	db.Delete(key)
	jsonResp(w, 200, Response{"ok", "deleted"})
	return
}

func Upsert(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	err := render.DecodeJSON(r.Body, &payload)
	if err != nil {
		jsonResp(w, 400, &ErrorResponse{
			"error",
			"Json decoding error. Make sure you send strings in key/value fields"},
		)
		return
	}
	db.Store(payload.Key, payload.Value)

	jsonResp(w, 200, Response{"ok", "key and value added/updated"})
	return
}
