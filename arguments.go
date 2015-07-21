package mscore

import (
	"flag"
)

/*
Data type defining the possible command line arguments / flags
*/
type Arguments struct {
	Migrate bool
	Workers bool
}

/*
Parse and return the command line arguments / flags
*/
func ParseArguments() Arguments {
	migrate := flag.Bool("migrate", false,
		"Auto-migrate the database")
	workers := flag.Bool("workers", false,
		"Run the service workers")

	flag.Parse()

	return Arguments{
		Migrate: *migrate,
		Workers: *workers,
	}
}
