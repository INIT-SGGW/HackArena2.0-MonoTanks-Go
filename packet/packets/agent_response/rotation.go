package agent_response

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Rotation represents the direction of rotation.
// type Rotation uint64

// const (
// 	Left  Rotation = 0
// 	Right Rotation = 1
// )

// // UnmarshalJSON implements custom JSON unmarshalling to handle integer values for Rotation.
// func (r *Rotation) UnmarshalJSON(data []byte) error {
// 	var value int
// 	if err := json.Unmarshal(data, &value); err != nil {
// 		return err
// 	}
// 	switch value {
// 	case 0:
// 		*r = Left
// 	case 1:
// 		*r = Right
// 	default:
// 		return fmt.Errorf("invalid Rotation value: %d", value)
// 	}
// 	return nil
// }

// // MarshalJSON implements custom JSON marshalling to ensure Rotation is serialized as an integer.
// func (r Rotation) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(uint64(r))
// }
