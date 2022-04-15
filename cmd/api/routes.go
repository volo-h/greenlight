package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
  // Initialize a new httprouter router instance.
  router := httprouter.New()

  // Register the relevant methods, URL patterns and handler functions for our 
  // endpoints using the HandlerFunc() method. Note that http.MethodGet and 
  // http.MethodPost are constants which equate to the strings "GET" and "POST"
  // respectively.

  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

  router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMoviesHandler)
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

  // Add the route for the PUT /v1/movies/:id endpoint.
  router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
  router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

  // Return the httprouter instance.
  return router
}
