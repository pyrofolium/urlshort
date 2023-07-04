package urlshort

import (
	_ "encoding/json"
	json2 "encoding/json"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return func(writer http.ResponseWriter, request *http.Request) {
		url, ok := pathsToUrls[request.URL.Path]
		if !ok {
			fallback.ServeHTTP(writer, request)
		} else {
			http.Redirect(writer, request, url, http.StatusFound)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type PathURL struct {
	Path string `yaml: "path" json:"path"`
	URL  string `yaml: "url" json:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	urlmap := make(map[string]string)
	var data []PathURL
	if err := yaml.Unmarshal(yml, &data); err != nil {
		return nil, err
	} else {
		for _, item := range data {
			urlmap[item.Path] = item.URL
		}
		return MapHandler(urlmap, fallback), nil
	}
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlmap := make(map[string]string)
	var data []PathURL
	if err := json2.Unmarshal(json, &data); err != nil {
		return nil, err
	} else {
		for _, item := range data {
			urlmap[item.Path] = item.URL
		}
		return MapHandler(urlmap, fallback), nil
	}
}
