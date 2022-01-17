package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	
	"path/filepath"

)

var (
	policy				string					 	// security policy for scoring
	token				string 						// GH access token
	message				string						// for output handling
	dryRun, versionFlag     	bool						// for CI testing
)

func main() {
	parseAndValidateInput()

	if versionFlag {
		printVersion()
		return
	}

	// used in CI
	if dryRun {
		printOutput("dryRun: " + message)
		return
	}


	// scorer code here
	printOutput("Result: " + message)
}

func parseAndValidateInput() {
	flag.StringVar(&message, "message", "", "i/o for dry runs")
	flag.StringVar(&token, "token", "123", "GH personal access token, repo scope needed")
	flag.StringVar(&policy, "policy", "policies/standard.json", "policy for scoring")
	flag.BoolVar(&dryRun, "dryRun", false, "for testing only")
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
		
		policyPath := filepath.FromSlash(policy)
		fmt.Println(policyPath)

		
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
