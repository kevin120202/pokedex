# Pokedex Go CLI

A command-line interface (CLI) application built in Go that allows users to explore the Pokemon world, catch Pokemon, and manage their Pokedex. This application interacts with the [PokeAPI](https://pokeapi.co/) to provide real Pokemon data in an interactive terminal experience.

## 🚀 Features

- **Interactive CLI**: Command-line interface with command history and auto-completion
- **Location Exploration**: Browse Pokemon locations with pagination support
- **Pokemon Catching**: Attempt to catch Pokemon with realistic catch mechanics
- **Pokedex Management**: View and inspect caught Pokemon
- **API Caching**: Intelligent caching system to reduce API calls and improve performance
- **Error Handling**: Robust error handling for network issues and invalid inputs

## 📋 Prerequisites

- Go 1.24.2 or higher
- Internet connection (for API calls)

## 🛠️ Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/kevin120202/pokedex.git
   cd pokedex
   ```

2. **Build the application**:
   ```bash
   go build -o pokedex
   ```

3. **Run the application**:
   ```bash
   ./pokedex
   ```

## 🎮 Usage

Once you start the application, you'll see a prompt like `Pokedex >`. Here are the available commands:

### Core Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Display help information | `help` |
| `map` | Show the next page of locations | `map` |
| `mapb` | Show the previous page of locations | `mapb` |
| `explore <location>` | Explore a specific location to find Pokemon | `explore pallet-town` |
| `catch <pokemon>` | Attempt to catch a Pokemon | `catch pikachu` |
| `inspect <pokemon>` | View details of a caught Pokemon | `inspect pikachu` |
| `pokedex` | View all Pokemon in your Pokedex | `pokedex` |
| `exit` | Exit the application | `exit` |

### Example Session

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
explore <location_name>: Explore a location
catch: Attempt to catch a pokemon
inspect: Get details of a pokemon
pokedex: Look at your pokedex
exit: Exit the Pokedex

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
snowpoint-city-area
survival-area-area
resort-area-area
...

Pokedex > explore pallet-town
Exploring pallet-town...
Found Pokemon:
 - rattata
 - raticate
 - pidgey
 - pidgeotto

Pokedex > catch rattata
Throwing a Pokeball at rattata...
rattata was caught!

Pokedex > inspect rattata
Name: rattata
Height: 3
Weight: 35
Stats:
  -hp: 30
  -attack: 56
  -defense: 35
  -special-attack: 25
  -special-defense: 35
  -speed: 72
Types:
  - normal

Pokedex > pokedex
Your Pokedex:
 - rattata

Pokedex > exit
```

## 🏗️ Architecture

### Project Structure

```
pokedex/
├── main.go                 # Application entry point
├── repl.go                 # REPL (Read-Eval-Print Loop) implementation
├── command_*.go           # Individual command implementations
├── internal/
│   ├── pokeapi/           # PokeAPI client and data structures
│   │   ├── client.go      # HTTP client with caching
│   │   ├── location_*.go  # Location-related API calls
│   │   ├── pokemon_*.go   # Pokemon-related API calls
│   │   └── type_*.go      # Data type definitions
│   └── pokecache/         # Caching system
│       ├── pokecache.go   # Cache implementation
│       └── pokecache_test.go
├── go.mod                 # Go module definition
└── go.sum                 # Dependency checksums
```

### Key Components

#### 1. **REPL System** (`repl.go`)
- Interactive command-line interface
- Command parsing and execution
- History management using readline library

#### 2. **PokeAPI Client** (`internal/pokeapi/`)
- HTTP client with configurable timeouts
- Integration with caching system
- Methods for fetching locations, Pokemon data, and encounters

#### 3. **Caching System** (`internal/pokecache/`)
- Thread-safe cache implementation
- Automatic cleanup of expired entries
- Configurable cache duration

#### 4. **Command System**
Each command is implemented as a separate file:
- `command_map.go` - Location browsing with pagination
- `command_explore.go` - Location exploration
- `command_catch.go` - Pokemon catching mechanics
- `command_inspect.go` - Pokemon details display
- `command_pokedex.go` - Pokedex management
- `command_help.go` - Help system
- `command_exit.go` - Application exit

## 🔧 Configuration

The application uses the following default configurations:

- **HTTP Timeout**: 5 seconds for API requests
- **Cache Duration**: 5 minutes for cached responses
- **Catch Mechanics**: Based on Pokemon's base experience (lower XP = easier to catch)

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

---

**Happy Pokemon hunting! 🎮⚡** 