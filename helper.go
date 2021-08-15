package imgr

import (
	"encoding/base64"
)

func b64Image(image []byte) string {
	b64img := base64.StdEncoding.EncodeToString(image)
	return b64img
}
