package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/saltside/valiant/config"
)

var (
	errFileExt     = errors.New("Unsupported file extension, supported: '.yml'")
	errTestFailure = errors.New("Failed test")
)

// runTests reads a directory and runs all the tests in that directory.
func runTests(address, testDir string) error {
	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		return err
	}

	gotErr := false

	for _, file := range files {
		spec, err := parseFile(filepath.Join(testDir, file.Name()))

		if err == errFileExt {
			continue
		}

		if err != nil {
			gotErr = true
			fmt.Println(err)
			continue
		}

		fmt.Printf("--- test: %s [%s %s]\n", file.Name(), spec.Request.Method, spec.Request.Path)

		err = runSingleTest(address, spec)
		if err == errTestFailure {
			gotErr = true
			continue
		}

		if err != nil {
			gotErr = true
			fmt.Println(err)
			continue
		}

		fmt.Println("ok")
	}

	if !gotErr {
		return nil
	}

	return errTestFailure
}

// runSingleTest runs a single test and returns an error if there are any.
func runSingleTest(address string, testSpec *config.TestSpec) error {
	resp, err := testSpec.SendRequest(address)
	if err != nil {
		return err
	}

	issues, err := testSpec.Validate(*resp)
	if err != nil {
		return err
	}

	for _, issue := range issues {
		fmt.Printf("%s\n", issue)
	}

	if len(issues) != 0 {
		return errTestFailure
	}

	return nil
}

// parseFile reads a file and returns the parsed content to the TestSpec
// struct.
func parseFile(fullPath string) (*config.TestSpec, error) {
	ext := filepath.Ext(fullPath)
	if ext != ".yml" {
		return nil, errFileExt
	}

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return config.Parse(data)
}
