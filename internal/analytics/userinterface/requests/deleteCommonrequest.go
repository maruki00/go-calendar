package requests

type DeleteCommonRequest struct {
	Id string `validate:"required" json:"id,omitempty"`
}
