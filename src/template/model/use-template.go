// Handles content validation.
package template_model

import (
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

type UseTemplate struct {
	Name       string                `json:"name"`
	Language   common_model.Language `json:"language"`
	Components []Component           `json:"components"`
}
