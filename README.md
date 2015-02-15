# Introduction
MSCore is a simple micro service framework written in Go. It has the following dependencies:

- https://github.com/go-martini/martini
- https://github.com/jinzhu/gorm
- https://github.com/stretchr/testify

The project currently contains no executable code (including tests). It can
therefore only be tested through services making use of the framework.


# Functionality
The framework contains utilities for:

- JSON APIs based on the Martini HTTP framework
- Some Martini middle ware for parsing useful parameters from requests
- Parsing of a number of default command line arguments
- Database connection based on environment variables
- Database interaction through the Gorm ORM
- Database migration based on the auto-migration of the Gorm ORM
- Handling background workers
- Unit testing based on the Testify test framework
- Integration testing based on the Testify test framework and on the Martini HTTP framework

## JSON APIs
JSON API creation is done via the Martini framework, including routing and
middleware. The framework gives direct access to the underlying Martini
structs. In addition, it handles:

- Creating a connection to the database and mapping the connection as middleware
- Adding a return handler, which automatically converts returned HTTP status
codes and structs to HTTP responses containing JSON data

##Parameter middle ware 
MSCore currently includes 3 middle ware handlers for parsing often used request
parameters (available in the parameters.go file):

- `Pagination` parses the query parameters `page` and `items_per_page` from a
  request and saves them in the `PaginationParameters` struct, that also
  contains the helpful `Offset()` method for calculating the current item
  offset. Default values are 1 and 9 respectively if they are not set in the
  request. Example request: `http://example.com/?page=2&items_per_page`.
- `ResourceId` parses the `:id` parameter from a martini route, and fails with
  a http 422 status code if the id is not set. Example route: `/:id/info`.
- `SearchTerm` parses the `search` query parameter from a request at save it in a
  `SearchParameter` struct, if `search` is not set in the request
  `SearchParameter.Success` will be false. `SearchParameter` also contains the
  helpful `SearchTerm()` method that will return the `%Search%` useful for
  database querying. Example request: `http://example.com/?search=test`


## Command line argument parsing
The `ParseArguments()` function parses the command line arguments to an instance of the
following struct:

```go
type Arguments struct {
	Migrate bool // --migrate
}
```

## Database connection
A database connection is established with the `InitDB()` function. It takes no
arguments, but establishes the connection based on the following environment
variables:

- `DB`: The database type ("sqlite3", "mysql", etc.)
- `DBCONN`: The connection string
- `DBDEBUG`: Whether or not the connection is in debug mode (true/false)

The values default to "sqlite3", "database.db", false.


## Database interaction and migration
Database interaction and migration is done directly on the Gorm database
connection. Martini middleware ensures that the connection is available to the
endpoints.

##Handling background workers
Background workers can be used to have a routine of work performed a regular
intervals. A worker is defined by conforming to the `Worker` interface which has
a single method `Routine()` representing the routine of work that has to be
performed. After defining a worker and implementing its routine, we simply start
it using a `Handler` giving it some `Interval` between each routine of work:

```go
worker := YourWorker{}
handler := mscore.Handler{
    Worker: worker,
    Interval: time.Minute * 5,
}
handler.Start()
```

## Unit and integration testing
For testing, a test struct must be defined. MSCore provides helper methods for setup and teardown. See the example for more details.


# Usage example
Examples of an API as well as unit and integration tests are given in the
'examples' directory.
