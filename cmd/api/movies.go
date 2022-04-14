package main

import (
  "net/http"
  "fmt"
  "time" // New import
  "greenlight.alexedwards.net/internal/data" // New import
  // "greenlight.alexedwards.net/internal/validator" // New import
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
  id, err := app.readIDParam(r) 
  if err != nil { 
    http.NotFound(w, r)
    return 
  }

  // Create a new instance of the Movie struct, containing the ID we extracted from 
  // the URL and some dummy data. Also notice that we deliberately haven't set a 
  // value for the Year field.

  movie := data.Movie{
    ID: id,
    CreatedAt: time.Now(),
    Title: "Casablanca",
    Runtime: 102,
    Genres: []string{"drama", "romance", "war"},
    Version: 1,
  }

  // Encode the struct to JSON and send it as the HTTP response. 
  err = app.writeJSON(w, http.StatusOK, movie, nil)
  if err != nil {
    app.logger.Println(err)
    http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
  }
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
  // Declare an anonymous struct to hold the information that we expect to be in the
  // HTTP request body (note that the field names and types in the struct are a subset 
  // of the Movie struct that we created earlier). This struct will be our *target 
  // decode destination*.

  var input struct {
    Title string `json:"title"`
    Year int32 `json:"year"`
    Runtime data.Runtime `json:"runtime"` // Make this field a data.Runtime type.
    Genres []string `json:"genres"` }

  // Use the new readJSON() helper to decode the request body into the input struct. 
  // If this returns an error we send the client the error message along with a 400 
  // Bad Request status code, just like before.
  err := app.readJSON(w, r, &input) 
  if err != nil {
    app.logger.Println(err)
    // # TODO: correct pass error to http.Error
    http.Error(w, "bad", http.StatusInternalServerError)
  }

  // Initialize a new Validator instance. 
  // v := validator.New()

  // Call the ValidateMovie() function and return a response containing the errors if
  // any of the checks fail.
  // if data.ValidateMovie(v, movie); !v.Valid() {
  //   app.failedValidationResponse(w, r, v.Errors)
  //   return
  // }

  fmt.Fprintf(w, "%+v\n", input)
}
