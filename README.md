# Pokedex

![Pokedex_logo](https://github.com/user-attachments/assets/59a9afb0-eaa9-4a5d-9c8c-308f7a982d98)

Pokedex is a device tool that let us to look up information about Pokemon, eg. pokemon name, type and stats. 
It is built in a command-line REPL. We use the [Poke API](https://pokeapi.co/docs/v2#info) to fetch all the data we will need using GET request.

# Local Development
1) This is a simple Golang project, you will need to install the latest [Go toolchain](https://go.dev/doc/install).

2) Create a main.go file at the root of project.

3) Create a Go module int the root of project.

4) Build your project: (Alternative way to run and build your project, you can run (6) instead, ignore (4) and (5) steps.)
   ```
   go build
   ```

5) Run your program:
   ```
   ./pokedexcli
   ```

6) Build and run your program:
   ```
   ./scripts/build_run.sh
   ```

7) Run the tests
   ```
   go test ./...
   ```


# Commands
There are eight commands to use in Pokedex:
1) Help
    - It shows all the commands available in Pokedex. It is like pokedex menu.
    
      ![command-help](https://github.com/user-attachments/assets/052e37e7-e710-4764-b040-cff5efe45437)

2) Map
     - It displays next 20 locations in Pokemon world

      ![command-map](https://github.com/user-attachments/assets/049efa4a-b493-456f-aa96-d36ec2d48967)

3) Mapb
      - It displays previous 20 locations in Pokemon world
      
      ![command-mapb](https://github.com/user-attachments/assets/f6743b94-bbb5-4045-902a-eb51ae8d8c21)

4) Explore [location]
      - It displays the name of Pokemons found in that location

      ![command-explore](https://github.com/user-attachments/assets/68f16c48-e518-462a-b44b-5fac11e53e26)

5) Catch [Pokemon Name]
      - It helps user to catch pokemon, user can inspect pokemon details after caught them successfully.

      ![command-catch](https://github.com/user-attachments/assets/0ca70ccc-5148-4991-bf5a-52f6f1172511)

6) Inspect [Pokemon Name]
      - It shows the name, stats and type of a pokemon.

      ![command-inspect](https://github.com/user-attachments/assets/09f63377-0450-44fe-b38c-d014ab5cd663)

7) Pokedex
      
      - It shows all the pokemon name that user have caught.

      ![command-pokedex](https://github.com/user-attachments/assets/16d3de4e-386f-409b-ab02-2d6cd907e452)

8) Exit
      - It exits the CLI tool (Pokedex device)

      ![command-exit](https://github.com/user-attachments/assets/6f672a8a-1073-40bb-bf56-ac00022dd82a)

   


