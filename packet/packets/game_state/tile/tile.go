package tile

// import (
// 	"encoding/json"
// )

// // Tile represents a tile on the map.
// type Tile struct {
// 	Visible   bool        `json:"visible"`   // Whether the tile is currently visible by you on the map.
// 	ZoneIndex *uint8      `json:"zoneIndex"` // If tile is in a zone, this is the index of the zone it belongs to.
// 	Payload   TilePayload `json:"payload"`   // The specific payload of the tile, determining its content (e.g., empty, wall, tank, bullet).
// }

// // int64Ptr is a helper function to convert an int64 to a pointer.
// func int64Ptr(i int64) *int64 {
// 	return &i
// }

// // TilePayload is the interface for tile contents.
// type TilePayload interface {
// 	IsTilePayload()
// }

// // EmptyTile represents an empty tile with no contents.
// type EmptyTile struct{}

// func (e EmptyTile) IsTilePayload() {}

// // WallTile represents a tile containing a wall.
// type WallTile struct{}

// func (w WallTile) IsTilePayload() {}

// // TankTile represents a tile containing a tank.
// type TankTile struct {
// 	Tank Tank `json:"payload"`
// }

// func (t TankTile) IsTilePayload() {}

// // BulletTile represents a tile containing a bullet.
// type BulletTile struct {
// 	Bullet Bullet `json:"payload"`
// }

// func (b BulletTile) IsTilePayload() {}

// // UnmarshalJSON implements custom deserialization for Tile to handle the enum-like TilePayload field.
// func (t *Tile) UnmarshalJSON(data []byte) error {
// 	// Create an alias to avoid recursion during UnmarshalJSON
// 	type Alias Tile
// 	aux := &struct {
// 		Payload struct {
// 			Type    string          `json:"type"`
// 			Payload json.RawMessage `json:"payload"`
// 		} `json:"payload"`
// 		*Alias
// 	}{
// 		Alias: (*Alias)(t),
// 	}

// 	// Unmarshal into the aux struct
// 	if err := json.Unmarshal(data, aux); err != nil {
// 		return err
// 	}

// 	// Deserialize the Payload based on the "type" field
// 	switch aux.Payload.Type {
// 	case "empty":
// 		t.Payload = EmptyTile{}
// 	case "wall":
// 		t.Payload = WallTile{}
// 	case "tank":
// 		var tank Tank
// 		if err := json.Unmarshal(aux.Payload.Payload, &tank); err != nil {
// 			return err
// 		}
// 		t.Payload = TankTile{Tank: tank}
// 	case "bullet":
// 		var bullet Bullet
// 		if err := json.Unmarshal(aux.Payload.Payload, &bullet); err != nil {
// 			return err
// 		}
// 		t.Payload = BulletTile{Bullet: bullet}
// 	default:
// 		return json.Unmarshal(data, t) // fallback in case of unexpected type
// 	}

// 	return nil
// }

// // MarshalJSON implements custom serialization for Tile to handle the enum-like TilePayload field.
// func (t Tile) MarshalJSON() ([]byte, error) {
// 	type Alias Tile
// 	aux := &struct {
// 		Payload struct {
// 			Type    string      `json:"type"`
// 			Payload interface{} `json:"payload,omitempty"`
// 		} `json:"payload"`
// 		*Alias
// 	}{
// 		Alias: (*Alias)(&t),
// 	}

// 	// Serialize the Payload based on its type
// 	switch v := t.Payload.(type) {
// 	case EmptyTile:
// 		aux.Payload.Type = "empty"
// 	case WallTile:
// 		aux.Payload.Type = "wall"
// 	case TankTile:
// 		aux.Payload.Type = "tank"
// 		aux.Payload.Payload = v.Tank
// 	case BulletTile:
// 		aux.Payload.Type = "bullet"
// 		aux.Payload.Payload = v.Bullet
// 	}

// 	return json.Marshal(aux)
// }

// // NewTile creates a new Tile with the given visibility, zone index, and payload.
// func NewTile(visible bool, zoneIndex *uint8, payload TilePayload) Tile {
// 	return Tile{
// 		Visible:   visible,
// 		ZoneIndex: zoneIndex,
// 		Payload:   payload,
// 	}
// }

// package tile

// import (
// 	"encoding/json"
// 	"errors"
// )

// // Tile represents a tile on the map.
// type Tile struct {
// 	Visible   bool        `json:"visible"`   // Whether the tile is visible on the map.
// 	ZoneIndex *uint8      `json:"zoneIndex"` // Optional index of the zone the tile belongs to.
// 	Payload   TilePayload `json:"payload"`   // Content (payload) of the tile.
// }

// // TilePayload represents the possible contents of a tile.
// type TilePayload interface{}

// // Enum-like constants for TilePayload types.
// const (
// 	EmptyType  = "empty"
// 	WallType   = "wall"
// 	TankType   = "tank"
// 	BulletType = "bullet"
// )

// // Wall represents a wall tile.

// // Tank represents a tank tile.

// // Bullet represents a bullet tile.
// type Bullet struct {
// 	Direction int     `json:"direction"`
// 	ID        int     `json:"id"`
// 	Speed     float64 `json:"speed"`
// }

// // PayloadWrapper wraps the TilePayload interface to allow custom deserialization.
// type PayloadWrapper struct {
// 	Type    string          `json:"type"`
// 	Payload json.RawMessage `json:"payload"`
// }

// // UnmarshalJSON custom deserializer for TilePayload.
// func (t *Tile) UnmarshalJSON(data []byte) error {
// 	var wrapper PayloadWrapper
// 	if err := json.Unmarshal(data, &wrapper); err != nil {
// 		return err
// 	}

// 	switch wrapper.Type {
// 	case EmptyType:
// 		t.Payload = EmptyType
// 	case WallType:
// 		t.Payload = Wall{}
// 	case TankType:
// 		var tank Tank
// 		if err := json.Unmarshal(wrapper.Payload, &tank); err != nil {
// 			return err
// 		}
// 		t.Payload = tank
// 	case BulletType:
// 		var bullet Bullet
// 		if err := json.Unmarshal(wrapper.Payload, &bullet); err != nil {
// 			return err
// 		}
// 		t.Payload = bullet
// 	default:
// 		return errors.New("invalid payload type")
// 	}

// 	return nil
// }
