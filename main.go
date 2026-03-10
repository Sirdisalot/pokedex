package main

import (
	"time"

	"github.com/sirdisalot/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

	/*
		for {
			fmt.Print("Pokedex >")
			scanner.Scan()
			text := scanner.Text()
			cleaned := cleanInput(text)
			if len(cleaned) == 0 {
				continue
			}

			commandName := cleaned[0]

			availabelCommands := getCommands()
			if command, ok := availabelCommands[commandName]; ok {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unkown command")
			}
		}*/
}
