package requests


type DeleteRequest struct {
	Id string `validate:"required" json:"id,omitempty"`
}

