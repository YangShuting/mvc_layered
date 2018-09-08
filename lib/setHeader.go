package lib

import "net/http"

func setHeaders(w http.ResponseWriter) http.ResponseWriter{
	//w.Header().Set("Content-Type", "application/josn")
	return w
}
