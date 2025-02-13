package models
import (
    "errors"
)
// pet status in the store
type Pet_status int

const (
    AVAILABLE_PET_STATUS Pet_status = iota
    PENDING_PET_STATUS
    SOLD_PET_STATUS
)

func (i Pet_status) String() string {
    return []string{"available", "pending", "sold"}[i]
}
func ParsePet_status(v string) (any, error) {
    result := AVAILABLE_PET_STATUS
    switch v {
        case "available":
            result = AVAILABLE_PET_STATUS
        case "pending":
            result = PENDING_PET_STATUS
        case "sold":
            result = SOLD_PET_STATUS
        default:
            return 0, errors.New("Unknown Pet_status value: " + v)
    }
    return &result, nil
}
func SerializePet_status(values []Pet_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
