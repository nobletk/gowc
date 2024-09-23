# GoWC

A Command line word count tool written in go.

## Features

Prints character, newline, word, and byte counts for a file.

Reads from the file into a fixed buffer size.


## Usage 

The tool takes the file path as an argument like:

```shell
gowc [OPTIONS] [FILE]
# or with temp binary build
/tmp/bin/gowc [OPTIONS] [FILE]
```

Or you can use `cat` to pipe the content of the file into the tool via stdin like:

```shell
cat [FILE] | gowc [OPTIONS]
# or with temp binary build
cat [FILE] | /tmp/bin/gowc [OPTIONS]
```

The options below may be used to select which counts are printed:

* `-c` or `--bytes`: prints the byte counts

* `-m` or `--chars`: prints the character counts

* `-l` or `--lines`: prints the newline counts

* `-w` or `--chars`: prints the word counts

* `-h` or `--help`: prints usage and options

* If no `OPTIONS` are provided, the tool will always print the count in the following order:

```shell
# newlines, words, bytes [FILE]
  7145  58164  342190  test.txt
```

## Getting started

### Clone the repo

```shell
git clone https://github.com/nobletk/gowc

# then build the binary

make build
```

### Go
```shell
go install https://github.com/nobletk/gowc/cmd/gowc@latest
```
