# ipfixmessage


[![Go Report Card](https://goreportcard.com/badge/github.com/adrianuswarmenhoven/ipfix)](https://goreportcard.com/report/github.com/adrianuswarmenhoven/ipfix)
[![GoDoc](https://godoc.org/github.com/adrianuswarmenhoven/ipfix?status.svg)](https://godoc.org/github.com/adrianuswarmenhoven/ipfix)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/adrianuswarmenhoven/ipfix/blob/master/LICENSE.md)

Full implementation of IPFIX in Go. RFC7011 and RFC7012 (basiclist, subtemplatelist, subtemplatemultilist)

Before compiling, use 'go generate' to pull in some field id's.

Still WIP, but correctly creating and parsing IPFIX messages.

TODO:
    - Sessions (IPFIX has some knowledge about sessions)
    - Nice marshalling of structs
    - Examples