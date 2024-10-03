package tile

// import (
// 	"encoding/json"
// 	"testing"
// )

// func TestSerialize(t *testing.T) {
// 	direction := Up
// 	serialized, err := json.Marshal(direction)
// 	if err != nil {
// 		t.Fatalf("Error marshalling direction: %v", err)
// 	}
// 	expected := "0"
// 	if string(serialized) != expected {
// 		t.Fatalf("Expected %s, got %s", expected, string(serialized))
// 	}
// }

// func TestDeserialize(t *testing.T) {
// 	var deserialized Direction
// 	err := json.Unmarshal([]byte("1"), &deserialized)
// 	if err != nil {
// 		t.Fatalf("Error unmarshalling direction: %v", err)
// 	}
// 	if deserialized != Right {
// 		t.Fatalf("Expected %v, got %v", Right, deserialized)
// 	}
// }

// func TestDeserializeInvalid(t *testing.T) {
// 	var deserialized Direction
// 	err := json.Unmarshal([]byte("4"), &deserialized)
// 	if err == nil {
// 		t.Fatalf("Expected error, got nil")
// 	}
// }

// func TestDeserializeInvalidType(t *testing.T) {
// 	var deserialized Direction
// 	err := json.Unmarshal([]byte("\"1\""), &deserialized)
// 	if err == nil {
// 		t.Fatalf("Expected error, got nil")
// 	}
// }

// func TestDeserializeInvalidType2(t *testing.T) {
// 	var deserialized Direction
// 	err := json.Unmarshal([]byte("Up"), &deserialized)
// 	if err == nil {
// 		t.Fatalf("Expected error, got nil")
// 	}
// }
