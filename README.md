# Numeric Word

Simple library to convert a number from int64 to a word representation.
Currently supports any int64 value.

## Installation

```
$ go get github.com/tgrosinger/numericword
```

## Usage

```
import "github.com/tgrosinger/numericword"

number := int64(12345)
numerString := numericword.ToWord(number)
```
