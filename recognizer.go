package imgr

type Recognizer interface {
	Recognize(image []byte) (string, error)
}
