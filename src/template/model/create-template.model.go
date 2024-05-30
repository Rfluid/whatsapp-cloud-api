package template_model

type CreateTemplate struct {
	Name                string                    `json:"name"`
	Category            TemplateCategories        `json:"category"`
	AllowCategoryChange *bool                     `json:"allow_category_change,omitempty"`
	Language            string                    `json:"language"`
	Components          []CreateTemplateComponent `json:"components"`
}
