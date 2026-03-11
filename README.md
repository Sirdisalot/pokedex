<div align="center">
<h1>🔴⚪ Pokedex CLI</h1>
<p><i>A command-line Pokedex application built in Go.</i></p>
</div>
This tool allows you to explore the Pokemon world, discover what Pokemon live in different areas, and try to catch them to build up your own personal Pokedex!
Data is dynamically fetched from the PokeAPI and cached locally for a blazingly fast user experience.
✨ Features
🧭 World Exploration: Paginate through 20 locations at a time, moving forward and backward.
🔍 Location Inspection: See exactly which Pokemon can be found in a specific location area.
🎯 Catching Mechanics: Throw Pokeballs at Pokemon! Catch rates are calculated dynamically based on the Pokemon's base experience.
📖 Personal Pokedex: Keep track of the Pokemon you've successfully caught.
📊 Pokemon Inspection: View detailed stats, types, height, and weight of Pokemon you have caught.
⚡ Custom Caching: Includes a custom, thread-safe, auto-reaping cache package that minimizes network requests and speeds up exploration.
🚀 Installation and Usage
To run this project, you will need to have Go installed on your machine.
1. Clone the repository:
git clone [https://github.com/YOUR_USERNAME/pokedexcli.git](https://github.com/YOUR_USERNAME/pokedexcli.git)
cd pokedexcli


2. Build the executable:
go build


3. Run the Pokedex:
./pokedexcli


🎮 Available Commands
Once the Pokedex is running, you can use the following commands in the REPL:
Command
Description
Example
help
Displays a help message showing available commands.
help
exit
Closes the Pokedex CLI.
exit
map
Displays the next 20 location areas in the Pokemon world.
map
mapb
Displays the previous 20 location areas.
mapb
explore <area>
Shows a list of all Pokemon that can be found in the specified area.
explore canalave-city-area
catch <pokemon>
Throws a Pokeball at a Pokemon. If caught, it is added to your Pokedex!
catch pikachu
inspect <pokemon>
View the stats, types, height, and weight of a caught Pokemon.
inspect pikachu
pokedex
Prints a list of all the Pokemon you have successfully caught.
pokedex

🛠️ Technologies Used
Go (Golang): Standard library only (no third-party packages used).
PokeAPI: RESTful API for all Pokemon data.
Concurrency: Goroutines and Mutexes used for background cache clearing.
🎓 Acknowledgments
This project was built as part of the backend curriculum on Boot.dev.
