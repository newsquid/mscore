package mscore

import (
	"flag"
)

/*
Data type defining the possible command line arguments / flags
*/
type Arguments struct {
	Migrate bool
}

/*
Parse and return the command line arguments / flags
*/
func ParseArguments() Arguments {
	migrate := flag.Bool("migrate", false,
		"Auto-migrate the database")

	flag.Parse()

	return Arguments{
		Migrate: *migrate,
	}
}
