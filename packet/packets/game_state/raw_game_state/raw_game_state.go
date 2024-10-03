package raw_game_state

// import (
// 	"hack-arena-2024-h2-go/packet/packets/game_state"
// 	"hack-arena-2024-h2-go/packet/packets/game_state/tile"
// )

// // RawGameState represents the raw state of the game.
// type RawGameState struct {
// 	ID      string              `json:"id"`
// 	Tick    uint64              `json:"tick"`
// 	Players []game_state.Player `json:"players"`
// 	Map     RawMap              `json:"map"`
// }

// // ToGameState converts a RawGameState to a GameState.
// func (rawGameState *RawGameState) ToGameState() *game_state.GameState {
// 	x := len(rawGameState.Map.Tiles)
// 	y := len(rawGameState.Map.Tiles[0])

// 	mapData := make([][]tile.Tile, y)
// 	for i := range mapData {
// 		mapData[i] = make([]tile.Tile, x)
// 		for j := range mapData[i] {
// 			mapData[i][j] = tile.Tile{
// 				Visible:   false,
// 				ZoneIndex: nil,
// 				Payload:   tile.EmptyType,
// 			}
// 		}
// 	}

// 	// Payload
// 	for x, column := range rawGameState.Map.Tiles {
// 		for y, row := range column {
// 			if len(row) > 0 {
// 				payload := row[0]
// 				mapData[y][x].Payload = payload

// 			}
// 		}
// 	}

// 	// Visibility
// 	for y, row := range rawGameState.Map.Visibility {
// 		for x, column := range row {
// 			mapData[y][x].Visible = column == '1'
// 		}
// 	}

// 	// Zone index
// 	for _, zone := range rawGameState.Map.Zones {
// 		for y := zone.Y; y < zone.Y+zone.Height; y++ {
// 			for x := zone.X; x < zone.X+zone.Width; x++ {
// 				mapData[y][x].ZoneIndex = &zone.Index
// 			}
// 		}
// 	}

// 	return &game_state.GameState{
// 		Map:     mapData,
// 		Players: rawGameState.Players,
// 		Tick:    rawGameState.Tick,
// 		Zones:   rawGameState.Map.Zones,
// 	}
// }
