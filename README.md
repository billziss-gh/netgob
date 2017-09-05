# netgob - gob encoding with channels

[![GoDoc](https://godoc.org/github.com/billziss-gh/netgob/gob?status.svg)](https://godoc.org/github.com/billziss-gh/netgob/gob)

Package netgob is a replacement for Go's gob package that also supports encoding/decoding of channels. Package netgob can be used to marshal channels in networking scenarios; the [netchan](https://github.com/billziss-gh/netchan) library uses netgob for this purpose.

This is accomplished through the use of the new `NetgobEncoder` and `NetgobDecoder` interfaces. When netgob sees a channel during encoding/decoding it calls one of these interfaces. A marshaling layer can implement these interfaces to convert channels to marshaling references and vice-versa.

The same technique could be used to encode/decode functions, although netgob does not support them at this time.

## Usage

To use netgob simply:

```gob
import (
    "github.com/billziss-gh/netgob/gob"
)
```
