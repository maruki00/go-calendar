package requests

type UpdateRequest struct {
	Id     string         `validate:"required" json:"id,omitempty"`
	Fields map[string]any `validate:"required" json:"fields,omitempty"`
}
