package raw_game_state

// import (
// 	"encoding/json"
// 	"hack-arena-2024-h2-go/packet/packets/game_state/tile"
// 	"testing"
// )

// // Helper function to return a pointer to an integer.
// func int64Ptr(i int64) *int64 {
// 	return &i
// }

// // Compare two Tank structs for equality, checking pointer values as well.
// func tanksEqual(t1, t2 tile.Tank) bool {
// 	if t1.Direction != t2.Direction {
// 		return false
// 	}
// 	if (t1.Health == nil && t2.Health != nil) || (t1.Health != nil && t2.Health == nil) {
// 		return false
// 	}
// 	if t1.Health != nil && t2.Health != nil && *t1.Health != *t2.Health {
// 		return false
// 	}
// 	if t1.OwnerID != t2.OwnerID {
// 		return false
// 	}
// 	return turretsEqual(t1.Turret, t2.Turret)
// }

// // Compare two Turret structs for equality, checking pointer values as well.
// func turretsEqual(t1, t2 tile.Turret) bool {
// 	if (t1.BulletCount == nil && t2.BulletCount != nil) || (t1.BulletCount != nil && t2.BulletCount == nil) {
// 		return false
// 	}
// 	if t1.BulletCount != nil && t2.BulletCount != nil && *t1.BulletCount != *t2.BulletCount {
// 		return false
// 	}
// 	if (t1.TicksToRegenBullet == nil && t2.TicksToRegenBullet != nil) || (t1.TicksToRegenBullet != nil && t2.TicksToRegenBullet == nil) {
// 		return false
// 	}
// 	if t1.TicksToRegenBullet != nil && t2.TicksToRegenBullet != nil && *t1.TicksToRegenBullet != *t2.TicksToRegenBullet {
// 		return false
// 	}
// 	return t1.Direction == t2.Direction
// }

// func TestDeserializeWall(t *testing.T) {
// 	jsonData := `{"type": "wall"}`
// 	var payload RawTilePayload
// 	err := json.Unmarshal([]byte(jsonData), &payload)
// 	if err != nil {
// 		t.Fatalf("Error during Unmarshal: %v", err)
// 	}
// 	if payload.Type != "wall" {
// 		t.Fatalf("Expected 'wall', got: %v", payload.Type)
// 	}
// }

// func TestDeserializeTank(t *testing.T) {
// 	jsonData := `{
// 		"type": "tank",
// 		"payload": {
// 			"direction": 1,
// 			"health": 100,
// 			"ownerId": "player1",
// 			"turret": {
// 				"bulletCount": 10,
// 				"ticksToRegenBullet": 50,
// 				"direction": 0
// 			}
// 		}
// 	}`

// 	var payload RawTilePayload
// 	err := json.Unmarshal([]byte(jsonData), &payload)
// 	if err != nil {
// 		t.Fatalf("Error during Unmarshal: %v", err)
// 	}

// 	var tank tile.Tank
// 	err = json.Unmarshal(payload.Payload, &tank)
// 	if err != nil {
// 		t.Fatalf("Error during Unmarshal Tank: %v", err)
// 	}

// 	expectedTank := tile.Tank{
// 		Direction: 1,
// 		Health:    int64Ptr(100),
// 		OwnerID:   "player1",
// 		Turret: tile.Turret{
// 			BulletCount:        int64Ptr(10),
// 			TicksToRegenBullet: int64Ptr(50),
// 			Direction:          0,
// 		},
// 	}

// 	// Use the custom comparison function.
// 	if !tanksEqual(tank, expectedTank) {
// 		t.Fatalf("Expected %v, got %v", expectedTank, tank)
// 	}
// }

// func TestDeserializeBullet(t *testing.T) {
// 	jsonData := `{
// 		"type": "bullet",
// 		"payload": {
// 			"direction": 2,
// 			"id": 1,
// 			"speed": 5.0
// 		}
// 	}`

// 	var payload RawTilePayload
// 	err := json.Unmarshal([]byte(jsonData), &payload)
// 	if err != nil {
// 		t.Fatalf("Error during Unmarshal: %v", err)
// 	}

// 	var bullet tile.Bullet
// 	err = json.Unmarshal(payload.Payload, &bullet)
// 	if err != nil {
// 		t.Fatalf("Error during Unmarshal Bullet: %v", err)
// 	}

// 	expectedBullet := tile.Bullet{
// 		Direction: 2,
// 		ID:        1,
// 		Speed:     5.0,
// 	}

// 	if bullet != expectedBullet {
// 		t.Fatalf("Expected %v, got %v", expectedBullet, bullet)
// 	}
// }
