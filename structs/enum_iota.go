package structs

import "time"

type status int

// status iota (0, 1, 2)
const (
	UNKNOWN status = iota
	TODO
	DONE
)

// Task struct use enum(const)
type Task struct {
	title  string
	status status
	due    *time.Time
}

// ByteSize type float64
type ByteSize float64

// ByteSize iota
const (
	KB ByteSize = 1 << (10 * (1 + iota))
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
