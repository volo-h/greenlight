
The internal directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn’t application-specific and can potentially be reused will live in here. Our Go code under cmd/api will import the packages in the internal directory (but never the other way around).
Any packages which live under this directory can only be imported by code inside the parent of the internal directory. In our case, this means that any packages which live in internal can only be imported by code inside our greenlight project directory.
Any packages under internal cannot be imported by code outside of our project.

The remote directory will contain the configuration files and setup scripts for our production server.

The go.mod file will declare our project dependencies, versions and module path.

✗ curl -i localhost:4000/v1/healthcheck

/Users/lviv/Downloads5/scanlibs.com/greenlight

$ go run ./cmd/api -port=3030 -env=production

https://en.wikipedia.org/wiki/Representational_state_transfer

go get github.com/julienschmidt/httprouter@v1.3.0

curl -X POST localhost:4000/v1/movies

curl localhost:4000/v1/movies/123

curl -i -X OPTIONS localhost:4000/v1/healthcheck


> Additional Information 
Conflicting routes

if you do need to support conflicting routes (for example, you might need to replicate the endpoints of an existing API exactly for backwards-compatibility)
  - https://github.com/bmizerany/pat
  - https://github.com/go-chi/chi
  - https://github.com/gorilla/mux

httprouter - https://pkg.go.dev/github.com/julienschmidt/httprouter#Router

How to encode native Go objects into JSON using the https://pkg.go.dev/encoding/json package

Different techniques for customizing how Go objects are encoded to JSON — first by using struct tags, and then by leveraging the json.Marshaler interface. https://pkg.go.dev/encoding/json#Marshaler

call the json.Marshal() function
json.Encoder type


curl -i localhost:4000/v1/healthcheck

interface{} - empty interface - https://www.alexedwards.net/blog/interfaces-explained

func Marshal(v interface{}) ([]byte, error) - we’re able to pass any Go type to Marshal() as the v parameter.

Headers Location - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Location

[]byte slice
slice of structs

use json.Encoder type to perform the encoding

internal/data package - encapsulate all the custom data types for our project along with the logic for interacting with our database.


all the fields in our Movie struct are exported (i.e. start with a capital letter), which is necessary for them to be visible to Go’s encoding/json package. Any fields which aren’t exported won’t be included when encoding a struct to JSON.


you can customize the JSON by annotating the fields with struct tags. - https://github.com/golang/go/wiki/Well-known-struct-tags

JSON formats:
  https://jsonapi.org/
  https://github.com/omniti-labs/jsend

# Case Styles: Camel, Pascal, Snake, and Kebab Case
https://betterprogramming.pub/string-case-styles-camel-pascal-snake-and-kebab-case-981407998841


https://go.dev/doc/effective_go#methods
https://medium.com/globant/go-method-receiver-pointer-vs-value-ffc5ab7acdb

There are two approaches that you can take to decode JSON into a native Go object: using a json.Decoder type or using the json.Unmarshal() function.

