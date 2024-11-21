package errorhandler

// Detail ...
type Detail struct {
	Field     string `json:"field,omitempty"`
	ID        string `json:"id,omitempty"`
	Info      string `json:"info,omitempty"`
	Max       string `json:"max,omitempty"`
	Min       string `json:"min,omitempty"`
	MaxLength string `json:"max_length,omitempty"`
	MinLength string `json:"min_length,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Type      string `json:"type,omitempty"`
}
