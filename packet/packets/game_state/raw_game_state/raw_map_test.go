package raw_game_state_test

// import (
// 	"encoding/json"
// 	"testing"

// 	"hack-arena-2024-h2-go/packet/packets/game_state/raw_game_state"
// 	"hack-arena-2024-h2-go/packet/packets/game_state/tile"
// )

// func int64Ptr(i int64) *int64 {
// 	return &i
// }

// func intPtr(i int) *int {
// 	return &i
// }

// func TestRawMapUnmarshalJSON(t *testing.T) {
// 	input := `{
//         "tiles": [
//             [
//                 [],
//                 [{"type":"tank","payload":{"ownerId":"fa47ab49-3faa-49b5-b737-5acd5d36624f","direction":1,"turret":{"direction":1,"bulletCount":3,"ticksToRegenBullet":null},"health":100}}]
//             ],
//             [
//                 [{"type":"wall"}],
//                 []
//             ]
//         ],
//         "zones": [],
//         "visibility": [
//             "10",
//             "01"
//         ]
//     }`

// 	var rawMap raw_game_state.RawMap
// 	err := json.Unmarshal([]byte(input), &rawMap)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal JSON: %v", err)
// 	}

// 	// Manually convert the nested maps into the specific types
// 	for i := range rawMap.Tiles {
// 		for j := range rawMap.Tiles[i] {
// 			for k := range rawMap.Tiles[i][j] {
// 				tileMap, ok := rawMap.Tiles[i][j][k].(map[string]interface{})
// 				if !ok {
// 					t.Fatalf("Expected map[string]interface{}, got %T", rawMap.Tiles[i][j][k])
// 				}
// 				switch tileMap["type"] {
// 				case "tank":
// 					var tank tile.Tank
// 					payload, err := json.Marshal(tileMap["payload"])
// 					if err != nil {
// 						t.Fatalf("Failed to marshal payload: %v", err)
// 					}
// 					err = json.Unmarshal(payload, &tank)
// 					if err != nil {
// 						t.Fatalf("Failed to unmarshal tank: %v", err)
// 					}
// 					rawMap.Tiles[i][j][k] = tank
// 				case "wall":
// 					rawMap.Tiles[i][j][k] = tile.Wall{}
// 				default:
// 					t.Fatalf("Unknown tile type: %s", tileMap["type"])
// 				}
// 			}
// 		}
// 	}

// 	var tank = tile.Tank{
// 		OwnerID:   "fa47ab49-3faa-49b5-b737-5acd5d36624f",
// 		Direction: 1,
// 		Turret: tile.Turret{
// 			Direction:          1,
// 			BulletCount:        intPtr(3),
// 			TicksToRegenBullet: nil,
// 		},
// 		Health: intPtr(100),
// 	}

// 	var wall = tile.Wall{}

// 	// Check the tiles
// 	expectedTiles := [][][]tile.TilePayload{
// 		{
// 			{},
// 			{
// 				tank,
// 			},
// 		},
// 		{
// 			{
// 				wall,
// 			},
// 			{},
// 		},
// 	}

// 	// Ensure the rawMap.Tiles slice is not nil and has the expected lengths
// 	if rawMap.Tiles == nil || len(rawMap.Tiles) != len(expectedTiles) {
// 		t.Fatalf("Expected %d rows in tiles, got %d", len(expectedTiles), len(rawMap.Tiles))
// 	}

// 	for i := range expectedTiles {
// 		if len(rawMap.Tiles[i]) != len(expectedTiles[i]) {
// 			t.Fatalf("Expected %d columns in row %d, got %d", len(expectedTiles[i]), i, len(rawMap.Tiles[i]))
// 		}
// 		for j := range expectedTiles[i] {
// 			if len(rawMap.Tiles[i][j]) != len(expectedTiles[i][j]) {
// 				t.Fatalf("Expected %d tiles in row %d, column %d, got %d", len(expectedTiles[i][j]), i, j, len(rawMap.Tiles[i][j]))
// 			}
// 		}
// 	}

// 	// Test each tile individually, not in a loop
// 	tankTile, ok := rawMap.Tiles[0][1][0].(tile.Tank)
// 	if !ok {
// 		t.Errorf("Expected tank, got %T", rawMap.Tiles[0][1][0])
// 	} else if !tanksEqual(tankTile, expectedTiles[0][1][0].(tile.Tank)) {
// 		t.Errorf("Expected tank %v, got %v", expectedTiles[0][1][0], rawMap.Tiles[0][1][0])
// 	}
// 	if len(rawMap.Tiles[0][0]) != 0 {
// 		t.Errorf("Expected nil, got %v", rawMap.Tiles[0][0])
// 	}
// 	if len(rawMap.Tiles[1][1]) != 0 {
// 		t.Errorf("Expected nil, got %v", rawMap.Tiles[1][1])
// 	}
// 	wallTile, ok := rawMap.Tiles[1][0][0].(tile.Wall)
// 	if !ok {
// 		t.Errorf("Expected wall, got %T", rawMap.Tiles[1][0][0])
// 	} else if !wallsEqual(wallTile, expectedTiles[1][0][0].(tile.Wall)) {
// 		t.Errorf("Expected wall %v, got %v", expectedTiles[1][0][0], rawMap.Tiles[1][0][0])
// 	}

// 	// Check the visibility
// 	expectedVisibility := []string{"10", "01"}
// 	for i, visibility := range rawMap.Visibility {
// 		if visibility != expectedVisibility[i] {
// 			t.Errorf("Expected visibility %s, got %s", expectedVisibility[i], visibility)
// 		}
// 	}

// 	// Check the zones
// 	if len(rawMap.Zones) != 0 {
// 		t.Errorf("Expected 0 zones, got %d", len(rawMap.Zones))
// 	}
// }

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

// func turretsEqual(t1, t2 tile.Turret) bool {
// 	if t1.Direction != t2.Direction {
// 		return false
// 	}
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
// 	return true
// }

// func wallsEqual(w1, w2 tile.Wall) bool {
// 	// Since Wall struct has no fields, all Wall instances are considered equal
// 	return true
// }
