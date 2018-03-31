package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestHalfBlockWrite(t *testing.T) {
	assert.NoError(t, HalfBlockWrite(os.Stdout, L, "https://github.com/WindomZ/qrw"))
	assert.NoError(t, HalfBlockWrite(nil, L, "https://github.com/WindomZ/qrw"))
}

func TestHalfBlockWriteFile(t *testing.T) {
	assert.NoError(t, HalfBlockWriteFile("file_qr_half_block", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, HalfBlockWriteFile("test/file_qr_half_block", L, "https://github.com/WindomZ/qrw"))
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

func TestHalfBash(t *testing.T) {
	assert.NoError(t, HalfBash(L, "https://github.com/WindomZ/qrw"))
}

func TestPNG(t *testing.T) {
	assert.NoError(t, PNG("file_qr.png", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, PNG("test/file_qr.png", L, "https://github.com/WindomZ/qrw"))
}

func TestJPEG(t *testing.T) {
	assert.NoError(t, JPEG("file_qr.jpeg", L, "https://github.com/WindomZ/qrw"))
	assert.Error(t, JPEG("test/file_qr.jpeg", L, "https://github.com/WindomZ/qrw"))
}
