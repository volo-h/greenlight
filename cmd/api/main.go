package main

import (
  "context" // New import
  "database/sql" // New import
	"flag"
	"os"
	"time"
  "greenlight.alexedwards.net/internal/data" // New import
  "greenlight.alexedwards.net/internal/jsonlog" // New import

  // Import the pq driver so that it can register itself with the database/sql 
  // package. Note that we alias this import to the blank identifier, to stop the Go 
  // compiler complaining that the package isn't being used.
  _ "github.com/lib/pq"
)

// Declare a string containing the application version number. Later in the book we'll 
// generate this automatically at build time, but for now we'll just store the version 
// number as a hard-coded global constant.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our application. 
// For now, the only configuration settings will be the network port that we want the 
// server to listen on, and the name of the current operating environment for the 
// application (development, staging, production, etc.). We will read in these 
// configuration settings from command-line flags when the application starts.

// Add a db struct field to hold the configuration settings for our database connection 
// pool. For now this only holds the DSN, which we will read in from a command-line flag. 
type config struct {
  port int
  env string
  db struct { 
    dsn string
    maxOpenConns int
    maxIdleConns int
    maxIdleTime string
  }
  // Add a new limiter struct containing fields for the requests-per-second and burst 
  // values, and a boolean field which we can use to enable/disable rate limiting 
  // altogether.
  limiter struct {
    rps float64
    burst int
    enabled bool
  }
}

// Define an application struct to hold the dependencies for our HTTP handlers, helpers, 
// and middleware. At the moment this only contains a copy of the config struct and a 
// logger, but it will grow to include a lot more as our build progresses.

type application struct {
	config config
	// logger *log.Logger
  logger *jsonlog.Logger
  models data.Models
}

func main() {
  var cfg config

  flag.IntVar(&cfg.port, "port", 4000, "API server port")
  flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

  // Read the DSN value from the db-dsn command-line flag into the config struct. We 
  // default to using our development DSN if no flag is provided.
  flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:@localhost/greenlight?sslmode=disable", "PostgreSQL DSN")  

  // TODO:
  // Use the value of the GREENLIGHT_DB_DSN environment variable as the default value // for our db-dsn command-line flag.
  // flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

  //  Read the connection pool settings from command-line flags into the config struct.
  // Notice the default values that we're using?
  flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
  flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
  flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

  // Create command line flags to read the setting values into the config struct.
  // Notice that we use true as the default for the 'enabled' setting? 
  flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second") 
  flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst") 
  flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

  flag.Parse()

  // logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  // Initialize a new jsonlog.Logger which writes any messages *at or above* the INFO
  // severity level to the standard out stream.
  logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

  // Call the openDB() helper function (see below) to create the connection pool, 
  // passing in the config struct. If this returns an error, we log it and exit the 
  // application immediately.

  db, err := openDB(cfg)
  if err != nil {
    // logger.Fatal(err)

    // Use the PrintFatal() method to write a log entry containing the error at the 
    // FATAL level and exit. We have no additional properties to include in the log 
    // entry, so we pass nil as the second parameter.
    logger.PrintFatal(err, nil)
  }

  // Defer a call to db.Close() so that the connection pool is closed before the 
  // main() function exits.
  defer db.Close()

  // Also log a message to say that the connection pool has been successfully
  // established.
  // logger.Printf("database connection pool established")
  // Likewise use the PrintInfo() method to write a message at the INFO level. 
  logger.PrintInfo("database connection pool established", nil)

  app := &application{ 
    config: cfg,
    logger: logger,
    models: data.NewModels(db),
  }

  // Call app.serve() to start the server. 
  err = app.serve()
  if err != nil {
    logger.PrintFatal(err, nil)
  }
}

// The openDB() function returns a sql.DB connection pool. 
func openDB(cfg config) (*sql.DB, error) {
  // Use sql.Open() to create an empty connection pool, using the DSN from the config 
  // struct.
  db, err := sql.Open("postgres", cfg.db.dsn)
  if err != nil { 
    return nil, err
  }

  // Set the maximum number of open (in-use + idle) connections in the pool. Note that 
  // passing a value less than or equal to 0 will mean there is no limit. 
  db.SetMaxOpenConns(cfg.db.maxOpenConns)

  // Set the maximum number of idle connections in the pool. Again, passing a value 
  // less than or equal to 0 will mean there is no limit. 
  db.SetMaxIdleConns(cfg.db.maxIdleConns)

  // Use the time.ParseDuration() function to convert the idle timeout duration string // to a time.Duration type.
  duration, err := time.ParseDuration(cfg.db.maxIdleTime)
  if err != nil {
    return nil, err
  }

  // Set the maximum idle timeout. 
  db.SetConnMaxIdleTime(duration)

  // Create a context with a 5-second timeout deadline.
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  // Use PingContext() to establish a new connection to the database, passing in the 
  // context we created above as a parameter. If the connection couldn't be 
  // established successfully within the 5 second deadline, then this will return an 
  // error.

  err = db.PingContext(ctx)
  if err != nil { 
    return nil, err
  }

  // Return the sql.DB connection pool.
  return db, nil
}
