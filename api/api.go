package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"golang.org/x/net/idna"

	"github.com/thngkaiyuan/look-at-my-site/checker"
)

type API struct {
	checker checker.Checker
}

type CheckPayload struct {
	Domain string                  `json:"domain"`
	Valid  bool                    `json:"valid"`
	Checks []checker.CheckerResult `json:"checks,omitempty"`
}

func New() API {
	return API{checker: checker.New()}
}

func (api API) Check(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s\n", r.URL)
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("This API endpoint only allows %s method. \n", http.MethodGet)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	unicodeDomain := html.EscapeString(r.URL.Query().Get("domain"))
	if unicodeDomain == "" {
		http.Error(w, "Domain name not specified.", http.StatusBadRequest)
		return
	}

	comprehensive := r.URL.Query().Get("comprehensive") == "true"

	asciiDomain, err := idna.ToASCII(unicodeDomain)
	if err != nil {
		msg := fmt.Sprintf("Internal Error: Domain name conversion failed. (%s)\n", err)
		http.Error(w, msg, http.StatusInternalServerError)
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
