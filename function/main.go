package main

import (
	"flag"
	"net/http"

	"github.com/GoogleCloudPlatform/cloud-functions-go/nodego"
	"github.com/g-harel/svgsaurus"
)

func main() {
	flag.Parse()
	defer nodego.TakeOver()

	http.HandleFunc(nodego.HTTPTrigger, func(w http.ResponseWriter, r *http.Request) {
		svg, err := new(svgsaurus.Config).FromQuery(r.URL.Query()).Render()
		if err != nil {
			nodego.ErrorLogger.Print(err)
			status := http.StatusInternalServerError
			http.Error(w, http.StatusText(status), status)
			return
		}

		w.Header().Add("Content-Type", "image/svg+xml")
		w.Header().Add("Cache-Control", "public, max-age=900, s-maxage=900")
		w.Write(svg)
	})
}
