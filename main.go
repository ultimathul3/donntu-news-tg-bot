package main

import (
	"crypto/tls"
	"donntu-news-tg-bot/handler"
	"donntu-news-tg-bot/logger"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

func init() {
	initLogger()
	initEnv()
}

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("DOMAIN")),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", handler.HandleRequest)

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	logger.Log.Info("server started")

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	check := func(variable string) {
		if os.Getenv(variable) == "" {
			log.Fatalf("%s not set in .env file", variable)
		}
	}

	check("DOMAIN")
	check("ACCESS_TOKEN")
	check("SECRET_TOKEN")
}

func initLogger() {
	var err error

	logger.Log, err = logger.New("log")
	if err != nil {
		log.Fatal(err)
	}
}
