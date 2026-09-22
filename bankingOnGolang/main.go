package main

import (
    "fmt"
    "os"
    "time"
    "bankingOnGolang/internal/router"
    "bankingOnGolang/internal/utils"
    "log"
    "net/http"
)

func main() {

	// ----------------------------------------------------------------------------
    // Call function to get things initialized such as environment vars, database
    // connectivity, schema migration, etc...
    // ----------------------------------------------------------------------------
	utils.InitializeEnvironment()
	
    appRouter := router.Router()
    appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT") )
    fmt.Println("Starting server on the port ", appPort)

    server := &http.Server{
        Addr:         appPort,
        Handler:      appRouter,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Fatal(server.ListenAndServe())
}
