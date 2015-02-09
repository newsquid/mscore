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
- Parsing of a number of default command line arguments
- Database connection based on environment variables
- Database interaction through the Gorm ORM
- Database migration based on the auto-migration of the Gorm ORM
- Unit testing based on the Testify test framework
- Integration testing based on the Testify test framework and on the Martini HTTP framework

## JSON APIs
JSON API creation is done via the Martini framework, including routing and
middleware. The framework gives direct access to the underlying Martini
structs. In addition, it handles:

- Creating a connection to the database and mapping the connection as middleware
- Adding a return handler, which automatically converts returned HTTP status
codes and structs to HTTP responses containing JSON data


## Command line argument parsing
The ˋParseArguments()ˋ function parses the command line arguments to an instance of the
follow struct:

ˋˋˋgo
type Arguments struct {
	Migrate bool // --migrate
}
ˋˋˋ


## Database connection
A database connection is established with the ˋInitDB()ˋ function. It takes no
arguments, but establishes the connection based on the followin environment
variables:

- ˋDBˋ: The database type ("sqlite3", "mysql", etc.)
- ˋDBCONNˋ: The connection string
- ˋDBDEBUGˋ: Whether or not the connection is in debug mode (true/false)

The values default to "sqlite3", "database.db", false.


## Database interaction and migration
Database interaction and migration is done directly on the Gorm database
connection. Martini middleware ensures that the connection is available to the
endpoints.


## Unit and integration testing
For testing, a test struct must be defined. MSCore provides helper methods for setup and teardown. See the example for more details.


# Usage example
Examples of an API as well as unit and integration tests are given in the
'examples' directory.
