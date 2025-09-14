package requests

type UpdatePeriodRequest struct {
	Id     string         `validate:"required" json:"id,omitempty"`
	Fields map[string]any `validate:"required" json:"fields,omitempty"`
}
