package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("carousell-challenge", func() { // API defines the microservice endpoint and
	Title("Carousell Challenge")
	Description("Carousell Challenge.") // and exactly one API definition appearing in
	Scheme("http", "https")             // the design.
	Version("0.1")
	BasePath("api/v1")           // Base path to all API endpoints
	Consumes("application/json") // Media types supported by the API
	Produces("application/json") // Media types generated by the API
	Origin("*", func() {         // Define CORS policy, may be prefixed with "*" wildcard
		Headers("*")                                     // One or more authorized headers, use "*" to authorize all
		Methods("GET", "POST", "DELETE", "PUT", "PATCH") // One or more authorized HTTP methods
		Expose("X-Time")                                 // One or more headers exposed to clients
		MaxAge(600)                                      // How long to cache a preflight request response
		Credentials()                                    // Sets Access-Control-Allow-Credentials header
	})
})