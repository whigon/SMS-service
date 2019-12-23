package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sms/api"
)

var (
	listenAddress = flag.String("listen-address", ":9876", "The address to listen on for HTTP requests.")
)

func main() {
	http.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
		_ = api.SendSMS(w, r)
	})
	if os.Getenv("PORT") != "" {
		*listenAddress = ":" + os.Getenv("PORT")
	}

	log.Printf("Listening on %s", *listenAddress)

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
