# Pokedex CLI

A simple command-line interface (CLI) application for exploring Pokemon data using the PokeAPI.

## Requirements

- Go 1.18 or newer

## Installation

Clone the repository:

```
git clone https://github.com/marveeeen/pokedexcli.git
cd pokedexcli
```

Build the project:

```
go build -o pokedexcli
```

## Usage

Run the CLI:

```
./pokedexcli
```

At the prompt, type `help` to see all available commands and their descriptions.

## Project Structure

- `main.go` — Entry point
- `repl.go` — REPL logic and command handling
- `commands.go` — CLI command definitions
- `internal/pokeapi` — PokéAPI client and types
- `internal/pokecache` — Caching logic

## Testing

Run unit tests with:

```
go test ./...
```
