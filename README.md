# Go WebSocket Client for Hackathon 2024

This Go-based WebSocket client was developed for the Hackathon 2024, organized
by WULS-SGGW. It serves as a framework for participants to create AI agents that
can play the game.

To fully test and run the game, you will also need the game server and GUI
client, as the GUI provides a visual representation of gameplay. You can find
more information about the server and GUI client in the following repository:

- [Server and GUI Client Repository](https://github.com/INIT-SGGW/HackArena2024H2-Game)

## Development

The agent logic you are going to implement is located in `agent/agent.go`:

```go
// Agent represents an AI player in the game.
type Agent struct {
	MyID string
}

// OnJoiningLobby is called when the agent joins a lobby, creating a new instance of the agent.
// This method initializes the agent with the lobby's current state and other relevant details.
//
// Parameters:
//   - lobbyData: The initial state of the lobby when the agent joins.
//     Contains information like player data, game settings, etc.
//
// Returns:
// - A new instance of the agent.
func OnJoiningLobby(lobbyData *lobby_data.LobbyData) *Agent {
	return &Agent{
		MyID: lobbyData.PlayerID,
	}
}

// OnLobbyDataChanged is called whenever there is a change in the lobby data.
// This method is triggered under various circumstances, such as:
// - When a player joins or leaves the lobby.
// - When server-side game settings are updated.
//
// Parameters:
//   - lobbyData: The updated state of the lobby, containing information
//     like player details, game configurations, and other relevant data.
//     This is the same data structure as the one provided when the agent
//     first joined the lobby.
//
// Default Behavior:
// By default, this method performs no action. To add custom behavior
// when the lobby state changes, override this method in your implementation.
func (a *Agent) OnLobbyDataChanged(lobbyData *lobby_data.LobbyData) {
	// Implement the logic for handling lobby data changes
}

// NextMove is called after each game tick, when new game state data is received from the server.
// This method is responsible for determining the agent's next move based on the current game state.
//
// Parameters:
//   - gameState: The current state of the game, which includes all necessary information
//     for the agent to decide its next action, such as the entire map with walls, tanks, bullets, zones, etc.
//
// Returns:
// - AgentResponse: The action or decision made by the agent, which will be communicated back to the game server.
func (a *Agent) NextMove(gameState *game_state.GameState) *agent_response.AgentResponse {

	// Find my tank
	var myTank *game_state.Tank
	for _, tank := range gameState.Tanks {
		if tank.OwnerID == a.MyID {
			myTank = &tank
			break
		}
	}

	// If my tank is not found, it is dead
	if myTank == nil {
		return agent_response.NewPass()
	}

	switch r := rand.Float32(); {
	case r < 0.25:
		// Move the tank
		direction := movement.Forward
		if rand.Intn(2) == 1 {
			direction = movement.Backward
		}
		return agent_response.NewMovement(direction)
	case r < 0.50:
		// Rotate the tank and/or turret
		randomRotation := func() string {
			switch rand.Intn(3) {
			case 0:
				return rotation.Left
			case 1:
				return rotation.Right
			default:
				return ""
			}
		}
		return agent_response.NewRotation(randomRotation(), randomRotation())
	case r < 0.75:
		// Use ability
		abilities := []string{
			ability.FireBullet,
			ability.FireDoubleBullet,
			ability.UseLaser,
			ability.UseRadar,
			ability.DropMine,
		}
		abilityType := abilities[rand.Intn(len(abilities))]
		return agent_response.NewAbilityUse(abilityType)
	default:
		// Pass
		return agent_response.NewPass()
	}
}

// OnWarningReceived is called when a warning is received from the server.
// Please remember that if your agent is stuck processing a warning,
// the next move won't be called and vice versa.
//
// Parameters:
// - warning: The warning received from the server.
func (a *Agent) OnWarningReceived(warn warning.Warning, message *string) {
	switch warn {
	case warning.CustomWarning:
		msg := "No message"
		if message != nil {
			msg = *message
		}
		fmt.Printf("[System] ⚠️ Custom Warning: %s\n", msg)
	case warning.PlayerAlreadyMadeActionWarning:
		fmt.Println("[System] ⚠️ Player already made action warning")
	case warning.ActionIgnoredDueToDeadWarning:
		fmt.Println("[System] ⚠️ Action ignored due to dead warning")
	case warning.SlowResponseWarning:
		fmt.Println("[System] ⚠️ Slow response warning")
	}
}

// OnGameEnded is called when the game has concluded, providing the final game results.
// This method is triggered when the game ends, which is when a defined number of ticks in LobbyData has passed.
//
// Parameters:
// - gameEnd: The final state of the game, containing players' scores.
//
// Default Behavior:
// By default, this method performs no action. You can override it to implement any post-game behavior,
// such as logging, updating agent strategies, or other clean-up tasks.
//
// Notes:
// - This method is optional to override, but it can be useful for handling game result analysis and logging.
func (a *Agent) OnGameEnded(gameEnd *game_end.GameEnd) {
	var winner game_end.GameEndPlayer
	for _, player := range gameEnd.Players {
		if player.Score > winner.Score {
			winner = player
		}
	}

	if winner.ID == a.MyID {
		fmt.Println("I won!")
	}

	for _, player := range gameEnd.Players {
		fmt.Printf("Player: %s - Score: %d\n", player.Nickname, player.Score)
	}
}
```

The `Agent` struct in `agent/agent.go` implements the agent's behavior in the game. The `OnJoiningLobby` function is called when the agent is created, the `NextMove` function is called every game tick to determine the agent's next move, and the `OnGameEnded` function is called when the game ends to provide the final game state.

`NextMove` returns an `AgentResponse` struct from `packet/packets/agent_response/agent_response.go`, which can be one of the following:

- `Movement`: Move the tank forward or backward. The `Direction` field is set to "forward" for forward movement and "backward" for backward movement.
- `Rotation`: Rotate the tank body and/or turret. Both `TankRotation` and `TurretRotation` fields use the following values:
  - "": no rotation
  - "left": rotate left
  - "right": rotate right
- `AbilityUse`: Use an ability. The `AbilityType` field specifies which ability to use (e.g., "fireBullet", "fireDoubleBullet", "useLaser", "useRadar", "dropMine").
- `Pass`: Do nothing this turn.

The `GameState` struct in `packet/packets/game_state/game_state.go` represents the current state of the game, including information about tanks, walls, bullets, players, and zones.

You can modify the `agent/agent.go` file and create more files in the `agent` directory. Do not modify any other files, as this may prevent us from running your agent during the competition.

If you want to extend the functionality of the `GameState` struct or other structs, create your own methods or helper functions within the `agent` package.

### Including Static Files

If you need to include static files that your program should access during testing or execution, place them in the `data` folder. This folder is copied into the Docker image and will be accessible to your application at runtime. For example, you could include configuration files, pre-trained models, or any other data your agent might need.


## Running the Client

You can run this client locally, within a VS Code development container, or manually using Docker.

### Running Locally

To run the client locally, you must have Go 1.21 or later installed. Verify
your Go version by running:

```sh
go version
```

Assuming the game server is running on `localhost:5000` (refer to the server
repository's README for setup instructions), start the client by running:

```sh
go run main.go --nickname TEAM_NAME
```

The `--nickname` argument is required and must be unique. For additional
configuration options, run:

```sh
go run main.go --help
```

To build and run an optimized release version of the client, use:

```sh
go run main.go --nickname TEAM_NAME
```

### 2. Running in a VS Code Development Container

To run the client within a VS Code development container, ensure you have Docker
and Visual Studio Code (VS Code) installed, along with the Dev Containers
extension.

Steps:

1. Open the project folder in VS Code.
2. If prompted, choose to reopen the project in a development container and wait
   for the setup to complete.
3. If not prompted, manually reopen the project in a container by:
   - Opening the command palette (`F1`)
   - Searching for and selecting `>Dev Containers: Reopen in Container`

Once the container is running, you can execute all necessary commands in VS
Code's integrated terminal, as if you were running the project locally.

When you are running the client in a container and the server is running on
your local machine, use the `--host host.docker.internal` flag to connect the
Docker container to your local host.

```sh
go run main.go --host host.docker.internal --nickname TEAM_NAME
```

### 3. Running in a Docker Container (Manual Setup)

To run the client manually in a Docker container, ensure Docker is installed on
your system.

Steps:

1. Build the Docker image:
   ```sh
   docker build -t client .
   ```
2. Run the Docker container:
   ```sh
   docker run --rm client --nickname TEAM_NAME --host host.docker.internal
   ```

If the server is running on your local machine, use the
`--host host.docker.internal` flag to connect the Docker container to your local
host.

