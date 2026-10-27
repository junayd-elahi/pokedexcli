# Pokedex CLI

An interactive command-line Pokedex in Go, backed by the [PokeAPI](https://pokeapi.co/). Browse location areas, see which Pokemon appear there, catch them, and inspect the ones you've caught.

## Features

- **REPL** reading commands from stdin, with case-insensitive input.
- **Paginated browsing** of location areas, forwards (`map`) and backwards (`mapb`).
- **In-memory response cache** (`internal/pokecache`): API responses are stored by URL and reaped by a background goroutine once they pass their TTL. A mutex guards concurrent access between lookups and the reaper.
- **Catch mechanic** weighted by each Pokemon's base experience, so stronger Pokemon are harder to catch.
- **Unit tests** for input parsing and the cache.

## Commands

| Command | Description |
|---|---|
| `help` | List all commands |
| `map` | Show the next page of location areas |
| `mapb` | Show the previous page of location areas |
| `explore <area>` | List the Pokemon found in an area |
| `catch <pokemon>` | Try to catch a Pokemon |
| `inspect <pokemon>` | Show stats and types for a caught Pokemon |
| `pokedex` | List every Pokemon you've caught |
| `exit` | Quit |

## Running it

Requires Go 1.25+.

```bash
go run .
```

Run the tests:

```bash
go test ./...
```

## Structure

```
.
├── main.go, repl.go        # entry point, REPL loop and command registry
├── command_*.go            # one file per command
└── internal/
    ├── pokeapi/            # HTTP client and response types
    └── pokecache/          # thread-safe TTL cache
```
