package template_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type GetTemplateResponse struct {
	Data    []Template              `json:"data"`
	Paging  Paging                  `json:"paging"`
	Summary TemplateSummaryResponse `json:"summary,omitempty"`
}

type Paging struct {
	Cursors common_model.GraphCursors `json:"cursors"`
	Next    string                    `json:"next,omitempty"`
}
