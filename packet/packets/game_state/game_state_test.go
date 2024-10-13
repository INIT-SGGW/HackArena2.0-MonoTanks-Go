package game_state

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected GameState
	}{
		{
			name: "Basic GameState",
			jsonData: `{
                "id": "36af2d82-40df-4332-9138-1bb0e6b09fcb",
                "tick": 110,
                "players": [
                    {
                        "id": "80fd0035-364e-4ec6-adc1-96208a580bd4",
                        "nickname": "GO3",
                        "color": 4294944256,
                        "ping": 1
                    },
                    {
                        "id": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
                        "nickname": "GO1",
                        "color": 4294925049,
                        "ping": 0,
                        "score": 10,
                        "ticksToRegen": null
                    }
                ],
                "map": {
                    "tiles": [
                        [
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            [],
                            [],
                            []
                        ],
                        [
                            [],
                            [],
                            [],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            []
                        ],
                        [
                            [],
                            [
                                {
                                    "type": "tank",
                                    "payload": {
                                        "ownerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
                                        "direction": "up",
                                        "turret": {
                                            "direction": "right",
                                            "bulletCount": 0,
                                            "ticksToRegenBullet": 1
                                        },
                                        "health": 100,
                                        "secondaryItem": "laser"
                                    }
                                }
                            ],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            []
                        ],
                        [
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            [],
                            [
                                {
                                    "type": "bullet",
                                    "payload": {
                                        "direction": "down",
                                        "id": 1,
                                        "speed": 0.5,
                                        "type": "basic"
                                    }
                                }
                            ],
                            [],
                            [],
                            []
                        ],
                        [
                            [],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            [],
                            []
                        ],
                        [
                            [],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            [],
                            []
                        ],
                        [
                            [],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ],
                            [],
                            [],
                            [
                                {
                                    "type": "wall"
                                }
                            ]
                        ],
                        [
                            [],
                            [],
                            [],
                            [],
                            [],
                            [],
                            [],
                            [
                                {
                                    "type": "item",
                                    "payload": {
                                        "type": "doubleBullet"
                                    }
                                }
                            ]
                        ]
                    ],
                    "zones": [
                        {
                            "x": 1,
                            "y": 1,
                            "width": 4,
                            "height": 4,
                            "index": 65,
                            "status": {
                                "remainingTicks": 100,
                                "playerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
                                "type": "beingCaptured"
                            }
                        }
                    ],
                    "visibility": [
                        "1110000",
                        "0111111",
                        "0000000",
                        "0000000",
                        "0000000",
                        "0000000",
                        "0000000"
                    ]
                }
            }`,
			expected: GameState{
				ID:   "36af2d82-40df-4332-9138-1bb0e6b09fcb",
				Tick: 110,
				Players: []Player{
					{
						ID:       "80fd0035-364e-4ec6-adc1-96208a580bd4",
						Nickname: "GO3",
						Color:    4294944256,
						Ping:     uint64Ptr(1),
					},
					{
						ID:       "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
						Nickname: "GO1",
						Color:    4294925049,
						Ping:     uint64Ptr(0),
						Score:    uint64Ptr(10),
					},
				},
				Walls: []Wall{
					{X: 0, Y: 2},
					{X: 1, Y: 5},
					{X: 2, Y: 4},
					{X: 3, Y: 0},
					{X: 4, Y: 3},
					{X: 5, Y: 3},
					{X: 6, Y: 3},
					{X: 6, Y: 6},
				},
				Tanks: []Tank{
					{
						X:         2,
						Y:         1,
						Direction: "up",
						Health:    intPtr(100),
						OwnerID:   "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
						Turret: Turret{
							Direction:          "right",
							BulletCount:        intPtr(0),
							TicksToRegenBullet: intPtr(1),
						},
						SecondaryItem: stringPtr("laser"),
					},
				},
				Bullets: []Bullet{
					{
						X:         3,
						Y:         3,
						Direction: "down",
						ID:        1,
						Speed:     0.5,
						Type:      "basic",
					},
				},
				Zones: []Zone{
					{
						X:      1,
						Y:      1,
						Width:  4,
						Height: 4,
						Index:  65,
						Status: ZoneStatus{
							Type: "beingCaptured",
							BeingCaptured: &BeingCapturedStatus{
								RemainingTicks: 100,
								PlayerID:       "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
							},
						},
					},
				},
				Visibility: [][]bool{
					{true, true, true, false, false, false, false},
					{false, true, true, true, true, true, true},
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
				},
				Items: []Item{
					{
						X:    7,
						Y:    7,
						Type: "doubleBullet",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gameState GameState
			if err := json.Unmarshal([]byte(tt.jsonData), &gameState); err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}

			if gameState.ID != tt.expected.ID {
				t.Errorf("expected ID = %v, got %v", tt.expected.ID, gameState.ID)
			}

			if gameState.Tick != tt.expected.Tick {
				t.Errorf("expected Tick = %v, got %v", tt.expected.Tick, gameState.Tick)
			}

			if len(gameState.Players) != len(tt.expected.Players) {
				t.Fatalf("expected %d players, got %d", len(tt.expected.Players), len(gameState.Players))
			}

			for i, player := range gameState.Players {
				if player.ID != tt.expected.Players[i].ID {
					t.Errorf("expected Player ID = %v, got %v", tt.expected.Players[i].ID, player.ID)
				}
				if player.Nickname != tt.expected.Players[i].Nickname {
					t.Errorf("expected Player Nickname = %v, got %v", tt.expected.Players[i].Nickname, player.Nickname)
				}
				if player.Color != tt.expected.Players[i].Color {
					t.Errorf("expected Player Color = %v, got %v", tt.expected.Players[i].Color, player.Color)
				}
				if player.Ping != nil && tt.expected.Players[i].Ping != nil && *player.Ping != *tt.expected.Players[i].Ping {
					t.Errorf("expected Player Ping = %v, got %v", *tt.expected.Players[i].Ping, *player.Ping)
				}
				if player.Score != nil && tt.expected.Players[i].Score != nil && *player.Score != *tt.expected.Players[i].Score {
					t.Errorf("expected Player Score = %v, got %v", *tt.expected.Players[i].Score, *player.Score)
				}
			}

			if len(gameState.Walls) != len(tt.expected.Walls) {
				t.Fatalf("expected %d walls, got %d", len(tt.expected.Walls), len(gameState.Walls))
			}

			for i, wall := range gameState.Walls {
				if wall.X != tt.expected.Walls[i].X || wall.Y != tt.expected.Walls[i].Y {
					t.Errorf("expected Wall = (%v, %v), got (%v, %v)", tt.expected.Walls[i].X, tt.expected.Walls[i].Y, wall.X, wall.Y)
				}
			}

			if len(gameState.Tanks) != len(tt.expected.Tanks) {
				t.Fatalf("expected %d tanks, got %d", len(tt.expected.Tanks), len(gameState.Tanks))
			}

			for i, tank := range gameState.Tanks {
				if tank.X != tt.expected.Tanks[i].X || tank.Y != tt.expected.Tanks[i].Y {
					t.Errorf("expected Tank = (%v, %v), got (%v, %v)", tt.expected.Tanks[i].X, tt.expected.Tanks[i].Y, tank.X, tank.Y)
				}
				if tank.Direction != tt.expected.Tanks[i].Direction {
					t.Errorf("expected Tank Direction = %v, got %v", tt.expected.Tanks[i].Direction, tank.Direction)
				}
				if (tank.Health == nil && tt.expected.Tanks[i].Health != nil) ||
					(tank.Health != nil && tt.expected.Tanks[i].Health == nil) ||
					(tank.Health != nil && tt.expected.Tanks[i].Health != nil && *tank.Health != *tt.expected.Tanks[i].Health) {
					t.Errorf("expected Tank Health = %v, got %v", tt.expected.Tanks[i].Health, tank.Health)
				}
				if tank.OwnerID != tt.expected.Tanks[i].OwnerID {
					t.Errorf("expected Tank OwnerID = %v, got %v", tt.expected.Tanks[i].OwnerID, tank.OwnerID)
				}
				if tank.Turret.Direction != tt.expected.Tanks[i].Turret.Direction {
					t.Errorf("expected Tank Turret Direction = %v, got %v", tt.expected.Tanks[i].Turret.Direction, tank.Turret.Direction)
				}
				if (tank.Turret.BulletCount == nil && tt.expected.Tanks[i].Turret.BulletCount != nil) ||
					(tank.Turret.BulletCount != nil && tt.expected.Tanks[i].Turret.BulletCount == nil) ||
					(tank.Turret.BulletCount != nil && tt.expected.Tanks[i].Turret.BulletCount != nil && *tank.Turret.BulletCount != *tt.expected.Tanks[i].Turret.BulletCount) {
					t.Errorf("expected Tank Turret BulletCount = %v, got %v", tt.expected.Tanks[i].Turret.BulletCount, tank.Turret.BulletCount)
				}
				if (tank.Turret.TicksToRegenBullet == nil && tt.expected.Tanks[i].Turret.TicksToRegenBullet != nil) ||
					(tank.Turret.TicksToRegenBullet != nil && tt.expected.Tanks[i].Turret.TicksToRegenBullet == nil) ||
					(tank.Turret.TicksToRegenBullet != nil && tt.expected.Tanks[i].Turret.TicksToRegenBullet != nil && *tank.Turret.TicksToRegenBullet != *tt.expected.Tanks[i].Turret.TicksToRegenBullet) {
					t.Errorf("expected Tank Turret TicksToRegenBullet = %v, got %v", tt.expected.Tanks[i].Turret.TicksToRegenBullet, tank.Turret.TicksToRegenBullet)
				}
				if tank.SecondaryItem != nil && tt.expected.Tanks[i].SecondaryItem != nil {
					if *tank.SecondaryItem != *tt.expected.Tanks[i].SecondaryItem {
						t.Errorf("expected Tank SecondaryItem = %v, got %v", *tt.expected.Tanks[i].SecondaryItem, *tank.SecondaryItem)
					}
				} else if (tank.SecondaryItem == nil) != (tt.expected.Tanks[i].SecondaryItem == nil) {
					t.Errorf("expected Tank SecondaryItem = %v, got %v", tt.expected.Tanks[i].SecondaryItem, tank.SecondaryItem)
				}
			}

			if len(gameState.Bullets) != len(tt.expected.Bullets) {
				t.Fatalf("expected %d bullets, got %d", len(tt.expected.Bullets), len(gameState.Bullets))
			}

			for i, bullet := range gameState.Bullets {
				if bullet.X != tt.expected.Bullets[i].X || bullet.Y != tt.expected.Bullets[i].Y {
					t.Errorf("expected Bullet = (%v, %v), got (%v, %v)", tt.expected.Bullets[i].X, tt.expected.Bullets[i].Y, bullet.X, bullet.Y)
				}
				if bullet.Direction != tt.expected.Bullets[i].Direction {
					t.Errorf("expected Bullet Direction = %v, got %v", tt.expected.Bullets[i].Direction, bullet.Direction)
				}
				if bullet.ID != tt.expected.Bullets[i].ID {
					t.Errorf("expected Bullet ID = %v, got %v", tt.expected.Bullets[i].ID, bullet.ID)
				}
				if bullet.Speed != tt.expected.Bullets[i].Speed {
					t.Errorf("expected Bullet Speed = %v, got %v", tt.expected.Bullets[i].Speed, bullet.Speed)
				}
				if bullet.Type != tt.expected.Bullets[i].Type {
					t.Errorf("expected Bullet Type = %v, got %v", tt.expected.Bullets[i].Type, bullet.Type)
				}
			}

			if len(gameState.Zones) != len(tt.expected.Zones) {
				t.Fatalf("expected %d zones, got %d", len(tt.expected.Zones), len(gameState.Zones))
			}

			for i, zone := range gameState.Zones {
				if zone.X != tt.expected.Zones[i].X || zone.Y != tt.expected.Zones[i].Y {
					t.Errorf("expected Zone = (%v, %v), got (%v, %v)", tt.expected.Zones[i].X, tt.expected.Zones[i].Y, zone.X, zone.Y)
				}
				if zone.Width != tt.expected.Zones[i].Width {
					t.Errorf("expected Zone Width = %v, got %v", tt.expected.Zones[i].Width, zone.Width)
				}
				if zone.Height != tt.expected.Zones[i].Height {
					t.Errorf("expected Zone Height = %v, got %v", tt.expected.Zones[i].Height, zone.Height)
				}
				if zone.Index != tt.expected.Zones[i].Index {
					t.Errorf("expected Zone Index = %v, got %v", tt.expected.Zones[i].Index, zone.Index)
				}
				if zone.Status.Type != tt.expected.Zones[i].Status.Type {
					t.Errorf("expected Zone Status Type = %v, got %v", tt.expected.Zones[i].Status.Type, zone.Status.Type)
				}
				if zone.Status.BeingCaptured != nil && tt.expected.Zones[i].Status.BeingCaptured != nil {
					if zone.Status.BeingCaptured.RemainingTicks != tt.expected.Zones[i].Status.BeingCaptured.RemainingTicks {
						t.Errorf("expected Zone Status BeingCaptured RemainingTicks = %v, got %v", tt.expected.Zones[i].Status.BeingCaptured.RemainingTicks, zone.Status.BeingCaptured.RemainingTicks)
					}
					if zone.Status.BeingCaptured.PlayerID != tt.expected.Zones[i].Status.BeingCaptured.PlayerID {
						t.Errorf("expected Zone Status BeingCaptured PlayerID = %v, got %v", tt.expected.Zones[i].Status.BeingCaptured.PlayerID, zone.Status.BeingCaptured.PlayerID)
					}
				}
			}

			if len(gameState.Visibility) != len(tt.expected.Visibility) {
				t.Fatalf("expected %d visibility rows, got %d", len(tt.expected.Visibility), len(gameState.Visibility))
			}

			for y, row := range gameState.Visibility {
				if len(row) != len(tt.expected.Visibility[y]) {
					t.Fatalf("expected %d visibility columns in row %d, got %d", len(tt.expected.Visibility[y]), y, len(row))
				}
				for x, cell := range row {
					if cell != tt.expected.Visibility[y][x] {
						t.Errorf("expected Visibility[%d][%d] = %v, got %v", y, x, tt.expected.Visibility[y][x], cell)
					}
				}
			}

			if len(gameState.Items) != len(tt.expected.Items) {
				t.Fatalf("expected %d items, got %d", len(tt.expected.Items), len(gameState.Items))
			}

			for i, item := range gameState.Items {
				if item.X != tt.expected.Items[i].X || item.Y != tt.expected.Items[i].Y {
					t.Errorf("expected Item = (%v, %v), got (%v, %v)", tt.expected.Items[i].X, tt.expected.Items[i].Y, item.X, item.Y)
				}
				if item.Type != tt.expected.Items[i].Type {
					t.Errorf("expected Item Type = %v, got %v", tt.expected.Items[i].Type, item.Type)
				}
			}
		})
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
