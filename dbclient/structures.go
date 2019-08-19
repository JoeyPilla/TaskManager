package dbclient

// Config type
type Config struct {
	User string `json:"user"`
}

// Entry type
type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}
