package router

import (

#declareRouterImports()

    PulseIndicatorController__ "bankingOnGolang/internal/controller"

)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")

#declareMuxRouter()

    return router
}
