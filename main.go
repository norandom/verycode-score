package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"

)

var (
	token 			string
	dryRun, versionFlag     bool
)

func main() {
	parseAndValidateInput()

	if versionFlag {
		printVersion()
		return
	}

	if dryRun {
		printOutput("dryRun: ", "test")
		return
	}


	// code here
	printOutput("Result: ", "test")
}

func parseAndValidateInput() {
	flag.StringVar(&token, "GitHub personal access token", "", "Personal access token, repo scope")
	flag.BoolVar(&dryRun, "dryRun", false, "if true or if env var DRY_RUN=true, then a tweet will not be sent")
	flag.BoolVar(&versionFlag, "version", false, "output the version of tweeter")
	flag.Parse()

	if os.Getenv("DRY_RUN") == "true" {
		dryRun = true
	}

	if versionFlag {
		return
	}

	var err error

	if !dryRun {
		if token == "" {
			err = multierror.Append(err, errors.New("--token can't be empty"))
		}
	}

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
