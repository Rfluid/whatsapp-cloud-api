package message_model

// If you are sending media with link key you can store it in cache
// with this specifications.
//
// See more at https://developers.facebook.com/docs/whatsapp/cloud-api/guides/send-messages#media-messages.
type MediaCacheControl struct {
	CacheControl string `json:"Cache-Control,omitempty"` // Can be either max-age=n (n is time in cache in seconds), no-cache (can be loaded in cache but must be reloaded if Last-Modified is different than a previous response), no-store (no cache then), private (no cache because file is only for the receiver).
	LastModified string `json:"Last-Modified,omitempty"` // Date string that specifies file last modification.
	ETag         string `json:"ETAG,omitempty"`          // Specifies file version.
}

func (c *MediaCacheControl) ToMap() map[string]string {
	mapped := make(map[string]string)
	if c.CacheControl != "" {
		mapped["Cache-Control"] = c.CacheControl
	}
	if c.LastModified != "" {
		mapped["Last-Modified"] = c.LastModified
	}
	if c.ETag != "" {
		mapped["ETAG"] = c.ETag
	}
	return mapped
}
