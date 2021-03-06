
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

There are two approaches that you can take to decode JSON into a native Go object: 
using a json.Decoder type 
or
using the json.Unmarshal() function.

# run project
go run ./cmd/api

BODY='{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}'
curl -i -d "$BODY" localhost:4000/v1/movies

https://pkg.go.dev/encoding/json#Decoder.DisallowUnknownFields

When Go is decoding some JSON, it will check to see if the destination type satisfies the json.Unmarshaler interface.
If it does satisfy the interface, then Go will call it’s UnmarshalJSON() method to determine how to decode the provided JSON into the target type.

# Postgresql
  https://www.enterprisedb.com/postgres-tutorials/how-tune-postgresql-memory
  https://pgtune.leopard.in.ua/


go get github.com/lib/pq@v1.10.0


File: $HOME/.profile

export GREENLIGHT_DB_DSN='postgres://postgres:@localhost/greenlight?sslmode=disable'

$ source $HOME/.profile
$ echo $GREENLIGHT_DB_DSN

$ psql $GREENLIGHT_DB_DSN

By default PostgreSQL has a hard limit of 100 open connections and, if this hard limit is hit under heavy load, it will cause our pq driver to return a "sorry, too many clients already" error.
	The hard limit on open connections can be changed in your postgresql.conf file using the max_connections setting.

SetMaxOpenConns()
SetMaxIdleConns()	#	By default, the maximum number of idle connections is 2.

by default MySQL will automatically close any connections which haven’t been used for 8 hours.

the MaxIdleConns limit should always be less than or equal to MaxOpenConns

SetConnMaxLifetime() method sets the ConnMaxLifetime limit — the maximum length of time that a connection can be reused for


https://github.com/golang-migrate/migrate
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

brew install golang-migrate
$ migrate -version

✗ migrate create -seq -ext=.sql -dir=./migrations create_movies_table
✗ migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints

✗ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

SELECT * FROM schema_migrations;

# show structure of table in pg 
greenlight-> \d movies

$ migrate -path=./migrations -database=$EXAMPLE_DSN version
$ migrate -path=./migrations -database=$EXAMPLE_DSN goto 1
$ migrate -path=./migrations -database =$EXAMPLE_DSN down 1
$ migrate -path=./migrations -database=$EXAMPLE_DSN down

# to force the database version number to 1
  $ migrate -path=./migrations -database=$EXAMPLE_DSN force 1

# Decoupling database migrations from server startup: why and how
  https://pythonspeed.com/articles/schema-migrations-server-startup/ 

If you don’t like the term model then you might want to think of this as your data access or storage layer instead.

https://github.com/avelino/awesome-go#orm
https://github.com/avelino/awesome-go#database

https://pkg.go.dev/database/sql#DB.Exec
https://pkg.go.dev/database/sql#DB.QueryRow

Behind the scenes, the pq.Array() adapter takes our []string slice and converts it to a pq.StringArray type.

BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["action","adventure"]}'
curl -d "$BODY" localhost:4000/v1/movies

https://www.postgresql.org/docs/current/xfunc-sql.html


BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
curl -X PUT -d "$BODY" localhost:4000/v1/movies/2


# for DELETE use https://pkg.go.dev/database/sql#DB.Exec
One of the nice things about Exec() is that it returns a sql.Result object, which contains information about the number of rows that the query affected. In our scenario here, this is really useful information.


pointers have the zero-value nil

slices already have the zero-value nil

curl -X PATCH -d '{"year": 1985}' localhost:4000/v1/movies/4
curl -X PATCH -d '{"year": 1985, "title": ""}' localhost:4000/v1/movies/2


https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH

# About race condition
  https://stackoverflow.com/questions/34510/what-is-a-race-condition

http.Server handles each HTTP request in its own goroutine, so when this happens the code in our  updateMovieHandler will be running concurrently in two different goroutines.


# data race
This specific type of race condition is known as a data race. Data races can occur when two or more goroutines try to use a piece of shared data (in this example the movie record) at the same time, but the result of their operations is dependent on the exact order that the scheduler executes their instructions.

# optimistic locking
  https://stackoverflow.com/questions/129329/optimistic-vs-pessimistic-locking/129397#129397

This means that the first update request that reaches our database will succeed, and whoever is making the second update will receive an error message instead of having their change applied.


curl -i -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/2" & curl -i -X PATCH -d '{"genres": ["comedy","drama"]}' "localhost:4000/v1/movies/2"

UPDATE movies SET title = $1, year = $2, runtime = $3, genres = $4, version = uuid_generate_v4() WHERE id = $5 AND version = $6 RETURNING version


https://blog.josephscott.org/2011/10/14/timing-details-with-curl/

https://pkg.go.dev/context#Background


After 3 seconds, the context timeout is reached and our pq database driver sends a cancellation signal to PostgreSQL† .
PostgreSQL then terminates the running query, the corresponding resources are freed-up, and it returns the error message that we see above.
The client is then sent a 500 Internal Server Error response, and the error is logged so that we know something has gone wrong.

More precisely, our context (the one with the 3-second timeout) has a Done channel, and when the timeout is reached the Done channel will be closed. While the SQL query is running, our database driver pq is also running a background goroutine which listens on this Done channel. If the channel gets closed, then pq sends a cancellation signal to PostgreSQL. PostgreSQL terminates the query, and then sends the error message that we see above as a response to the original pq goroutine. That error message is then returned to our database model’s Get() method.

In fact, we can demonstrate this in our application by setting the maximum open connections to 1 and making two concurrent requests to our endpoint.

$ go run ./cmd/api -db-max-open-conns=1

# quering / pagination
We can retrieve the query string data from a request by calling the r.URL.Query() method. This returns a url.Values type, which is basically a map holding the query string data.
We can then extract values from this map using the Get() method, which will return the value for a specific key as a string type, or the empty string "" if no matching key exists in the query string.


curl "localhost:4000/v1/movies?title=godfather&genres=crime,drama&page=1&page_size=5&sort=year"

https://pkg.go.dev/database/sql#DB.QueryContext

curl localhost:4000/v1/movies

# reductive filter
https://ux.stackexchange.com/questions/88993/inclusive-additive-vs-exclusive-reductive-filtering-how-to-differentiate


SELECT id, created_at, title, year, runtime, genres, version FROM movies
  WHERE (LOWER(title) = LOWER($1) OR $1 = '')
  AND (genres @> $2 OR $2 = '{}') ORDER BY id

The @> symbol is the ‘contains’ operator for PostgreSQL arrays, and this condition will return true if all values in the placeholder parameter $2 are contained in the database genres field or the placeholder parameter contains an empty array.


# postgresql functions
  https://www.postgresql.org/docs/9.6/functions-array.html

PostgreSQL also provides a range of other useful array operators and functions, including the && ‘overlap’ operator, the <@ ‘contained by’ operator, and the array_length() function.


curl "localhost:4000/v1/movies?genres=adventure"
curl "localhost:4000/v1/movies?title=moana&genres=animation,adventure"
curl "localhost:4000/v1/movies?title=moana&genres=aaction,adventure"
curl "localhost:4000/v1/movies?title=moana&genres=action,adventure"
curl "localhost:4000/v1/movies?genres=action,adventure"

# about indexes
  https://www.postgresql.org/docs/13/indexes-intro.html
  https://www.postgresql.org/docs/13/indexes-types.html

# generate indexes and apply its 
  $ migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
  $ migrate -path ./migrations -database $GREENLIGHT_DB_DSN up

retrieve a list of all available configurations by running the \dF

if you wanted to use the english configuration to search our movies, you could update the SQL query like so:
  SELECT id, created_at, title, year, runtime, genres, version FROM movies WHERE (to_tsvector('english', title) @@ plainto_tsquery('english', $1) OR $1 = '') AND (genres @> $2 OR $2 = '{}') ORDER BY id

https://www.compose.com/articles/indexing-for-full-text-search-in-postgresql/
https://www.postgresql.org/docs/current/textsearch.html

The PostgreSQL STRPOS() function allows you to check for the existence of a substring in a particular database field.
https://www.postgresql.org/docs/current/functions-string.html

https://niallburkley.com/blog/index-columns-for-like-in-postgres/

# order by
  https://www.postgresql.org/docs/current/queries-order.html

The actual order in that case will depend on the scan and join plan types and the order on disk, but it must not be relied on.

# sort by year
  ✗ curl "localhost:4000/v1/movies?sort=-year"

# pagination
The LIMIT clause allows you to set the maximum number of records that a SQL query should return, 
and
OFFSET allows you to ‘skip’ a specific number of rows before starting to return records from the query.

LIMIT = page_size
OFFSET = (page - 1) * page_size

curl "localhost:4000/v1/movies?page_size=1"

# window function - counts the total number of filtered rows
  https://www.postgresql.org/docs/current/tutorial-window.html

curl "localhost:4000/v1/movies?page=1&page_size=2"

# mutexes
  https://www.alexedwards.net/blog/understanding-mutexes

# standard logger
  https://pkg.go.dev/log#pkg-overview

# another loggers
  https://github.com/avelino/awesome-go#logging
  https://github.com/rs/zerolog


for i in {1..6}; do curl http://localhost:4000/v1/healthcheck; done

using HAProxy 
or
Nginx as a load balancer 
or reverse proxy


go run ./cmd/api/ -limiter-burst=2
go run ./cmd/api/ -limiter-enabled=false

# graceful shutdown
https://pkg.go.dev/net/http#Server.Shutdown

Ctrl+C == interrupt signal == SIGINT signal

SIGINT          Ctrl+C
SIGQUIT         Ctrl+\
SIGKILL
SIGTERM


Catachable signals can be intercepted by our application and either ignored, or used to trigger a certain action (such as a graceful shutdown):
  SIGINT
  SIGQUIT

pgrep command to verify that this process exists
  pgrep -l api

# try sending a SIGKILL / SIGTERM signal to the api process
  pkill -SIGKILL api
  pkill -SIGTERM api

tools in the os/signals package that we can use to intercept catchable signals and trigger a graceful shutdown of our application.

https://gobyexample.com/channel-buffering

# sources!
  https://github.com/golang/go/blob/bc7e4d9257693413d57ad467814ab71f1585a155/src/os/signal/signal.go#L243

curl localhost:4000/v1/healthcheck & pkill -SIGTERM api

migrate create -seq -ext=.sql -dir=./migrations create_users_table

# register new user
  BODY='{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
  curl -i -d "$BODY" localhost:4000/v1/users


  BODY='{"name": "", "email": "bob@invalid.", "password": "pass"}'
  curl -d "$BODY" localhost:4000/v1/users

https://pkg.go.dev/golang.org/x/crypto/bcrypt#GenerateFromPassword
https://pkg.go.dev/golang.org/x/crypto/bcrypt#CompareHashAndPassword

case-insensitive
  alice@example.com
  ==
  alice@EXAMPLE.COM

SMTP service provider
  Postmark
  Sendgrid
  https://mailtrap.io/
  https://mailtrap.io/register/signup


embedded file systems


BODY='{"name": "Bob Jones", "email": "bob@example.com", "password": "pa55word"}'
curl -w '\nTime: %{time_total}\n' -d "$BODY" localhost:4000/v1/users


BODY='{"name": "Faith Smith", "email": "faith@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/users

# verificatrions process
  https://ux.stackexchange.com/questions/111005/what-is-the-point-of-email-verification/111023#111023

# ???
If a client sends the same PUT /v1/users/activated request multiple times, the first will succeed (assuming the token is valid) and then any subsequent requests will result in an error being sent to the client (because the token has been used and deleted from the database). But the important thing is that nothing in our application state (i.e. database) changes after that first request.
Basically, there are no application state side-effects from the client sending the same request multiple times, which means that the endpoint is idempotent and using PUT is more appropriate than POST .

curl -X PUT -d '{"token": "invalid"}' localhost:4000/v1/users/activated
curl -X PUT -d '{"token": "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}' localhost:4000/v1/users/activated

curl -X PUT -d '{"token": "GAHSM6PPCROP4VSSEBCGEU2XEU"}' localhost:4000/v1/users/activated

#
  https://portswigger.net/web-security/host-header

# sql joins
  https://www.dataquest.io/blog/sql-joins-tutorial/

#
  https://shahjerry33.medium.com/password-reset-token-leak-via-referrer-2e622500c2c1

# First-Class Functions in Golang !
  https://levelup.gitconnected.com/first-class-functions-in-golang-ef2a5001bb4f

# generics
  https://levelup.gitconnected.com/generics-example-in-go-minmax-40593f0adaa8

https://go.dev/play/p/eTz1xBm4W2a


Authentication is about confirming who a user is
Authorization is about checking whether that user is permitted to do something

five approaches:
  HTTP Basic authentication - 
    in header ex.: Authorization: Basic YWxpY2VAZXhhbXBsZS5jb206cGE1NXdvcmQ=
    Request.BasicAuth() https://pkg.go.dev/net/http#Request.BasicAuth
  Stateful token authentication - stored server-side in a database, alongside the user ID and an expiry time for the token.
    in header ex.: Authorization: Bearer <token>
  Stateless token authentication - stateless tokens encode the user ID and expiry time in the token itself. The token is cryptographically signed to prevent tampering and (in some cases) encrypted to prevent the contents being read.
  API key authentication. In an emergency, you could effectively revoke all tokens by changing the secret used for signing your tokens (forcing all users to re-authenticate), or another workaround is to maintain a blocklist of revoked tokens in a database (although that defeats the ‘stateless’ aspect of having stateless tokens).
    JWT
    https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/
    https://curity.io/resources/learn/jwt-best-practices/
  OAuth 2.0 / OpenID Connect

  API-key authentication - The idea behind API-key authentication is that a user has a non-expiring secret ‘key’ associated with their account. This key should be a high-entropy cryptographically-secure random string, and a fast hash of the key (SHA256 or SHA512) should be stored alongside the corresponding user ID in your database.
  OAuth 2.0 is not an authentication protocol, and you shouldn’t really use it for authenticating users - https://oauth.net/articles/authentication/ю
  If you want to implement authentication checks against a third-party identity provider, you should use OpenID Connect - https://openid.net/connect/
    in header ex: Authorization: Key <key>
    https://github.com/coreos/go-oidc

    when your API has a microservice architecture with different services for performing authentication and performing other tasks, then use stateless authentication tokens.


SELECT * FROM tokens WHERE scope = 'authentication';

Authorization is a request header, not a response header - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization

401 Unauthorized response should be used when you have missing or bad authentication
403 Forbidden response should be used afterwards, when the user is authenticated but isn’t allowed to perform the requested operation.



The difference between simple and preflight cross-origin requests.

scheme    https|https  
host      foo.com
port      443

same origin

# different path
  http://foo.com/a & http://foo.com/b
# different query string
  http://foo.com/a & http://foo.com/a?b=c
# different fragment
  http://foo.com/a#b & http://foo.com/a#c

A webpage on one origin can send data to a different origin. For example, it’s OK for a HTML form in a webpage to submit data to a different origin.
But a webpage on one origin is not allowed to receive data from a different origin.


What is the ratio of successful HTTP responses to both client and server errors?

https://pkg.go.dev/expvar

go mod tidy - command to prune any unused dependencies from the go.mod and go.sum files, and add any missing dependencies.

go mod verify - command to check that the dependencies on your computer (located in your module cache located at $GOPATH/pkg/mod ) haven’t been changed since they were downloaded and that they match the cryptographic hashes in your go.sum file. Running this helps ensure that the dependencies being used are the exact ones that you expect.

go fmt ./... command to format all .go files in the project directory, according to the Go standard. This will reformat files ‘in place’ and output the names of any changed files.

go vet ./... command to check all .go files in the project directory. The go vet tool runs a variety of analyzers which carry out static analysis of your code and warn you about things which might be wrong but won’t be picked up by the compiler such as unreachable code, unnecessary assignments, and badly-formed build tags.

go test -race -vet=off ./... command to run all tests in the project directory. By default, go test automatically executes a small subset of the go vet checks before running any tests, so to avoid duplication we’ll use the -vet=off flag to turn this off. The -race flag enables Go’s race detector, which can help pick up certain classes of race conditions while tests are running.

https://staticcheck.io/ - is a state of the art linter for the Go programming language. Using static analysis, it finds bugs and performance issues, offers simplifications, and enforces style rules.
  go install honnef.co/go/tools/cmd/staticcheck@latest


1. module proxies (also known as module mirrors)
  go env

  go env | grep GOPROXY

  export GOPROXY=https://goproxy.io,https://proxy.golang.org,direct

  # disable module mirrors altogether
    export GOPROXY=direct

go mod vendor - Vendoring dependencies in this way basically stores a complete copy of the source code for third-party packages in a vendor folder in your project.

go mod tidy - command will make sure the go.mod and go.sum files list all the necessary dependencies for our project (and no unnecessary ones).

go mod verify - command will verify that the dependencies stored in your module cache (located on your machine at $GOPATH/pkg/mod ) match the cryptographic hashes in the go.sum file.

go mod vendor - command will then copy the necessary source code from your module cache into a new vendor directory in your project root.

tree -L 3 ./vendor/

when you run a command such as 
  go run , 
  go test or 
  go build , 

the go tool will recognize the presence of a vendor folder and the dependency code in the vendor folder will be used — rather than the code in the module cache on your local machine.

# to remove everything from your local module cache
go clean -modcache

vendor/modules.txt file is essentially a manifest of the vendored packages and their version numbers. When vendoring is being used, the go tool will check that the module version numbers in modules.txt are consistent with the version numbers in the go.mod file. If there’s any inconsistency, then the go tool will report an error.


go mod verify - will verify that the dependencies in your module cache match the go.sum file
go mod vendor - will copy those same dependencies from the module cache into your vendor folder.

go fmt ./...            the current directory
go vet ./...            all sub-directories
go test ./... .         excluding the vendor directory


go build -o=./bin/api ./cmd/api

# debugger
  https://github.com/go-delve/delve

  https://www.sourceware.org/gdb/

  https://github.com/golang/go/issues/26074

go tool dist list

you can specify the operating system and architecture that you want to create the binary for by setting GOOS and GOARCH environment variables when running go build . For example:
$ GOOS=linux GOARCH=amd64 go build {args}

# where your build cache is
  go env GOCACHE

  /Users/lviv/Library/Caches/go-build

$ go build -a -o=/bin/foo ./cmd/foo        # Force all packages to be rebuilt 
$ go clean -cache                          # Remove everything from the build cache

https://www.digitalocean.com/pricing
https://caddyserver.com/
https://letsencrypt.org/

virtual machine known in Digital Ocean terminology as a droplet.
https://www.digitalocean.com/products/droplets

https://cloud.digitalocean.com/registrations/new

As part of the registration process, you’ll be prompted confirm your email address and then add a minimum of $5 USD pre-paid credit to your account using either a credit/debit card or PayPal.


ssh-keygen -t rsa -b 4096 -C "greenlight@greenlight.alexedwards.net" -f $HOME/.ssh/id_rsa_greenlight
  $HOME/.ssh/id_rsa_greenlight file contains your private key. Make sure to keep this secure, because anyone who has access to it will be able to impersonate you.
  
  $HOME/.ssh/id_rsa_greenlight.pub file contains your public key. We’ll upload a copy of this public key to Digital Ocean.

# SSH key listed in the output
  ssh-add -l

# add key to SSH agent
  ssh-add $HOME/.ssh/id_rsa_greenlight

add pub key from $HOME/.ssh/id_rsa_greenlight.pub TO 'popup window'

# 

# create droplet
  https://docs.digitalocean.com/reference/api/api-reference/
  https://www.digitalocean.com/community/tutorials/how-to-use-doctl-the-official-digitalocean-command-line-client

https://docs.digitalocean.com/products/billing/bandwidth/

# automated backup
https://www.digitalocean.com/docs/images/backups/

$ ssh root@45.55.49.87

# add the greenlight user to the sudo group, so that it can perform actions as root if necessary.

https://www.fail2ban.org/wiki/index.php/Main_Page
https://github.com/golang-migrate/migrate
https://caddyserver.com/docs/install#debian-ubuntu-raspbian

$ mkdir -p remote/setup 
$ touch remote/setup/01.sh

# run in production
./api -port=4000 -db-dsn=$GREENLIGHT_DB_DSN -env=production

http://<your_droplet_ip>:4000/v1/healthcheck

/home/greenlight/api -port=4000 -db-dsn=$GREENLIGHT_DB_DSN -env=production


https://www.freedesktop.org/wiki/Software/systemd/

https://www.freedesktop.org/software/systemd/man/systemd.unit.html

https://en.wikipedia.org/wiki/Public-key_cryptography
https://www.digitalocean.com/community/tutorials/ssh-essentials-working-with-ssh-servers-clients-and-keys
https://www.ssh.com/academy/ssh/keygen

ps -U greenlight

curl 45.55.49.87:4000/v1/healthcheck

# Disable port 4000
  ~$ sudo ufw delete allow 4000/tcp

sudo journalctl -u api

to configure Caddy to act as a reverse proxy and forward any HTTP requests that it receives onward to our API.

https://caddyserver.com/docs/caddyfile

web -> 80 -> Caddy (proxy) -> internal 4000 port


https://www.freenom.com/ru/index.html

Caddy will automatically handle provisioning and renewing TLS certificates for your domain via Let’s Encrypt, as well as redirecting all HTTP requests to HTTPS.

