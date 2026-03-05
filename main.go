package main

import (
	"bufio"
	"fmt"
	"os"
)
	
func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
	}

}


