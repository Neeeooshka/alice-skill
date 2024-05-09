package main

import (
	"encoding/json"
	"fmt"
	"github.com/Neeeooshka/alice-skill.git/internal/config"
	"github.com/Neeeooshka/alice-skill.git/internal/storage"
	"github.com/thanhpk/randstr"
	"io"
	"net/http"
	"strings"
)

type app struct {
	options config.Options
	storage storage.LinkStorage
}

func (a *app) GetBaseURL() string {
	return a.options.GetBaseURL()
}

func (a *app) GenerateShortLink() string {
	return randstr.Base62(8)
}

func newAppInstance(opt config.Options, s storage.LinkStorage) *app {
	return &app{options: opt, storage: s}
}

func (a *app) ExpanderHandler(w http.ResponseWriter, r *http.Request) {
	link, _ := strings.CutPrefix(r.URL.String(), "/")
	if fullLink, ok := a.storage.Get(link); ok {
		w.Header().Set("Location", fullLink)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (a *app) ShortenerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortLink := a.GenerateShortLink()
	a.storage.Add(shortLink, string(body))

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, a.GetBaseURL()+"/"+shortLink)
}

func (a *app) APIShortenerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortLink := a.GenerateShortLink()
	a.storage.Add(shortLink, req.URL)

	resp := struct {
		Result string `json:"result"`
	}{
		Result: a.GetBaseURL() + "/" + shortLink,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resp)
}

func (a *app) APIBatchShortenerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type reqURL struct {
		ID  string `json:"correlation_id"`
		URL string `json:"original_url"`
	}

	var req []reqURL
	var resp []storage.Batch

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, e := range req {
		resp = append(resp, storage.Batch{ID: e.ID, URL: e.URL, Result: a.GenerateShortLink()})
	}

	if err := a.storage.AddBatch(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resp)
}
