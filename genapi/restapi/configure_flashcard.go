// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	apiModel "github.com/pytorchtw/flashcard-go/genapi/models"
	"github.com/pytorchtw/flashcard-go/genapi/restapi/operations"
	"github.com/pytorchtw/flashcard-go/genapi/restapi/operations/deck"
	"github.com/pytorchtw/flashcard-go/services"
)

//go:generate swagger generate server --target ../../genapi --name Flashcard --spec ../../swagger.yml

func configureFlags(api *operations.FlashcardAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FlashcardAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeckLoadDeckHandler = deck.LoadDeckHandlerFunc(
		func(params deck.LoadDeckParams) middleware.Responder {
			log.Println(*params.Body.URL)
			url := *params.Body.URL
			parts := strings.Split(url, "/")
			log.Println(parts)
			if len(parts) < 5 {
				return deck.NewLoadDeckDefault(500)
			}
			if parts[2] != "github.com" {
				log.Println(parts[2])
				return deck.NewLoadDeckDefault(500)
			}
			actualURL := fmt.Sprintf("https://api.github.com/repos/%v/%v/contents/%v", parts[3], parts[4], parts[len(parts)-1])
			log.Println(actualURL)
			myDeck, err := services.LoadDeckFromURL(actualURL)
			if err != nil {
				log.Println(err)
				return deck.NewLoadDeckDefault(500)
			}
			flashcards, err := services.MakeFlashcards(myDeck.Content)
			if err != nil {
				log.Println(err)
				return deck.NewLoadDeckDefault(500)
			}
			var cards []*apiModel.Flashcard
			for _, card := range flashcards {
				newCard := apiModel.Flashcard{}
				newCard.Front = card.Front
				newCard.Back = card.Back
				cards = append(cards, &newCard)
			}
			params.Body.Flashcards = cards
			return deck.NewLoadDeckCreated().WithPayload(params.Body)
		})

	if api.GetGreetingHandler == nil {
		api.GetGreetingHandler = operations.GetGreetingHandlerFunc(func(params operations.GetGreetingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetGreeting has not yet been implemented")
		})
	}
	if api.DeckLoadDeckHandler == nil {
		api.DeckLoadDeckHandler = deck.LoadDeckHandlerFunc(func(params deck.LoadDeckParams) middleware.Responder {
			return middleware.NotImplemented("operation deck.LoadDeck has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	/*
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			handler.ServeHTTP(w, r)
			return
		})
	*/
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		//AllowedOrigins: []string{"http://flashcardgo.com:8082"},
		AllowedOrigins: []string{"*"},
		//AllowedMethods:   []string{"GET", "POST", "PUT", "OPTION", "DELETE"},
		AllowedMethods:   []string{"POST"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(handler)
}
