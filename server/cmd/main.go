package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	srvPort := ":8443"
	fs := http.FileServer(http.Dir("../client/public"))
	http.Handle("/", fs)
	server := http.Server{
		Addr:         srvPort,
		Handler:      fs,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	// router := mux.NewRouter()
	// router.HandleFunc("/api/transactions", handlers.GetTransactions).Methods("GET")
	// router.HandleFunc("/api/transaction", handlers.CreateTransaction).Methods("POST")

	log.Printf("Starting https Server on port %s\n\n", srvPort)
	go runServerTLS(&server)

	select {} // Block Main flow
}

func runServerTLS(srv *http.Server) {
	err := srv.ListenAndServeTLS("/home/rodoslav_sk/ca/proxy.local.crt.pem", "/home/rodoslav_sk/ca/proxy.local.key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
