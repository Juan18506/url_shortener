package main

import (
	"net/http"
)

type PathUrlMap map[string]string

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func mapHandler(pathsToUrls PathUrlMap, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		}

		fallback.ServeHTTP(w, r)
	}
}

func buildMap(urlData []PathUrl) PathUrlMap {
	pathsToUrls := make(PathUrlMap)

	for _, pathUrl := range urlData {
		pathsToUrls[pathUrl.Path] = pathUrl.URL
	}

	return pathsToUrls
}
