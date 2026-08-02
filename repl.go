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
	experienceBar		int
	badgeCase			int
	victoryCase			int
	championCase		int
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
		"amie": {
			name:			"amie",
			description:	"Play with your registered Pokemon",
			callback:		commandAmie,
		},
		"experience": {
			name:			"experience",
			description:	"Displays how much experience has been gained so far",
			callback:		commandExperience,
		},
		"buybadge": {
			name:			"buy badge",
			description:	"Buy a Gym badge for 200 experience points each",
			callback:		commandBuyBadge,
		},
		"badgecase": {
			name:			"badge case",
			description:	"Look at the amount of badges obtained",
			callback:		commandBadgeCase,
		},
		"victorybadge": {
			name:			"victory badge",
			description:	"Buy a Victory Road badge for 500 experience points each",
			callback:		commandVictoryBadge,
		},
		"victorycase": {
			name:			"victory case",
			description:	"Look at the amount of victory badges obtained",
			callback:		commandVictoryCase,
		},
		"championbadge": {
			name:			"champion badge",
			description:	"Buy a Champion Badge for 1000 experience points each",
			callback:		commandChampionBadge,
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
	fmt.Println("30 exp gained!")
	fmt.Println("Pokemon data has now been registered")
	cfg.registeredPokemon[pokemon.Name] = pokemon
	cfg.experienceBar += 30
	if len(cfg.registeredPokemon) % 25 == 0 {
		fmt.Printf("Congratulations! You've managed to catch %v Pokemon! But there are many more still to catch. Don't stop traveling now!\n", len(cfg.registeredPokemon))
	}
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
	if len(encounter) != 1 {
		return errors.New("No Pokemon encountered")
	}
	name := encounter[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	response := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Approaching %s for a battle...\n", pokemon.Name)
	if len(cfg.registeredPokemon) == 0 {
		fmt.Println("No Pokemon in party. Unable to battle.")
	} else if len(cfg.registeredPokemon) != 0 && response > 60 {
		fmt.Printf("%s has won the battle. You run away in a hurry...\n", pokemon.Name)
	} else if len(cfg.registeredPokemon) != 0 {
		fmt.Printf("You have defeated %s! 50 exp gained!\n", pokemon.Name)
		cfg.experienceBar += 50
	}
	return nil
}

func commandAmie(cfg *config, encounter ...string) error {
	if len(encounter) != 1 {
		return errors.New("No Pokemon encountered")
	}
	name := encounter[0]
	pokemon, ok := cfg.registeredPokemon[name]
	if !ok {
		return errors.New("No Pokemon of that name are registered")
	}
	if len(pokemon.Name) > 10 {
		fmt.Printf("Your %s dances in a circle\n", pokemon.Name)
	} else if len(pokemon.Name) < 11 && len(pokemon.Name) > 7 {
		fmt.Printf("%s nuzzles against your hand\n", pokemon.Name)
	} else if len(pokemon.Name) < 8 && len(pokemon.Name) > 5 {
		fmt.Printf("%s wishes to spend more time with you\n", pokemon.Name)
	} else if len(pokemon.Name) < 6 {
		fmt.Printf("Your %s hungers for battle\n", pokemon.Name)
	}
	cfg.experienceBar += 10
	return nil
}

func commandExperience(cfg *config, encounter ...string) error {
	fmt.Printf("Experience gained: %v\n", cfg.experienceBar)
	return nil
}

func commandBuyBadge(cfg *config, encounter ...string) error {
	if cfg.experienceBar >= 200 {
		fmt.Println("Spending experience to buy a badge\n")
		fmt.Println("Congrats on obtaining a badge! Badge has been added to your Badge Case!\n")
		cfg.experienceBar -= 200
		cfg.badgeCase++
	} else {
		fmt.Printf("Not enough experience to obtain badge. Current experience: %v\n", cfg.experienceBar)
	}
	if cfg.badgeCase % 8 == 0 {
		fmt.Printf("You've obtained %v badges! You are now able to obtain a set of 4 Victory badges!\n", cfg.badgeCase)
	}
	return nil
}
func commandBadgeCase(cfg *config, encounter ...string) error {
	fmt.Printf("Badges obtained: %v\n", cfg.badgeCase)
	return nil
}

func commandVictoryBadge(cfg *config, encounter ...string) error {
	if cfg.badgeCase % 8 == 0 {
		if cfg.experienceBar >= 500 {
			fmt.Println("Spending experience to buy a badge\n")
			fmt.Println("Congrats on obtaining a Victory Badge! The Victory Badge has been added to your Victory Case!\n")
			cfg.experienceBar -= 500
			cfg.victoryCase++
		} else {
			fmt.Printf("Not enough experience or badges to obtain a Victory Badge. Current experience: %v\n", cfg.experienceBar)
		}
	}
	if cfg.victoryCase % 4 == 0 {
		fmt.Println("You've obtained %v Victory Badges! You are now able to obtain a Champion Badge!\n", cfg.victoryCase)
	}
	return nil
}

func commandVictoryCase(cfg *config, encounter ...string) error {
	fmt.Printf("Victory Badges obtained: %v\n", cfg.victoryCase)
	return nil
}

func commandChampionBadge(cfg *config, encounter ...string) error {
	fmt.Println("Development on the Champion badge requirements are currently under production.")
	return nil
}