package models

// Password represents the requisition format to password update
type Password struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
