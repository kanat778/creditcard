package data

import (
	"flag"
	"fmt"
	"os"
)

type AppArgs struct {
	Command string
	Inputs  []string
	Brands  string
	Issuers string
	Brand   string
	Issuer  string
	Stdin   bool
	Pick    bool
}

func ParseArgs() *AppArgs {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
		os.Exit(1)
	}

	app := &AppArgs{}
	app.Command = os.Args[1]
	
	fs := flag.NewFlagSet(app.Command, flag.ExitOnError)
	fs.BoolVar(&app.Stdin, "stdin", false, "read input from stdin")
	fs.BoolVar(&app.Pick, "pick", false, "pick random entry (for generate)")
	fs.StringVar(&app.Brands, "brands", "", "path to brands.txt")
	fs.StringVar(&app.Issuers, "issuers", "", "path to issuers.txt")
	fs.StringVar(&app.Brand, "brand", "", "brand name (for issue)")
	fs.StringVar(&app.Issuer, "issuer", "", "issuer name (for issue)")

	if err := fs.Parse(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing arguments:", err)
		os.Exit(1)
	}

	app.Inputs = fs.Args()
	return app
}

//func printUsage() {
//
//}
