package data

import (
	"errors"
	"flag"
	"os"
)

type CommandLineInput struct {
	Command   string
	Args      []string
	UseStdin  bool
	Pick      bool
	Brands    string
	Issuers   string
	BrandName string
	Issuer    string
}

func ParseCommandLine(args []string) (*CommandLineInput, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("No command provided")
	}
	cmd := args[0]

	fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
	useStdin := fs.Bool("stdin", false, "read input from stdin")
	pick := fs.Bool("pick", false, "randomly pick one number")
	brands := fs.String("brands", "", "path to brands file")
	issuers := fs.String("issuers", "", "path to issuers file")
	brand := fs.String("brand", "", "brand name for issue command")
	issuer := fs.String("issuer", "", "issuer name for issue command")

	if err := fs.Parse(args[1:]); err != nil {
		return nil, err
	}
	return &CommandLineInput{
		Command:   cmd,
		Args:      fs.Args(),
		UseStdin:  *useStdin,
		Pick:      *pick,
		Brands:    *brands,
		Issuers:   *issuers,
		BrandName: *brand,
		Issuer:    *issuer,
	}, nil
}
