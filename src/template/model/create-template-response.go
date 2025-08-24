package template_model

type CreateTemplateResponse struct {
	ID       string           `json:"id"`
	Status   Status           `json:"status" validate:"required,template_status"`
	Category TemplateCategory `json:"category" validate:"required,template_category"`
}
