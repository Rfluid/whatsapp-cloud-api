package template

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type GetTemplateResponse struct {
	Data    []Template              `json:"data"`
	Paging  Paging                  `json:"paging"`
	Summary TemplateSummaryResponse `json:"summary,omitempty"`
}

type Paging struct {
	Cursors common.GraphCursors `json:"cursors"`
	Next    string                    `json:"next,omitempty"`
}
