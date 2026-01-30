// Handles content validation.
package template

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

type UseTemplate struct {
	Name       string                `json:"name"`
	Language   common.Language `json:"language"`
	Components []Component           `json:"components"`
}
