package common_model

import "net/url"

type GraphCursors struct {
	After  *string `json:"after,omitempty"`  // Used to paginate.
	Before *string `json:"before,omitempty"` // Used to paginate.
}

func (gba *GraphCursors) BuildQuery(v *url.Values) {
	if gba.After != nil {
		v.Set("after", *gba.After)
	}
	if gba.Before != nil {
		v.Set("before", *gba.Before)
	}
}
