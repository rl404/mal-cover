package utils

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/mal-cover/internal/errors"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				RespondWithCSS(
					w,
					http.StatusInternalServerError,
					"",
					stack.Wrap(r.Context(),
						fmt.Errorf("%s", debug.Stack()),
						fmt.Errorf("%v", rvr),
						errors.ErrInternalServer))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
