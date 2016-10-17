package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"golang.org/x/net/idna"

	"github.com/thngkaiyuan/look-at-my-site/checker"
	"github.com/abiosoft/semaphore"
)

type API struct {
	checker checker.Checker
	semaphore *semaphore.Semaphore
}

type CheckPayload struct {
	Domain string                  `json:"domain"`
	Valid  bool                    `json:"valid"`
	Checks []checker.CheckerResult `json:"checks,omitempty"`
}

func New(s *semaphore.Semaphore) API {
	return API{checker: checker.New(), semaphore: s}
}

func (api API) Check(w http.ResponseWriter, r *http.Request) {
	api.semaphore.Acquire()
	fmt.Printf("Received request: %s\n", r.URL)
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("This API endpoint only allows %s method. \n", http.MethodGet)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		api.semaphore.Release()
		return
	}

	unicodeDomain := html.EscapeString(r.URL.Query().Get("domain"))
	if unicodeDomain == "" {
		http.Error(w, "Domain name not specified.", http.StatusBadRequest)
		api.semaphore.Release()
		return
	}

	comprehensive := r.URL.Query().Get("comprehensive") == "true"

	asciiDomain, err := idna.ToASCII(unicodeDomain)
	if err != nil {
		msg := fmt.Sprintf("Internal Error: Domain name conversion failed. (%s)\n", err)
		http.Error(w, msg, http.StatusInternalServerError)
		api.semaphore.Release()
		return
	}

	payload := CheckPayload{
		Domain: asciiDomain,
		Valid:  false,
	}

	var results []checker.CheckerResult
	if comprehensive {
		results = api.checker.CheckAll(asciiDomain)
	} else {
		results = api.checker.CheckBasic(asciiDomain)
	}

	if len(results) > 0 {
		payload.Checks = results
		payload.Valid = true
	}

	respondWithJSON(w, payload)
	api.semaphore.Release()
}

func respondWithJSON(w http.ResponseWriter, v interface{}) {
	payload, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		msg := fmt.Sprintf("Internal error: JSON marshalling failed. (%s)\n", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "%s\n", payload)
}
