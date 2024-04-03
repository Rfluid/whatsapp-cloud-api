package media_model

import (
	"bytes"
	"mime/multipart"
	"net/url"

	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

type Upload struct {
	File                          string                          `json:"file"` // Path to file.
	Type                          common_model.SupportedMimeTypes `json:"type"`
	BufferBytes                   []byte                          `json:"buffer_bytes"`
	common_model.MessagingProduct                                 // Use SetDefault() method to set this property.
}

func (u *Upload) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("file", string(u.File))
	formData.Set("type", string(u.Type))
	formData.Set("messaging_product", u.MessagingProduct.MessagingProduct)

	return formData
}

func (u *Upload) ToFormValues() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	part, err := w.CreateFormFile("file", u.File)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	part.Write(u.BufferBytes)

	err = w.WriteField("type", string(u.Type))
	if err != nil {
		return &bytes.Buffer{}, err
	}
	err = w.WriteField("messaging_product", u.MessagingProduct.MessagingProduct)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	err = w.Close()
	if err != nil {
		return &bytes.Buffer{}, err
	}

	return buf, nil
}
