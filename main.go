package api

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	mux := goji.NewMux()

	corz := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"*", "x-requested-with", "Content-Type", "If-Modified-Since", "If-None-Match"},
		ExposedHeaders: []string{"Content-Length"},
	})

	mux.Use(corz.Handler)
	mux.UseC(Recoverer)

	mux.HandleFuncC(pat.Post("/onhub/connect/:device"), OnhubConnect)
	mux.HandleFuncC(pat.Post("/onhub/disconnect/:device"), OnhubDisconnect)

	mux.HandleFuncC(pat.Post("/nest/motion-detected"), NestMotionDetected)

	mux.HandleFunc(pat.Get("/*"), func(w http.ResponseWriter, req *http.Request) {
		renderJSON(w, 400, map[string]interface{}{
			"error": "404 Not found",
		})
	})

	http.Handle("/", mux)
}

func Recoverer(inner goji.Handler) goji.Handler {
	mw := func(c context.Context, w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ctx := appengine.NewContext(req)
				log.Errorf(ctx, "Recovered error: %v", err)
				renderJSON(w, 500, map[string]interface{}{
					"error": "There was an unexpected error",
				})
			}
		}()
		inner.ServeHTTPC(c, w, req)
	}

	return goji.HandlerFunc(mw)
}

func renderJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.WriteHeader(code)

	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(d)
	}
}
