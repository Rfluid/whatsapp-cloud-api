package common_model

import (
	"errors"
	"strings"
)

type SupportedMimeTypes string

const (
	AudioAac                                                             SupportedMimeTypes = "audio/aac"
	AudioMp4                                                             SupportedMimeTypes = "audio/mp4"
	AudioMpeg                                                            SupportedMimeTypes = "audio/mpeg"
	AudioAmr                                                             SupportedMimeTypes = "audio/amr"
	AudioOgg                                                             SupportedMimeTypes = "audio/ogg"
	TextPlain                                                            SupportedMimeTypes = "text/plain"
	ApplicationPdf                                                       SupportedMimeTypes = "application/pdf"
	ApplicationVndMsPowerpoint                                           SupportedMimeTypes = "application/vnd.ms-powerpoint"
	ApplicationMsword                                                    SupportedMimeTypes = "application/msword"
	ApplicationVndMsExcel                                                SupportedMimeTypes = "application/vnd.ms-excel"
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument   SupportedMimeTypes = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation SupportedMimeTypes = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet         SupportedMimeTypes = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	ImageJpeg                                                            SupportedMimeTypes = "image/jpeg"
	ImagePng                                                             SupportedMimeTypes = "image/png"
	VideoMp4                                                             SupportedMimeTypes = "video/mp4"
	Video3gp                                                             SupportedMimeTypes = "video/3gp"
	ImageWebp                                                            SupportedMimeTypes = "image/webp"
)

func ParseMimeType(mimeType string) (SupportedMimeTypes, error) {
	switch strings.ToLower(mimeType) {
	case "audio/aac":
		return AudioAac, nil
	case "audio/mp4":
		return AudioMp4, nil
	case "audio/mpeg":
		return AudioMpeg, nil
	case "audio/amr":
		return AudioAmr, nil
	case "audio/ogg":
		return AudioOgg, nil
	case "text/plain":
		return TextPlain, nil
	case "application/pdf":
		return ApplicationPdf, nil
	case "application/vnd.ms-powerpoint":
		return ApplicationVndMsPowerpoint, nil
	case "application/msword":
		return ApplicationMsword, nil
	case "application/vnd.ms-excel":
		return ApplicationVndMsExcel, nil
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument, nil
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation, nil
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet, nil
	case "image/jpeg":
		return ImageJpeg, nil
	case "image/png":
		return ImagePng, nil
	case "video/mp4":
		return VideoMp4, nil
	case "video/3gp":
		return Video3gp, nil
	case "image/webp":
		return ImageWebp, nil
	default:
		return TextPlain, errors.New("unsupported type")
	}
}
