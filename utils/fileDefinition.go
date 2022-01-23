package utils

import "time"

type fileInfo struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Extension    string    `json:"extension"`
	Size         int64     `json:"size"`
	CreationDate time.Time `json:"creationDate"`
}
