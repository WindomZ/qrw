# QRCode Writer

> A golang library for outputting QR codes

[![Build Status](https://travis-ci.org/WindomZ/qrw.svg?branch=master)](https://travis-ci.org/WindomZ/qrw)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/qrw/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/qrw?branch=master)
[![GoDoc](https://godoc.org/github.com/WindomZ/qrw?status.svg)](https://godoc.org/github.com/WindomZ/qrw)

## Features
- [x] Output to text
  - [x] two characters as a block
  - [x] half character as a block
- [x] Output to terminal
  - [x] bash
  - [x] zsh
- [x] Output to file
  - [x] text
  - [x] png
  - [x] jpeg

## Example
As shown in the figure, you can get started quickly with the following example:

![image](https://user-images.githubusercontent.com/14875359/37864413-600aab9a-2fa9-11e8-85ff-9f3c1007bb5f.png)

### BlockWriter
Show a QR block by two characters.

##### Output to `io.Writer`:
```
BlockWrite(os.Stdout, L, "Hello world!")
```

##### Output to `file`:
```
BlockWriteFile("file_path", L, "Hello world!")
```

### HalfBlockWriter
Show a QR block by half character.

##### Output to `io.Writer`:
```
CharWrite(os.Stdout, L, "Hello world!")
```

##### Output to `file`:
```
CharWriteFile("file_path", L, "Hello world!")
```

## Usage
Common functions:
```
func BlockWrite(io.Writer, Level, string) error
func BlockWriteFile(string, Level, string) error
func CharWrite(io.Writer, Level, string) error
func CharWriteFile(string, Level, string) error
func Bash(Level, string) error
func PNG(string, Level, string) error
func JPEG(string, Level, string) error
```

See [document](https://godoc.org/github.com/WindomZ/qrw).

### Level(error correction level)
From least to most tolerant of errors:
- `L` 20% redundant
- `M` 38% redundant
- `Q` 55% redundant
- `H` 65% redundant

The definitions comes from [rsc/qr](https://github.com/rsc/qr/blob/master/qr.go#L23).

## Install
```bash
go get -u github.com/WindomZ/qrw
```

## Related
- Driven by [rsc/qr](https://github.com/rsc/qr)
- Inspired by [mdp/qrterminal](https://github.com/mdp/qrterminal)
