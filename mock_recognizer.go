package imgr

type MockRecognizer struct{}

func NewMockRecognizer() *MockRecognizer {
	return &MockRecognizer{}
}

func (r *MockRecognizer) Recognize(image []byte) (string, error) {
	return "aThing", nil
}
