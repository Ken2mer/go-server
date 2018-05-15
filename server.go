package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// cf.
// https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
// https://blog.golang.org/context#TOC_3.2.
// https://deeeet.com/writing/2017/02/23/go-context-value/
type key string

const paramsKey key = "params"

func recoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	})
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	ps, ok := r.Context().Value(paramsKey).(httprouter.Params)
	if !ok {
		fmt.Printf("ps is not type httprouter.Params")
	}
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// We could also put *httprouter.Router in a field to not get access to the original methods (GET, POST, etc. in uppercase)
// cf. https://gist.github.com/nmerouze/5ed810218c661b40f5c4
type router struct {
	*httprouter.Router
}

func (r *router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func newRouter() *router {
	return &router{httprouter.New()}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// cf. https://github.com/julienschmidt/httprouter/issues/198
		ctx := r.Context()
		ctx = context.WithValue(ctx, paramsKey, ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

func main() {
	commonHandlers := alice.New(loggingHandler, recoverHandler)
	router := newRouter()
	router.Get("/hello/:name", commonHandlers.ThenFunc(helloHandle))
	http.ListenAndServe(":8080", router)
}
