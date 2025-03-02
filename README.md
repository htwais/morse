

# morse

A CLI tool that takes a text which contains either standard characters or Morse code and converts it to the respective other representation.


# HowTo

build it with

    $ go build .
    $ go test ./...
    $ go doc ./convert

# Usage

    $ ./morse path/to/input.txt [path/to/output/file]
    
    $ ./morse path/to/input.morse [path/to/output/file]

if no output file is given, stdout is used


# Limitations

Morse input must not contain anything but

-   dot, ascii 0x2e, the morse dit
-   dash, ascii 0x2d, the morse dah
-   space, ascii 0x20, character separator
-   slash, ascii 0x2f, single slash to separate words, two slashes for newline

Text input must be valid UTF8.
Any character that is not representable as morse is silently ignored.
Conversion from morse to utf8 is all lower case.

# TODO
- Nonstandard punctuation?
- non-Latin extensions?
- Cut numbers? 
- Prosigns
- allow for other representations of dit and dah?
- better error messages: add position in file and bad character
- document errors
- test output: convert space to hex to make it more readable
