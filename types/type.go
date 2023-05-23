package types

import "time"

type Order struct {
	Id        int
	Uid       int
	Weight    float64
	CreatedAt time.Time
}
