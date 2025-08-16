package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %v", err)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %v", err)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

