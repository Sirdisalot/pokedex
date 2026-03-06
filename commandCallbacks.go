package main

import (
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	locatioResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locatioResp.Next
	cfg.prevLocationsURL = locatioResp.Previous

	for _, loc := range locatioResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("No previous page of locations")
		return nil
	}

	locatioResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locatioResp.Next
	cfg.prevLocationsURL = locatioResp.Previous

	for _, loc := range locatioResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
