package ctinfo

import (
	"mime"
	"strings"
)

type CTInfo struct {
	CT         string
	Extensions []string
	Filetype   string
}

func GetCTInfo(ct string) *CTInfo {
	exts, _ := mime.ExtensionsByType(ct)

	filetype := strings.Split(ct, "/")[0]

	return &CTInfo{
		CT:         ct,
		Extensions: exts,
		Filetype:   filetype,
	}
}

func GetFiletype(ct string) string {
	return GetCTInfo(ct).Filetype
}

func GetExtensions(ct string) []string {
	return GetCTInfo(ct).Extensions
}
