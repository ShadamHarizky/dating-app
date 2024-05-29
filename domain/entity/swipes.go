package entity

type Swipes struct {
	ProfileId uint64 `json:"profile_id"`
	Direction string `json:"direction"`
}
