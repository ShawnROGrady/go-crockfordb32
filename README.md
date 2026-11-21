# go-crockfordb32

[![Go Reference](https://pkg.go.dev/badge/github.com/ShawnROGrady/go-crockfordb32.svg)](https://pkg.go.dev/github.com/ShawnROGrady/go-crockfordb32)

A Go implementation of [Douglas Crockford's Base32 encoding](https://www.crockford.com/base32.html).

## Overview

Crockford's Base32 uses a 32-character alphabet (`0-9` and `a-z`, excluding `i`, `l`, `o`, and `u`) chosen to be unambiguous when written or read by humans. Padding is omitted.

This package is a thin wrapper around the standard library's `encoding/base32`, exposing the encoding with the Crockford alphabet as package-level functions.

## Requirements

Go 1.22 or later.

## Installation

```sh
go get github.com/ShawnROGrady/go-crockfordb32
```

## Usage

```go
import "github.com/ShawnROGrady/go-crockfordb32/crockfordb32"

encoded := crockfordb32.EncodeToString([]byte("hello"))

decoded, err := crockfordb32.DecodeString(encoded)
```

If you need the underlying `*base32.Encoding` directly (e.g. to pass to another library), use `crockfordb32.Encoding()`.
