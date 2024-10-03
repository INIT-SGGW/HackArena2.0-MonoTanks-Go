package agent_response

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // MoveDirection represents the direction of movement.
// type MoveDirection uint64

// const (
// 	Forward  MoveDirection = 0
// 	Backward MoveDirection = 1
// )

// // UnmarshalJSON implements custom JSON unmarshalling to handle integer values for MoveDirection.
// func (md *MoveDirection) UnmarshalJSON(data []byte) error {
// 	var value int
// 	if err := json.Unmarshal(data, &value); err != nil {
// 		return err
// 	}
// 	switch value {
// 	case 0:
// 		*md = Forward
// 	case 1:
// 		*md = Backward
// 	default:
// 		return fmt.Errorf("invalid MoveDirection value: %d", value)
// 	}
// 	return nil
// }

// // MarshalJSON implements custom JSON marshalling to ensure MoveDirection is serialized as an integer.
// func (md MoveDirection) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(uint64(md))
// }
