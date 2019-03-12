package log

// Log represents a time log file
type Log struct {
	Segments Segment
}

// Segment represents a tim log segment
type Segment struct {
	Project string
	Login   string
	Logout  string
}
