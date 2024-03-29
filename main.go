// Copyright 2019 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v0.0.1"
)

var (
	port = flag.Int("port", 8013, "The port to listen on.")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kto nie skacze ten z UJ!\n")
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

	http.HandleFunc("/whoami", func(w http.ResponseWriter, r *http.Request) {
		hostName, _ := os.Hostname()
		fmt.Fprintf(w, "I am %s\n", hostName)
	})

	s := http.Server{Addr: fmt.Sprintf(":%d", *port)}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	cleanShutdown(s)
}

func cleanShutdown(s http.Server) {
	signalChan := make(chan os.Signal, 1)
        signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
        <-signalChan

        fmt.Println("Shutdown signal received, exiting...")

        s.Shutdown(context.Background())
}
