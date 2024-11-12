package model

type Item struct {
	Id     uint64  `json:"id"`
	Name   string  `json:"name"`
	Weight float32 `json:"weigth"`
}
