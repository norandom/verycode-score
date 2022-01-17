package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"

)

var (
	token, message			string
	dryRun, versionFlag     	bool
)

func main() {
	parseAndValidateInput()

	if versionFlag {
		printVersion()
		return
	}

	if dryRun {
		printOutput("dryRun: " + message)
		return
	}


	// code here
	printOutput("Result: " + message)
}

func parseAndValidateInput() {
	flag.StringVar(&message, "message", "", "i/o for dryRun")
	flag.StringVar(&token, "token", "123", "Personal access token, repo scope")
	flag.BoolVar(&dryRun, "dryRun", false, "if true or if env var DRY_RUN=true, then a tweet will not be sent")
	flag.BoolVar(&versionFlag, "version", false, "output the version of verycode")
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


func printVersion() {
	versionStr := "dirty"
	fmt.Printf("verycode version: %s", versionStr)
}

func printOutput(message string) {
	fmt.Printf("message: %s", message)
}
