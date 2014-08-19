// Copyright 2013 Google. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

// +build !appengine

package main

import (
	"log"
	"net/url"
	"os"

	"github.com/shopkeep/go-rpcgen/example_ae/whoami"
	"github.com/shopkeep/go-rpcgen/webrpc"
)

func main() {
	for _, arg := range os.Args[1:] {
		url, err := url.Parse(arg)
		if err != nil {
			log.Printf("invald url %q: %s", arg, err)
			continue
		}

		svc := whoami.NewWhoamiServiceWebClient(webrpc.JSON, url)

		in, out := whoami.Empty{}, whoami.YouAre{}
		if err := svc.Whoami(&in, &out); err != nil {
			log.Printf("whoami(%q): %s", url, err)
			continue
		}
		log.Printf("You are %s", *out.IpAddr)
	}
}
