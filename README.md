# QRCode Writer

> A golang library for outputting QR codes

## Features
- [x] Output to character
  - [x] half character as a block
  - [x] two characters as a block
- [x] Output to terminal
- [x] Output to file
  - [x] characters
  - [ ] png

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
```

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
