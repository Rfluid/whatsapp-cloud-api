package template_model

type CreateTemplateResponse struct {
	Id       string           `json:"id"`
	Status   string           `json:"status"`
	Category TemplateCategory `json:"category"`
}
