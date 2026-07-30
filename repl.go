package main

import (
	"strings"
	"bufio"
	"fmt"
	"errors"
	"os"
	"math/rand"
	"github.com/JRPOGM/PokeDex/TheApi/pokeapi"
)

type regisCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

type config struct {
	pokeapiClient		pokeapi.Client
	nextLocationURL		*string
	previousLocationURL	*string
	registeredPokemon	map[string]pokeapi.Pokemon
}

func getCommands() map[string]regisCommand {
	return map[string]regisCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a message",
			callback:		commandHelp,
		},
		"map": {
			name: 			"map",
			description:	"Pulls up 20 named areas",
			callback: 		commandMap,	
		},
		"mapb": {
			name:			"map back",
			description:	"Displays previous 20 area names",
			callback:		commandMapB,
		},
		"explore": {
			name:			"explore",
			description:	"Displays all the Pokemon in an area",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Attempts to catch a Pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Prints information of a Pokemon",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Lists all registered Pokemon",
			callback:		commandPokedex,
		},
		"battle": {
			name:			"battle",
			description:	"Fight a wild Pokemon with a registered Pokemon",
			callback:		commandBattle,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	texting := strings.Fields(words)
	return texting
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		encounter := []string{}
		if len(words) > 1{
			encounter = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, encounter...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func commandExit(cfg *config, encounter ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, encounter ...string) error {
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

func commandMap(cfg *config, encounter ...string) error {
	localResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = localResponse.Next
	cfg.previousLocationURL = localResponse.Previous
	for _, local := range localResponse.Results {
		fmt.Println(local.Name)
	}
	return nil
}

func commandMapB(cfg *config, encounter ...string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("No prior locations to submit")
	}
	localResponse, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = localResponse.Next
	cfg.previousLocationURL = localResponse.Previous
	for _, local := range localResponse.Results {
		fmt.Println(local.Name)
	}
	return nil
}

func commandExplore(cfg *config, encounter ...string) error {
	if len(encounter) != 1 {
		return errors.New("No Pokemon encountered")
	}
	name := encounter[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Encountered Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, encounter ...string) error {
	if len(encounter) != 1 {
		return errors.New("No Pokemon encountered")
	}
	name := encounter[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	response := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if response > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("Pokemon data has now been registered")
	cfg.registeredPokemon[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, encounter ...string) error {
	if len(encounter) != 1 {
		return errors.New("No Pokemon encountered")
	}
	name := encounter[0]
	pokemon, ok := cfg.registeredPokemon[name]
	if !ok {
		return errors.New("No Pokemon of that name are registered")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeName := range pokemon.Types {
		fmt.Println(" -", typeName.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, encounter ...string) error {
	fmt.Println("Your Pokedex:")
	for _, captures := range cfg.registeredPokemon {
		fmt.Printf(" -%s\n", captures.Name)
	}
	return nil
}

func commandBattle(cfg *config, encounter ...string) error {
	fmt.Println("Battling wild Pokemon is currently under development.")
	return nil
}