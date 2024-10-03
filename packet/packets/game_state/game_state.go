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

type RawTank struct {
	Direction int    `json:"direction"`
	Health    *int   `json:"health"`
	OwnerID   string `json:"ownerId"`
	Turret    Turret `json:"turret"`
}

type Turret struct {
	BulletCount        *int `json:"bulletCount"`
	TicksToRegenBullet *int `json:"ticksToRegenBullet"`
	Direction          int  `json:"direction"`
}

type Tank struct {
	x         int
	y         int
	Direction int
	Health    int
	OwnerID   string
	Turret    Turret
}

type Wall struct {
	x int
	y int
}

type RawBullet struct {
	Direction int     `json:"direction"`
	ID        int     `json:"id"`
	Speed     float64 `json:"speed"`
}

type Bullet struct {
	x         int
	y         int
	Direction int
	ID        int
	Speed     float64
}

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
	Type           string                `json:"type"`
	BeingCaptured  *BeingCapturedStatus  `json:"beingCaptured,omitempty"`
	Captured       *CapturedStatus       `json:"captured,omitempty"`
	BeingContested *BeingContestedStatus `json:"beingContested,omitempty"`
	BeingRetaken   *BeingRetakenStatus   `json:"beingRetaken,omitempty"`
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

// Examples of zones
// {
// 	"x": 4,
// 	"y": 13,
// 	"width": 4,
// 	"height": 4,
// 	"index": 65,
// 	"status": {
// 	  "type": "neutral"
// 	}
//   },
//   {
// 	"x": 3,
// 	"y": 3,
// 	"width": 4,
// 	"height": 4,
// 	"index": 66,
// 	"status": {
// 	  "type": "neutral"
// 	}
//   },
// {
// 	"x": 1,
// 	"y": 1,
// 	"width": 4,
// 	"height": 4,
// 	"index": 65,
// 	"status": {
// 	  "remainingTicks": 100,
// 	  "playerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
// 	  "type": "beingCaptured"
// 	}
//   }

// Example of game state
// {
//     "id": "36af2d82-40df-4332-9138-1bb0e6b09fcb",
//     "tick": 110,
//     "players": [
//       {
//         "id": "80fd0035-364e-4ec6-adc1-96208a580bd4",
//         "nickname": "RUST3",
//         "color": 4294944256,
//         "ping": 1
//       },
//       {
//         "id": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
//         "nickname": "R1",
//         "color": 4294925049,
//         "ping": 0,
//         "score": 10,
//         "ticksToRegen": null
//       }
//     ],
//     "map": {
//       "tiles": [
//         [
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           [],
//           [],
//           []
//         ],
//         [
//           [],
//           [],
//           [],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           []
//         ],
//         [
//           [],
//           [
//             {
//               "type": "tank",
//               "payload": {
//                 "ownerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
//                 "direction": 0,
//                 "turret": {
//                   "direction": 1,
//                   "bulletCount": 0,
//                   "ticksToRegenBullet": 1
//                 },
//                 "health": 100
//               }
//             }
//           ],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           []
//         ],
//         [
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           [],
//           [],
//           [],
//           [],
//           []
//         ],
//         [
//           [],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           [],
//           []
//         ],
//         [
//           [],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           [],
//           []
//         ],
//         [
//           [],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ],
//           [],
//           [],
//           [
//             {
//               "type": "wall"
//             }
//           ]
//         ]
//       ],
//       "zones": [
//         {
//           "x": 1,
//           "y": 1,
//           "width": 4,
//           "height": 4,
//           "index": 65,
//           "status": {
//             "remainingTicks": 100,
//             "playerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
//             "type": "beingCaptured"
//           }
//         }
//       ],
//       "visibility": [
//         "1110000",
//         "0111111",
//         "0000000",
//         "0000000",
//         "0000000",
//         "0000000",
//         "0000000"
//       ]
//     }
//   }

// Implement UnmarshalJSON for GameState

// Custom struct to unmarshal the JSON data
type rawGameState struct {
	ID      string          `json:"id"`
	Tick    uint64          `json:"tick"`
	Players []Player        `json:"players"`
	Map     json.RawMessage `json:"map"`
}

// Custom struct to unmarshal the map data
type rawMap struct {
	Tiles      [][][]json.RawMessage `json:"tiles"`
	Zones      []Zone                `json:"zones"`
	Visibility []string              `json:"visibility"`
}

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
					Type string `json:"type"`
				}
				if err := json.Unmarshal(cell[0], &tileType); err != nil {
					return err
				}

				switch tileType.Type {
				case "wall":
					gameState.Walls = append(gameState.Walls, Wall{x: x, y: y})
				case "tank":
					var rawTank struct {
						Payload RawTank `json:"payload"`
					}
					if err := json.Unmarshal(cell[0], &rawTank); err != nil {
						return err
					}
					tank := Tank{
						x:         x,
						y:         y,
						Direction: rawTank.Payload.Direction,
						Health:    *rawTank.Payload.Health,
						OwnerID:   rawTank.Payload.OwnerID,
						Turret:    rawTank.Payload.Turret,
					}
					gameState.Tanks = append(gameState.Tanks, tank)
				case "bullet":
					var rawBullet struct {
						Payload RawBullet `json:"payload"`
					}
					if err := json.Unmarshal(cell[0], &rawBullet); err != nil {
						return err
					}
					bullet := Bullet{
						x:         x,
						y:         y,
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
