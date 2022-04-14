package main

import (
  "net/http"
)

// Note that the errors parameter here has the type map[string]string, which is exactly 
// the same as the errors map contained in our Validator type.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	// app.errorResponse(w, r, http.StatusUnprocessableEntity, errors) // TODO: failed 
}
