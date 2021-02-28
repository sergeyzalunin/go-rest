package server

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	Code int
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.Code = statusCode
}
