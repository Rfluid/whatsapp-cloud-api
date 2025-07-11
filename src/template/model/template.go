package template_model

type Template struct {
	Name       string               `json:"name,omitempty"`
	Components *[]TemplateComponent `json:"components,omitempty"`
	Language   string               `json:"language,omitempty"`
	Status     Status               `json:"status,omitempty" validate:"template_status"`
	Category   string               `json:"category,omitempty"`
	Id         string               `json:"id,omitempty"`
}

type TemplateFields string

const (
	Name           TemplateFields = "name"
	Components     TemplateFields = "components"
	Language       TemplateFields = "language"
	TemplateStatus TemplateFields = "status"
	Category       TemplateFields = "category"
	Id             TemplateFields = "id"
)
