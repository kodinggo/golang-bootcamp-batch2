package httphandler

type Response struct {
	Data     any             `json:"data"`
	Metadata *map[string]any `json:"metadata,omitempty"`
}
