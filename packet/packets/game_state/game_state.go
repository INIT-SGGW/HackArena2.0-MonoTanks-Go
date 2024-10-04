package game_state

import (
	"encoding/json"
	"fmt"
)

// GameState represents the current state of the game.
type GameState struct {
	// A unique identifier for the game state.
	ID string

	// A slice of Tank objects representing all the tanks in the game.
	Tanks []Tank

	// A slice of Wall objects representing all the walls in the game.
	Walls []Wall

	// A slice of Bullet objects representing all the bullets in the game.
	Bullets []Bullet

	// A slice of Player objects representing all the players in the game.
	Players []Player

	// The current tick of the game.
	Tick uint64

	// A slice of Zone objects representing different zones in the game.
	Zones []Zone

	// A 2D slice representing the visibility map, where each element is a boolean.
	Visibility [][]bool
}

// RawTank represents the raw JSON structure of a tank.
type RawTank struct {
	// The direction the tank is facing. 0 means up, 1 means right, 2 means down and 3 means left. Other values are not allowed.
	Direction int `json:"direction"`

	// The health of the tank. It is nil for other players tanks.
	Health *int `json:"health"`

	// The ID of the player who owns the tank.
	OwnerID string `json:"ownerId"`

	// The turret of the tank.
	Turret Turret `json:"turret"`
}

// Turret represents the turret of a tank.
type Turret struct {
	// The number of bullets the turret has. It is nil for other players tanks.
	BulletCount *int `json:"bulletCount"`

	// The number of ticks until the turret regenerates a bullet. It is nil for other players tanks.
	TicksToRegenBullet *int `json:"ticksToRegenBullet"`

	// The direction the turret is facing. 0 means up, 1 means right, 2 means down and 3 means left. Other values are not allowed.
	Direction int `json:"direction"`
}

// Tank represents a tank in the game.
type Tank struct {
	// The x-coordinate of the tank.
	X int

	// The y-coordinate of the tank.
	Y int

	// The direction the tank is facing.
	Direction int

	// The health of the tank. It is nil for other players tanks.
	Health *int

	// The ID of the player who owns the tank.
	OwnerID string

	// The turret of the tank.
	Turret Turret
}

// Wall represents a wall in the game.
type Wall struct {
	// The x-coordinate of the wall.
	X int

	// The y-coordinate of the wall.
	Y int
}

// RawBullet represents the raw JSON structure of a bullet.
type RawBullet struct {
	// The direction the bullet is traveling.
	Direction int `json:"direction"`

	// The unique identifier for the bullet.
	ID int `json:"id"`

	// The speed of the bullet.
	Speed float64 `json:"speed"`
}

// Bullet represents a bullet in the game.
type Bullet struct {
	// The x-coordinate of the bullet.
	X int

	// The y-coordinate of the bullet.
	Y int

	// The direction the bullet is traveling. 0 means up, 1 means right, 2 means down and 3 means left. Other values are not allowed.
	Direction int

	// The unique identifier for the bullet.
	ID int

	// The speed of the bullet.
	Speed float64
}

// Player represents a player in the game.
type Player struct {
	// A unique identifier for the player.
	ID string `json:"id"`

	// The player's chosen nickname or alias.
	Nickname string `json:"nickname"`

	// Represents the player's color, used in visual representation as a color for nickname and tank.
	Color uint64 `json:"color"`

	// The player's current ping, representing latency, if available.
	Ping *uint64 `json:"ping,omitempty"`

	// The player's score in the game, if available.
	Score *uint64 `json:"score,omitempty"`

	// Number of ticks (time units) remaining until the player's health or resource regenerates, if applicable. This is when player is dead.
	TicksToRegen *uint64 `json:"ticksToRegen,omitempty"`
}

// Zone represents a zone in the game world.
type Zone struct {
	// The unique index of the zone.
	Index uint8 `json:"index"`

	// The x-coordinate of the left side of the zone.
	X uint64 `json:"x"`

	// The y-coordinate of the top side of the zone.
	Y uint64 `json:"y"`

	// The width of the zone.
	Width uint64 `json:"width"`

	// The height of the zone.
	Height uint64 `json:"height"`

	// The current status of the zone.
	Status ZoneStatus `json:"status"`
}

// ZoneStatus represents the status of a zone.
type ZoneStatus struct {
	// The type of the zone status.
	Type string `json:"type"`

	// The status of the zone being captured, if applicable.
	BeingCaptured *BeingCapturedStatus `json:"beingCaptured,omitempty"`

	// The status of the zone being captured, if applicable.
	Captured *CapturedStatus `json:"captured,omitempty"`

	// The status of the zone being contested, if applicable.
	BeingContested *BeingContestedStatus `json:"beingContested,omitempty"`

	// The status of the zone being retaken, if applicable.
	BeingRetaken *BeingRetakenStatus `json:"beingRetaken,omitempty"`
}

// BeingCapturedStatus represents the status of a zone being captured.
type BeingCapturedStatus struct {
	// The remaining ticks until the zone is captured.
	RemainingTicks uint64 `json:"remainingTicks"`

	// The ID of the player capturing the zone.
	PlayerID string `json:"playerId"`
}

// CapturedStatus represents the status of a zone that has been captured.
type CapturedStatus struct {
	// The ID of the player who captured the zone.
	PlayerID string `json:"playerId"`
}

// BeingContestedStatus represents the status of a zone being contested.
type BeingContestedStatus struct {
	// The ID of the player who captured the zone, if any.
	CapturedByID *string `json:"capturedById,omitempty"`
}

// BeingRetakenStatus represents the status of a zone being retaken.
type BeingRetakenStatus struct {
	// The remaining ticks until the zone is retaken.
	RemainingTicks uint64 `json:"remainingTicks"`

	// The ID of the player who previously captured the zone.
	CapturedByID string `json:"capturedById"`

	// The ID of the player retaking the zone.
	RetakenByID string `json:"retakenById"`
}

// rawGameState is a custom struct to unmarshal the JSON data for the game state.
type rawGameState struct {
	// A unique identifier for the game state.
	ID string `json:"id"`

	// The current tick of the game.
	Tick uint64 `json:"tick"`

	// A slice of Player objects representing all the players in the game.
	Players []Player `json:"players"`

	// The raw JSON message for the map data.
	Map json.RawMessage `json:"map"`
}

// rawMap is a custom struct to unmarshal the map data.
type rawMap struct {
	// A 3D slice representing the tiles in the map.
	Tiles [][][]json.RawMessage `json:"tiles"`

	// A slice of Zone objects representing different zones in the game.
	Zones []Zone `json:"zones"`

	// A slice of strings representing the visibility map.
	Visibility []string `json:"visibility"`
}

// UnmarshalJSON custom unmarshals the JSON data into a GameState object.
func (gameState *GameState) UnmarshalJSON(data []byte) error {
	var raw rawGameState
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	gameState.ID = raw.ID
	gameState.Tick = raw.Tick
	gameState.Players = raw.Players

	var rawMapData rawMap
	if err := json.Unmarshal(raw.Map, &rawMapData); err != nil {
		return err
	}

	gameState.Zones = rawMapData.Zones

	// Process tiles
	for x, column := range rawMapData.Tiles {
		for y, cell := range column {
			if len(cell) > 0 {
				var tileType struct {
					// The type of the tile.
					Type string `json:"type"`
				}
				if err := json.Unmarshal(cell[0], &tileType); err != nil {
					return err
				}

				switch tileType.Type {
				case "wall":
					gameState.Walls = append(gameState.Walls, Wall{X: x, Y: y})
				case "tank":
					var rawTank struct {
						// The payload containing the raw tank data.
						Payload RawTank `json:"payload"`
					}
					if err := json.Unmarshal(cell[0], &rawTank); err != nil {
						return err
					}
					tank := Tank{
						X:         x,
						Y:         y,
						Direction: rawTank.Payload.Direction,
						Health:    rawTank.Payload.Health,
						OwnerID:   rawTank.Payload.OwnerID,
						Turret:    rawTank.Payload.Turret,
					}
					gameState.Tanks = append(gameState.Tanks, tank)
				case "bullet":
					var rawBullet struct {
						// The payload containing the raw bullet data.
						Payload RawBullet `json:"payload"`
					}
					if err := json.Unmarshal(cell[0], &rawBullet); err != nil {
						return err
					}
					bullet := Bullet{
						X:         x,
						Y:         y,
						Direction: rawBullet.Payload.Direction,
						ID:        rawBullet.Payload.ID,
						Speed:     rawBullet.Payload.Speed,
					}
					gameState.Bullets = append(gameState.Bullets, bullet)
				default:
					return fmt.Errorf("unknown tile type: %s", tileType.Type)
				}
			}
		}
	}

	// Process visibility
	gameState.Visibility = make([][]bool, len(rawMapData.Visibility))
	for y, row := range rawMapData.Visibility {
		gameState.Visibility[y] = make([]bool, len(row))
		for x, cell := range row {
			gameState.Visibility[y][x] = cell == '1'
		}
	}

	return nil
}
