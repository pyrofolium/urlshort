package main

import (
	"flag"
	"fmt"
	"github.com/pyrofolium/urlshort"
	"io"
	"net/http"
	"os"
)

func getyamlfile(yamlfile string) []byte {
	if len(yamlfile) != 0 {
		file, err := os.Open(yamlfile)
		if err != nil {
			fmt.Println(err)
			panic(err)
		} else {
			yamlBytes, err := io.ReadAll(file)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			return yamlBytes
		}
	} else {
		fmt.Println("yaml file could not be evaluated, using default.")
		return []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	}
}

func getjsonfile(jsonfile string) []byte {
	if len(jsonfile) != 0 {
		file, err := os.Open(jsonfile)
		if err != nil {
			fmt.Println(err)
			panic(err)
		} else {
			jsonBytes, err := io.ReadAll(file)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			return jsonBytes
		}
	} else {
		fmt.Println("yaml file could not be evaluated, using default.")
		return []byte(`
[
  {
    "path": "/urlshort",
    "url": "https://github.com/gophercises/urlshort"
  },
  {
    "path": "/urlshort-final",
    "url": "https://github.com/gophercises/urlshort/tree/solution"
  }
]
`)
	}
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	//pathsToUrls := map[string]string{
	//	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	//	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	//}
	//mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//Build the YAMLHandler using the mapHandler as the
	//fallback

	//yamlfile := flag.String("yamlfile", "", "yaml file")
	jsonfile := flag.String("jsonfile", "", "json file")
	flag.Parse()
	//yaml := getyamlfile(*yamlfile)
	json := getjsonfile(*jsonfile)
	//yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	jsonHandler, err := urlshort.JSONHandler(json, mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	//http.ListenAndServe(":8080", yamlHandler)
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
