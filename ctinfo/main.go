package ctinfo

type CTInfo struct {
	CT         string
	Extensions []string
	Filetype   string
}

var fileData = []CTInfo{
	{
		CT: "image/png",
		Extensions: []string{
			"png",
		},
		Filetype: "image",
	},
	{
		CT: "image/jpeg",
		Extensions: []string{
			"jpg",
			"jpeg",
		},
		Filetype: "image",
	},
	{
		CT: "image/gif",
		Extensions: []string{
			"gif",
		},
		Filetype: "image",
	},
}

func GetCTInfo(ct string) *CTInfo {
	for _, file := range fileData {
		if ct == file.CT {
			return &file
		}
	}
	return &CTInfo{
		CT:         ct,
		Extensions: []string{},
		Filetype:   "",
	}
}

func GetFiletype(ct string) string {
	return GetCTInfo(ct).Filetype
}

func GetExtensions(ct string) []string {
	return GetCTInfo(ct).Extensions
}
