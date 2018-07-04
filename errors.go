package figma

import "fmt"

// Error represents an error returned from the Figma API.
type Error struct {
	URL        string
	StatusCode int
	Body       string
}

func (e *Error) Error() string {
	return fmt.Sprintf("figma: %v %v '%.100s'", e.StatusCode, e.URL, e.Body)
}
