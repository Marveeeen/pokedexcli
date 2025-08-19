package main

import (
	"fmt"
	"os"
	"math/rand"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
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

func commandMapf(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a location name to explore")
	}

	locationName := args[0]

	locationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a Pokémon name to catch")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res < 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil	
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a Pokémon name to inspect")
	}

	pokemonName := args[0]
	_, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("you haven't caught a Pokémon named %s", pokemonName)
	}

	fmt.Println("Name:", pokemonName)
	fmt.Println("Height:", cfg.caughtPokemon[pokemonName].Height)
	fmt.Println("Weight:", cfg.caughtPokemon[pokemonName].Weight)
	fmt.Println("Stats:")
	for _, stat := range cfg.caughtPokemon[pokemonName].Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range cfg.caughtPokemon[pokemonName].Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.caughtPokemon {
		fmt.Println(" -", name)
	}

	return nil
}

