package main

import (
    "fmt"
    "os"
    "banking-on-golang/internal/router"
    "banking-on-golang/internal/utils"
    "log"
    "net/http"
)


func main() {

	//----------------------------------------------------------------------------	
    // Call function to get things initialized such as environment vars, database
    // connectivity, schema migration, etc...
    //----------------------------------------------------------------------------
	utils.InitializeEnvironment()
	
    appRouter := router.Router()
    appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT") )
    fmt.Println("Starting server on the port ", appPort)

    log.Fatal(http.ListenAndServe(appPort, appRouter))
}
