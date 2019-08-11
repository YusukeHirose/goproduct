package models

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Status  int      `json:"status"`
	Message []string `json:"message"`
}
