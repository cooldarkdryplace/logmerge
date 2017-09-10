# logmerge
Tiny tool for merging multiple log files. Reads stream, so should not use a lot of RAM.

## Installation
`go get github.com/cooldarkdryplace/logmerge`

## Usage
`logmerge -i your_input_directory -o output.file -tf Jan  2 15:04:05`

Please refer to time package documentation for more info.
https://golang.org/pkg/time/#pkg-constants
