package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/go-chi/chi/middleware"
)

// RespondWithCSS to write response with CSS format.
func RespondWithCSS(w http.ResponseWriter, statusCode int, data string, err error) {
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(statusCode)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_, _ = w.Write([]byte(data))
	}
}

// Recoverer is custom recoverer middleware.
// Will return 500.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}

				RespondWithCSS(w, http.StatusInternalServerError, "", errors.New("panic"))
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
