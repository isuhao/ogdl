// Copyright 2012-2014, Rolf Veen and contributors.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ogdl is used to process OGDL, the Ordered Graph Data Language.
//
// OGDL is a textual format to write trees or graphs of text, where
// indentation and spaces define the structure. Here is an example:
//
//    network
//      ip 192.168.1.100
//      gw 192.168.1.9
//
// The languange is simple, either in its textual representation or its
// number of productions (the specification rules), allowing for compact
// implementations.
//
// OGDL character streams are normally formed by Unicode characters, and
// encoded as UTF-8 strings, but any encoding that is ASCII transparent
// is compatible with the specification and the implementations.
//
// See the full spec at http://ogdl.org.
//
// Installation
//
// To install this package just do:
//
//     go get github.com/rveen/ogdl
//
// An example
//
// If we have a text file 'config.g' containing:
//
//    eth0
//      ip
//        192.168.1.1
//      gateway
//        192.168.1.10
//      mask
//        255.255.255.0
//      timeout
//        20
//
// then,
//
//    g := ogdl.ParseFile("config.g")
//    ip,_ := g.GetString("eth0.ip")
//    to,_ := g.GetInt64("eth0.timeout")
//
//    println("ip:",ip,", timeout:",to)
//
// will print
//
//    ip: 192.168.1.1, timeout: 20
//
// The configuration file can be written in a conciser way:
//
//    eth0
//      ip      192.168.1.1
//      gateway 192.168.1.10
//      mask    255.255.255.0
//      timeout 20
//
// A template example
//
// The package includes a template processor. It takes an arbitrary input stream
// with some variables in it, and produces an output stream with the variables
// resolved out of a Graph object which acts as context.
//
// For example (given the previous config file):
//
//     g := ogdl.ParseFile("config.g")
//     t := ogdl.NewTemplate("The gateway's IP is $eth0.gateway.")
//     b := t.Process(g)
//
// string(b) is then:
//
//     The gateway's IP is 192.168.1.10.
//
package ogdl
