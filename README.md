Pokedex CLI

A command-line Pokedex application built in Go. This tool allows you to explore the Pokemon world, discover what Pokemon live in different areas, and try to catch them to build up your own personal Pokedex!

Data is dynamically fetched from the PokeAPI and cached locally for a faster user experience.

Features

Interactive REPL: A continuous command-line interface to explore the Pokemon world.

World Exploration: Paginate through 20 locations at a time, moving forward and backward.

Location Inspection: See exactly which Pokemon can be found in a specific location area.

Catching Mechanics: Throw Pokeballs at Pokemon! Catch rates are calculated dynamically based on the Pokemon's base experience.

Personal Pokedex: Keep track of the Pokemon you've successfully caught.

Pokemon Inspection: View detailed stats, types, height, and weight of Pokemon you have caught.

Installation and Usage

To run this project, you will need to have Go installed on your machine.

Clone the repository:

git clone [https://github.com/YOUR_USERNAME/pokedexcli.git](https://github.com/YOUR_USERNAME/pokedexcli.git)
cd pokedexcli


Build the executable:

go build


Run the Pokedex:

./pokedexcli


Available Commands

Once the Pokedex is running, you can use the following commands:

help - Displays a help message showing available commands.

exit - Closes the Pokedex CLI.

map - Displays the next 20 location areas in the Pokemon world.

mapb - Displays the previous 20 location areas.

explore <location_name> - Shows a list of all Pokemon that can be found in the specified area.

Example: explore canalave-city-area

catch <pokemon_name> - Throws a Pokeball at a Pokemon. If caught, it is added to your Pokedex!

Example: catch pikachu

inspect <pokemon_name> - View the stats, types, height, and weight of a Pokemon you have caught.

Example: inspect pikachu

pokedex - Prints a list of all the Pokemon you have successfully caught.
