package web

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AppHandler func(http.ResponseWriter, *http.Request) error
type WebError interface {
	error
	Message() string
}

func ErrorWrap(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("panic error: %s", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(w, r)
		if err != nil {
			log.SetFormatter(&log.TextFormatter{
				ForceColors:     true,
				DisableColors:   false,
				ForceQuote:      false,
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339,
			})
			log.Errorf("request url : %s  method: %s error: %s", r.URL.Path, r.Method, err)
			switch {
			case os.IsNotExist(err):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case os.IsPermission(err):
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			default:
				var webError WebError
				if errors.As(err, &webError) {
					http.Error(w, webError.Message(), http.StatusBadRequest)
				} else {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
			return
		}
	}
}
