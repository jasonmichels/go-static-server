package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jasonmichels/go-static-server/middleware"
	"github.com/newrelic/go-agent"
)

// PublicPath The location to the public folder
const PublicPath = "./public"

// IndexHTML The default index.html location
const IndexHTML = PublicPath + "/index.html"

// Get the file to serve from the path name
func getFileToServe(url *url.URL, pathPrefix string) string {
	var file string

	if pathPrefix == "/" {
		pathPrefix = ""
	}

	switch filepath.Ext(url.Path) {
	case ".html", ".htm", "":
		file = IndexHTML
	default:
		path := strings.TrimPrefix(url.Path, pathPrefix)
		file = PublicPath + path
	}

	return file
}

// Handle serving static files
func handler(w http.ResponseWriter, r *http.Request) {
	filePath := getFileToServe(r.URL, os.Getenv("PATH_PREFIX"))
	http.ServeFile(w, r, filePath)
}

func main() {
	newRelicName := os.Getenv("NEW_RELIC_NAME")
	newRelicKey := os.Getenv("NEW_RELIC_KEY")

	config := newrelic.NewConfig(newRelicName, newRelicKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Println("New relic issue and not being used: ", err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":80"
	}

	http.Handle(middleware.NewRelicMiddleware(app, "/", middleware.LoggingMiddleware(handler)))
	http.ListenAndServe(port, nil)
}
