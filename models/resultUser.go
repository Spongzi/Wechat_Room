package models

type CheckUser struct {
	Name    string `json:"name" db:"name"`
	Sing    string `json:"sing" db:"sing"`
	Photo   string `json:"photo" db:"photo"`
	Address string `json:"address" db:"address"`
	Tel     int    `json:"tel" db:"tel"`
	Account string `json:"account" db:"account"`
	Status  string `json:"status" db:"status"`
}
