package client

import (
	"net/http"
	"os"

	"github.com/EricFrancis12/url-checker-lambda/pkg"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	authToken := os.Getenv(pkg.EnvAuthToken)
	if authToken != "" {
		header := r.Header.Get(pkg.HttpHeaderAuthorization)
		if pkg.BearerHeader(authToken) != header {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	hostname := os.Getenv(pkg.EnvHostname)
	if hostname == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.NewData(hostname).Json())
}
