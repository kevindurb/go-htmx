package main

import (
	"time"
)

type AutoIncr struct {
	ID      uint64
	Created time.Time
}

type Todo struct {
	AutoIncr

	Description string
	Done        bool
}
