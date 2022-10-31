package main

import (
	"crypto/tls"
	"donntu-news-tg-bot/logger"
	"donntu-news-tg-bot/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

var (
	fileLog *logger.Logger
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != os.Getenv("SECRET_TOKEN") {
		fileLog.Info("invalid secret token from request")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fileLog.Info("body read error:", err)
		return
	}

	var update types.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		fileLog.Info("json error:", err)
		return
	}

	fileLog.Info("update:", fmt.Sprintf("%+v", update))
}

func init() {
	var err error

	fileLog, err = logger.New("log")
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("DOMAIN")),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}
