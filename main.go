package main

//
import (
	"github.com/corinnekrych/graphql-service/handler"
	"github.com/corinnekrych/graphql-service/loader"
	"github.com/corinnekrych/graphql-service/resolver"
	"github.com/corinnekrych/graphql-service/schema"
	"github.com/corinnekrych/graphql-service/witapi/client"
	goaclient "github.com/goadesign/goa/client"
	graphql "github.com/graph-gophers/graphql-go"

	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Tweak configuration values here.
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	// Create client struct
	witClient := client.New(goaclient.HTTPClientDoer(http.DefaultClient))
	witClient.Host = "openshift.io"

	root, err := resolver.NewRoot(witClient)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
		return
	}
	sch, err := graphql.ParseSchema(schema.String(), root)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
		return
	}
	// Create the request handler; inject dependencies.
	h := handler.GraphQL{
		// Parse and validate schema. Panic if unable to do so.
		Schema:  sch,
		Loaders: loader.Initialize(witClient),
	}
	// Register handlers to routes.
	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h) // Register without a trailing slash to avoid redirect.

	// Configure the HTTP server.
	s := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
	// Begin listening for requests.
	log.Printf("Listening for requests on %s", s.Addr)

	if err = s.ListenAndServe(); err != nil {
		log.Println("server.ListenAndServe:", err)
	}
}
