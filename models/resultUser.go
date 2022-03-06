package models

type CheckUser struct {
	Name    string `json:"name" `
	Sing    string `json:"sing" `
	Photo   string `json:"photo" `
	Address string `json:"address" `
	Tel     int    `json:"tel" `
	Status  string `json:"status" `
	LoginId string `json:"login_id"`
}
