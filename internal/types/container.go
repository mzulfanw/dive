package types

// Container represents a running container.
type Container struct {
	ID      string `json:"ID"`
	Image   string `json:"Image"`
	Names   string `json:"Names"`
	Command string `json:"Command"`
	Status  string `json:"Status"`
}
