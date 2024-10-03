package game_state

// // Zone represents a zone in the game world.
// type Zone struct {
// 	// The unique index of the zone.
// 	Index uint8 `json:"index"`

// 	// The x-coordinate of the left side of the zone.
// 	X uint64 `json:"x"`

// 	// The y-coordinate of the top side of the zone.
// 	Y uint64 `json:"y"`

// 	// The width of the zone.
// 	Width uint64 `json:"width"`

// 	// The height of the zone.
// 	Height uint64 `json:"height"`

// 	// The current status of the zone.
// 	Status ZoneStatus `json:"status"`
// }

// // ZoneStatus represents the status of a zone.
// type ZoneStatus struct {
// 	Type           string                `json:"type"`
// 	BeingCaptured  *BeingCapturedStatus  `json:"beingCaptured,omitempty"`
// 	Captured       *CapturedStatus       `json:"captured,omitempty"`
// 	BeingContested *BeingContestedStatus `json:"beingContested,omitempty"`
// 	BeingRetaken   *BeingRetakenStatus   `json:"beingRetaken,omitempty"`
// }

// // BeingCapturedStatus represents the status of a zone being captured.
// type BeingCapturedStatus struct {
// 	// The remaining ticks until the zone is captured.
// 	RemainingTicks uint64 `json:"remainingTicks"`

// 	// The ID of the player capturing the zone.
// 	PlayerID string `json:"playerId"`
// }

// // CapturedStatus represents the status of a zone that has been captured.
// type CapturedStatus struct {
// 	// The ID of the player who captured the zone.
// 	PlayerID string `json:"playerId"`
// }

// // BeingContestedStatus represents the status of a zone being contested.
// type BeingContestedStatus struct {
// 	// The ID of the player who captured the zone, if any.
// 	CapturedByID *string `json:"capturedById,omitempty"`
// }

// // BeingRetakenStatus represents the status of a zone being retaken.
// type BeingRetakenStatus struct {
// 	// The remaining ticks until the zone is retaken.
// 	RemainingTicks uint64 `json:"remainingTicks"`

// 	// The ID of the player who previously captured the zone.
// 	CapturedByID string `json:"capturedById"`

// 	// The ID of the player retaking the zone.
// 	RetakenByID string `json:"retakenById"`
// }

// // {
// // 	"x": 4,
// // 	"y": 13,
// // 	"width": 4,
// // 	"height": 4,
// // 	"index": 65,
// // 	"status": {
// // 	  "type": "neutral"
// // 	}
// //   },
// //   {
// // 	"x": 3,
// // 	"y": 3,
// // 	"width": 4,
// // 	"height": 4,
// // 	"index": 66,
// // 	"status": {
// // 	  "type": "neutral"
// // 	}
// //   },
// // {
// // 	"x": 1,
// // 	"y": 1,
// // 	"width": 4,
// // 	"height": 4,
// // 	"index": 65,
// // 	"status": {
// // 	  "remainingTicks": 100,
// // 	  "playerId": "e21b7b97-0451-4800-b1ba-0e5cfc983aa3",
// // 	  "type": "beingCaptured"
// // 	}
// //   }
