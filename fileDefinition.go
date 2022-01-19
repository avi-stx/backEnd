package main

import "time"

type fileInfo struct {
	id           int64
	name         string
	extension    string
	size         int64
	creationData time.Time
}
