package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	HandleAllRequest := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-control-Allow-Origin", "*")
		w.Header().Set("Access-control-Allow-Methods", "GET , POST, PUT, PETCH, DELETE,OPTIONS")
		w.Header().Set("Access-control-Allow-Headers", "Content-Type, arafat")
		w.Header().Set("Content-Type", "Application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)

		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(HandleAllRequest)

}
