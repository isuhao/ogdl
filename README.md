# OGDL for Go

Package ogdl is used to process [OGDL](http://ogdl.org), the Ordered Graph Data Language.

OGDL is a textual format to write trees or graphs of text, where indentation and spaces define the structure. Here is an example:

    network
      ip 192.168.1.100
      gw 192.168.1.9

The languange is simple, either in its textual representation or its number of productions (the specification rules), allowing for compact implementations.

OGDL character streams are normally formed by Unicode characters, and encoded as UTF-8 strings, but any encoding that is ASCII transparent is compatible with the specification and the implementations.

## Installation

    go get github.com/rveen/ogdl

## Example: a configuration file

If we have a text file 'conf.g' like this:

    eth0
      ip
        192.168.1.1
      gateway
        192.168.1.10
      mask
        255.255.255.0
      timeout
        20
then,

    g := ogdl.ParseFile("conf.g")
    ip,_ := g.GetString("eth0.ip")
    to,_ := g.GetInt("eth0.timeout")
    println("ip:",ip,", timeout:",to)

will print

    ip: 192.168.1.1, timeout: 20

## Documentation

The [documentation](http://godoc.org/github.com/rveen/ogdl) of the package is kindly generated by godoc.org.
