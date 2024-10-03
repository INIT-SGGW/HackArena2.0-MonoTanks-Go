package raw_game_state

import (
	"encoding/json"
)

// // Represents a tile on the map.
// type Tile struct {
// 	// Whether the tile is currently visible by you on the map.
// 	Visible bool `json:"visible"`
// 	// If the tile is in a zone, this is the index of the zone it belongs to.
// 	ZoneIndex *uint8 `json:"zoneIndex"`
// 	// The specific payload of the tile, determining its content (e.g., empty, wall, tank, bullet).
// 	Payload TilePayload `json:"payload"`
// }

// Enum representing the possible contents (payloads) of a tile.
type RawTilePayload struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// type Tank struct {
// 	Direction int    `json:"direction"`
// 	Health    *int   `json:"health"`
// 	OwnerID   string `json:"ownerId"`
// 	Turret    Turret `json:"turret"`
// }

// type Bullet struct {
// 	Direction int     `json:"direction"`
// 	ID        int     `json:"id"`
// 	Speed     float64 `json:"speed"`
// }

// type Turret struct {
// 	BulletCount        *int `json:"bulletCount"`
// 	TicksToRegenBullet *int `json:"ticksToRegenBullet"`
// 	Direction          int  `json:"direction"`
// }

// UnmarshalTilePayload unmarshals the JSON data into the appropriate TilePayload type.
func UnmarshalTilePayload(data []byte) (RawTilePayload, error) {
	var payload RawTilePayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return RawTilePayload{}, err
	}
	return payload, nil
}
