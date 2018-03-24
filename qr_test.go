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
	assert.NoError(t, CharWriteFile("file_qr_half_block", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, CharWriteFile("test/file_qr_half_block", L, "https://github.com/WindomZ/qrw"))
}

func TestBlockWrite(t *testing.T) {
	assert.NoError(t, BlockWrite(os.Stdout, L, "https://github.com/WindomZ/qrw"))
	assert.NoError(t, BlockWrite(nil, L, "https://github.com/WindomZ/qrw"))
}

func TestBlockWriteFile(t *testing.T) {
	assert.NoError(t, BlockWriteFile("file_qr_block", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, BlockWriteFile("test/file_qr_block", L, "https://github.com/WindomZ/qrw"))
}

func TestBash(t *testing.T) {
	assert.NoError(t, Bash(L, "https://github.com/WindomZ/qrw"))
}

func TestPNG(t *testing.T) {
	assert.NoError(t, PNG("file_qr.png", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, PNG("test/file_qr.png", L, "https://github.com/WindomZ/qrw"))
}
