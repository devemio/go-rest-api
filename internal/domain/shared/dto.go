package shared

type NotFoundDto struct {
	Message string `json:"message"`
}

func NewNotFoundDto() *NotFoundDto {
	return &NotFoundDto{
		Message: "Not Found",
	}
}
