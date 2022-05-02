package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  // Initialize a new httprouter router instance.
  router := httprouter.New()

  // Register the relevant methods, URL patterns and handler functions for our 
  // endpoints using the HandlerFunc() method. Note that http.MethodGet and 
  // http.MethodPost are constants which equate to the strings "GET" and "POST"
  // respectively.

  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

  // Use the requireActivatedUser() middleware on our five /v1/movies** endpoints. 
  router.HandlerFunc(http.MethodGet, "/v1/movies", app.requireActivatedUser(app.listMoviesHandler)) 
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.requireActivatedUser(app.createMovieHandler)) 

  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requireActivatedUser(app.showMovieHandler)) 
  router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requireActivatedUser(app.updateMovieHandler)) 
  router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requireActivatedUser(app.deleteMovieHandler))

  // Add the route for the POST /v1/users endpoint. 
  router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
  router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

  // Add the route for the POST /v1/tokens/authentication endpoint.
  router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

  // Wrap the router with the rateLimit() middleware. 
  return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
