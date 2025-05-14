package template_model

type ComponentFormat string

const (
	ImageFormat    ComponentFormat = "IMAGE"
	VideoFormat    ComponentFormat = "VIDEO"
	StickerFormat  ComponentFormat = "STICKER"
	DocumentFormat ComponentFormat = "DOCUMENT"
	AudioFormat    ComponentFormat = "AUDIO"
	TextFormat     ComponentFormat = "TEXT"
)
