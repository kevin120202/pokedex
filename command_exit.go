package main

import "os"

func exitCallback(config *config) error {
	os.Exit(0)
	return nil
}
