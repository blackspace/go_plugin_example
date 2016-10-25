// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("my_plugin.so")
	if err != nil {
		log.Fatalf("plugin.Open failed: %v", err)
	}


	A, err := p.Lookup("A")
	if err != nil {
		log.Fatalf(`Lookup("A") failed: %v`, err)
	}
	log.Println(*A.(*int))


	Fhello, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}


	Fhello.(func())()
}
