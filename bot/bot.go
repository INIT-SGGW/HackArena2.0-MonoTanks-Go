package bot

import (
	"fmt"
	"math/rand"

	"hackarena2-0-mono-tanks-go/packet/packets/bot_response"
	"hackarena2-0-mono-tanks-go/packet/packets/bot_response/ability"
	"hackarena2-0-mono-tanks-go/packet/packets/bot_response/movement"
	"hackarena2-0-mono-tanks-go/packet/packets/bot_response/rotation"
	"hackarena2-0-mono-tanks-go/packet/packets/game_end"
	"hackarena2-0-mono-tanks-go/packet/packets/game_state"
	"hackarena2-0-mono-tanks-go/packet/packets/lobby_data"
	"hackarena2-0-mono-tanks-go/packet/warning"
)

// Bot represents an AI player in the game.
type Bot struct {
	MyID string
}

// OnJoiningLobby is called when the bot joins a lobby, creating a new instance of the bot.
// This method initializes the bot with the lobby's current state and other relevant details.
//
// Parameters:
//   - lobbyData: The initial state of the lobby when the bot joins.
//     Contains information like player data, game settings, etc.
//
// Returns:
// - A new instance of the bot.
func OnJoiningLobby(lobbyData *lobby_data.LobbyData) *Bot {
	return &Bot{
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
//     This is the same data structure as the one provided when the bot
//     first joined the lobby.
//
// Default Behavior:
// By default, this method performs no action. To add custom behavior
// when the lobby state changes, override this method in your implementation.
func (b *Bot) OnLobbyDataChanged(lobbyData *lobby_data.LobbyData) {
	// Implement the logic for handling lobby data changes
}

// NextMove is called after each game tick, when new game state data is received from the server.
// This method is responsible for determining the bot's next move based on the current game state.
//
// Parameters:
//   - gameState: The current state of the game, which includes all necessary information
//     for the bot to decide its next action, such as the entire map with walls, tanks, bullets, zones, etc.
//
// Returns:
// - BotResponse: The action or decision made by the bot, which will be communicated back to the game server.
func (b *Bot) NextMove(gameState *game_state.GameState) *bot_response.BotResponse {

	// Print map as ascii
	row_number := len(gameState.Visibility)
	col_number := len(gameState.Visibility[0])

	fmt.Println("Map:")
	for row := 0; row < row_number; row++ {

	out:
		for col := 0; col < col_number; col++ {

			isVisible := gameState.Visibility[row][col]

			for _, wall := range gameState.Walls {
				if wall.X == col && wall.Y == row {
					fmt.Print("# ")
					continue out
				}
			}

			for _, tank := range gameState.Tanks {
				if tank.X == col && tank.Y == row {
					if tank.OwnerID == b.MyID {
						if tank.Direction == "up" {
							fmt.Print("^ ")
						} else if tank.Direction == "down" {
							fmt.Print("v ")
						} else if tank.Direction == "left" {
							fmt.Print("< ")
						} else {
							fmt.Print("> ")
						}
					} else {
						fmt.Print("T ")
					}
					continue out
				}
			}

			for _, bullet := range gameState.Bullets {
				if bullet.X == col && bullet.Y == row {
					if bullet.Type == "basic" {
						if bullet.Direction == "up" {
							fmt.Print("↑ ")
						} else if bullet.Direction == "down" {
							fmt.Print("↓ ")
						} else if bullet.Direction == "left" {
							fmt.Print("← ")
						} else {
							fmt.Print("→ ")
						}
					} else {
						if bullet.Direction == "up" {
							fmt.Print("⇈ ")
						} else if bullet.Direction == "down" {
							fmt.Print("⇊ ")
						} else if bullet.Direction == "left" {
							fmt.Print("⇇ ")
						} else {
							fmt.Print("⇉ ")
						}
					}
					continue out
				}
			}

			for _, laser := range gameState.Lasers {
				if laser.X == col && laser.Y == row {
					if laser.Orientation == "horizontal" {
						fmt.Print("═ ")
					} else {
						fmt.Print("║ ")
					}
					continue out
				}
			}

			for _, mine := range gameState.Mines {
				if mine.X == col && mine.Y == row {
					fmt.Print("X ")
					continue out
				}
			}

			for _, item := range gameState.Items {
				if item.X == col && item.Y == row {
					if item.Type == "doubleBullet" {
						fmt.Print("D ")
					} else if item.Type == "laser" {
						fmt.Print("L ")
					} else if item.Type == "radar" {
						fmt.Print("R ")
					} else if item.Type == "mine" {
						fmt.Print("M ")
					}
					continue out
				}
			}

			for _, zone := range gameState.Zones {
				start_x := int(zone.X)
				start_y := int(zone.Y)
				end_x := int(zone.X) + int(zone.Width)
				end_y := int(zone.Y) + int(zone.Height)

				if col >= start_x && col <= end_x && row >= start_y && row <= end_y {
					if isVisible {
						fmt.Print(string(zone.Index) + " ")
					} else {
						fmt.Print(string(zone.Index+32) + " ")
					}
					continue out
				}
			}

			if isVisible {
				fmt.Print(". ")
				continue out
			}

			fmt.Print("  ")

		}
		fmt.Println()
	}

	// Find my tank
	var myTank *game_state.Tank
	for _, tank := range gameState.Tanks {
		if tank.OwnerID == b.MyID {
			myTank = &tank
			break
		}
	}

	// If my tank is not found, it is dead
	if myTank == nil {
		return bot_response.NewPass()
	}

	switch r := rand.Float32(); {
	case r < 0.25:
		// Move the tank
		direction := movement.Forward
		if rand.Intn(2) == 1 {
			direction = movement.Backward
		}
		return bot_response.NewMovement(direction)
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
		return bot_response.NewRotation(randomRotation(), randomRotation())
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
		return bot_response.NewAbilityUse(abilityType)
	default:
		// Pass
		return bot_response.NewPass()
	}
}

// OnWarningReceived is called when a warning is received from the server.
// Please remember that if your bot is stuck processing a warning,
// the next move won't be called and vice versa.
//
// Parameters:
// - warning: The warning received from the server.
func (b *Bot) OnWarningReceived(warn warning.Warning, message *string) {
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
// such as logging, or other clean-up tasks.
//
// Notes:
// - This method is optional to override, but it can be useful for handling game result analysis and logging.
func (b *Bot) OnGameEnded(gameEnd *game_end.GameEnd) {
	var winner game_end.GameEndPlayer
	for _, player := range gameEnd.Players {
		if player.Score > winner.Score {
			winner = player
		}
	}

	if winner.ID == b.MyID {
		fmt.Println("I won!")
	}

	for _, player := range gameEnd.Players {
		fmt.Printf("Player: %s - Score: %d\n", player.Nickname, player.Score)
	}
}
