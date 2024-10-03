package tile

// import (
// 	"encoding/json"
// 	"hack-arena-2024-h2-go/packet/packets/game_state/tile"
// 	"testing"
// )

// func int64Ptr(i int64) *int64 {
// 	return &i
// }

// func tanksEqual(t1, t2 Tank) bool {
// 	if t1.OwnerID != t2.OwnerID || t1.Direction != t2.Direction || t1.Health == nil || t2.Health == nil || *t1.Health != *t2.Health {
// 		return false
// 	}
// 	if t1.Turret.Direction != t2.Turret.Direction || t1.Turret.BulletCount == nil || t2.Turret.BulletCount == nil || *t1.Turret.BulletCount != *t2.Turret.BulletCount {
// 		return false
// 	}
// 	return true
// }

// func TestTileUnmarshalJSON(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected tile.Tile
// 	}{
// 		{
// 			name:  "WallTile",
// 			input: `{"visible":true,"zoneIndex":null,"payload":{"type":"wall"}}`,
// 			expected: tile.Tile{
// 				Visible:   true,
// 				ZoneIndex: nil,
// 				Payload:   tile.WallTile{},
// 			},
// 		},
// 		{
// 			name:  "TankTile",
// 			input: `{"visible":true,"zoneIndex":null,"payload":{"type":"tank","payload":{"ownerId":"fa47ab49-3faa-49b5-b737-5acd5d36624f","direction":1,"turret":{"direction":1,"bulletCount":3,"ticksToRegenBullet":null},"health":100}}}`,
// 			expected: tile.Tile{
// 				Visible:   true,
// 				ZoneIndex: nil,
// 				Payload: tile.TankTile{
// 					Tank: tile.Tank{
// 						OwnerID:   "fa47ab49-3faa-49b5-b737-5acd5d36624f",
// 						Direction: 1,
// 						Turret: tile.Turret{
// 							Direction:          1,
// 							BulletCount:        int64Ptr(3),
// 							TicksToRegenBullet: nil,
// 						},
// 						Health: int64Ptr(100),
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:  "EmptyTile",
// 			input: `{"visible":true,"zoneIndex":null,"payload":{"type":"empty"}}`,
// 			expected: tile.Tile{
// 				Visible:   true,
// 				ZoneIndex: nil,
// 				Payload:   tile.EmptyTile{},
// 			},
// 		},
// 		{
// 			name:  "BulletTile",
// 			input: `{"visible":true,"zoneIndex":null,"payload":{"type":"bullet","payload":{"direction":1, "id": 1, "speed": 1.0}}}`,
// 			expected: tile.Tile{
// 				Visible:   true,
// 				ZoneIndex: nil,
// 				Payload: tile.BulletTile{
// 					Bullet: tile.Bullet{
// 						Direction: 1,
// 						ID:        1,
// 						Speed:     1.0,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var tTile tile.Tile
// 			if err := json.Unmarshal([]byte(tt.input), &tTile); err != nil {
// 				t.Fatalf("Unmarshal failed: %v", err)
// 			}
// 			if tTile.Visible != tt.expected.Visible || (tTile.ZoneIndex == nil && tt.expected.ZoneIndex != nil) || (tTile.ZoneIndex != nil && tt.expected.ZoneIndex == nil) || (tTile.ZoneIndex != nil && tt.expected.ZoneIndex != nil && *tTile.ZoneIndex != *tt.expected.ZoneIndex) {
// 				t.Errorf("Expected %v, got %v", tt.expected, tTile)
// 			}
// 			switch expectedPayload := tt.expected.Payload.(type) {
// 			case tile.WallTile:
// 				if _, ok := tTile.Payload.(tile.WallTile); !ok {
// 					t.Errorf("Expected tile.WallTile, got %T", tTile.Payload)
// 				}
// 			case tile.TankTile:
// 				actualPayload, ok := tTile.Payload.(tile.TankTile)
// 				if !ok {
// 					t.Errorf("Expected tile.TankTile, got %T", tTile.Payload)
// 				}
// 				if !tanksEqual(actualPayload.Tank, expectedPayload.Tank) {
// 					t.Errorf("Expected %v, got %v", expectedPayload.Tank, actualPayload.Tank)
// 				}
// 			case tile.EmptyTile:
// 				if _, ok := tTile.Payload.(tile.EmptyTile); !ok {
// 					t.Errorf("Expected tile.EmptyTile, got %T", tTile.Payload)
// 				}
// 			case tile.BulletTile:
// 				actualPayload, ok := tTile.Payload.(tile.BulletTile)
// 				if !ok {
// 					t.Errorf("Expected tile.BulletTile, got %T", tTile.Payload)
// 				}
// 				if actualPayload.Bullet != expectedPayload.Bullet {
// 					t.Errorf("Expected %v, got %v", expectedPayload.Bullet, actualPayload.Bullet)
// 				}
// 			}
// 		})
// 	}
// }

// Compare two Tank structs for equality, checking pointer values as well.
// func tanksEqual(t1, t2 Tank) bool {
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
// func turretsEqual(t1, t2 Turret) bool {
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

// // Tests
// func TestDeserializeWall(t *testing.T) {
// 	jsonData := `{"type": "wall"}`
// 	var deserialized Tile
// 	if err := json.Unmarshal([]byte(jsonData), &deserialized); err != nil {
// 		t.Fatalf("Failed to deserialize wall: %v", err)
// 	}

// 	if _, ok := deserialized.Payload.(Wall); !ok {
// 		t.Fatalf("Expected Wall type, got: %v", deserialized.Payload)
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

// 	var deserialized Tile
// 	if err := json.Unmarshal([]byte(jsonData), &deserialized); err != nil {
// 		t.Fatalf("Failed to deserialize tank: %v", err)
// 	}

// 	expectedTank := Tank{
// 		Direction: 1,
// 		Health:    ptr(100),
// 		OwnerID:   "player1",
// 		Turret: Turret{
// 			BulletCount:        ptr(10),
// 			TicksToRegenBullet: ptr(50),
// 			Direction:          0,
// 		},
// 	}

// 	tankPayload, ok := deserialized.Payload.(Tank)
// 	if !ok {
// 		t.Fatalf("Expected Tank type, got: %T", deserialized.Payload)
// 	}
// 	if !tanksEqual(tankPayload, expectedTank) {
// 		t.Fatalf("Expected tank: %v, got: %v", expectedTank, deserialized.Payload)
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

// 	var deserialized Tile
// 	if err := json.Unmarshal([]byte(jsonData), &deserialized); err != nil {
// 		t.Fatalf("Failed to deserialize bullet: %v", err)
// 	}

// 	expectedBullet := Bullet{
// 		Direction: 2,
// 		ID:        1,
// 		Speed:     5.0,
// 	}

// 	if deserialized.Payload != expectedBullet {
// 		t.Fatalf("Expected bullet: %v, got: %v", expectedBullet, deserialized.Payload)
// 	}
// }

// func TestDeserializeInvalidType(t *testing.T) {
// 	jsonData := `{"type": "invalid"}`
// 	var deserialized Tile
// 	err := json.Unmarshal([]byte(jsonData), &deserialized)
// 	if err == nil {
// 		t.Fatalf("Expected error for invalid type, got: %v", deserialized.Payload)
// 	}
// }

// func ptr[T any](v T) *T {
// 	return &v
// }
