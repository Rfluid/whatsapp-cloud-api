package media

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"

	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

type UploadPayload struct {
	FileData                      io.Reader                       // This will handle the media data
	FileName                      string                          // File name for the media being uploaded	Type                          common.SupportedMimeTypes `json:"type"`
	Type                          common.SupportedMimeTypes `json:"type"`
	common.MessagingProduct                                 // Use SetDefault() method to set this property.
}

// Creates the request body for the file upload
func (u *UploadPayload) CreateForm() (*bytes.Buffer, string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a MIMEHeader to set Content-Disposition and Content-Type
	headers := textproto.MIMEHeader{}
	headers.Set("Content-Type", string(u.Type))
	headers.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, u.FileName))

	part, err := writer.CreatePart(headers)
	if err != nil {
		return nil, "", err
	}

	// Write file data to the multipart
	_, err = io.Copy(part, u.FileData)
	if err != nil {
		return nil, "", err
	}

	// Add other fields
	err = writer.WriteField("type", string(u.Type))
	if err != nil {
		return nil, "", err
	}

	err = writer.WriteField("messaging_product", u.MessagingProduct.MessagingProduct)
	if err != nil {
		return nil, "", err
	}

	// Close the writer to finalize the request body
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
