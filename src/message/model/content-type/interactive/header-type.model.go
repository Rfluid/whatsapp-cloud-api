package message_type_interactive_model

import message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"

type HeaderType message_type_common_model.Type

var (
	Document HeaderType = HeaderType(message_type_common_model.Document)
	Image    HeaderType = HeaderType(message_type_common_model.Image)
	Video    HeaderType = HeaderType(message_type_common_model.Video)
	Text     HeaderType = HeaderType(message_type_common_model.Text)
)
