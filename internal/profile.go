package profile

// Profile represents a clinks profile
type Profile struct {
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

// Link represents a single link entry
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
