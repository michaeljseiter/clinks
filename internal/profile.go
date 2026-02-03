package profile

import (
	"errors"
	"fmt"
	"net/url"
)

// Validate checks if the profile is valid
func (p *Profile) Validate() error {
	if p.Name == "" {
		return errors.New("name is required and must not empty")
	}

	if len(p.Links) == 0 {
		return errors.New("links must be a not-empty array")
	}

	for i, link := range p.Links {
		if err := link.Validate(); err != nil {
			return fmt.Errorf("link[%d]: %w", i, err)
		}
	}

	return nil
}

// Validate check if a link is valid
func (l *Link) Validate() error {
	if l.Label == "" {
		return errors.New("label is required and must not be empty")
	}

	if l.URL == "" {
		return errors.New("url is required and must not be empty")
	}

	parsedURL, err := url.Parse(l.URL)
	if err != nil {
		return fmt.Errorf("invaid url: %w", err)
	}
}

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
