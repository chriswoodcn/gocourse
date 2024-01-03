package errorHandle

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorWrap(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("recover error: %s", r)
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
			log.Errorf("request error: %s", err)
			switch {
			case os.IsNotExist(err):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case os.IsPermission(err):
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
	}
}
