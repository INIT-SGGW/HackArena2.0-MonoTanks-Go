package tile

// import (
// 	"encoding/json"
// 	"errors"
// )

// // Direction represents the four cardinal directions.
// type Direction uint64

// const (
// 	// Represents upward direction.
// 	Up Direction = iota

// 	// Represents rightward direction.
// 	Right

// 	// Represents downward direction.
// 	Down

// 	// Represents leftward direction.
// 	Left
// )

// // MarshalJSON implements the json.Marshaler interface.
// func (d Direction) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(uint64(d))
// }

// // UnmarshalJSON implements the json.Unmarshaler interface.
// func (d *Direction) UnmarshalJSON(data []byte) error {
// 	var value uint64
// 	if err := json.Unmarshal(data, &value); err != nil {
// 		return err
// 	}

// 	switch value {
// 	case 0, 1, 2, 3:
// 		*d = Direction(value)
// 		return nil
// 	default:
// 		return errors.New("invalid direction value")
// 	}
// }
