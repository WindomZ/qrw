package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestCharWrite(t *testing.T) {
	assert.NoError(t, CharWrite(os.Stdout, L, "https://github.com/WindomZ/qrw"))
	assert.NoError(t, CharWrite(nil, L, "https://github.com/WindomZ/qrw"))
}

func TestCharWriteFile(t *testing.T) {
	assert.NoError(t, CharWriteFile("QR", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, CharWriteFile("test/QR", L, "https://github.com/WindomZ/qrw"))
}
