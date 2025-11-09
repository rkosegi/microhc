/*
Copyright 2025 Richard Kosegi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		checkUrl string
		to       time.Duration
		okCode   int
		silent   bool
	)

	flag.StringVar(&checkUrl, "url", "", "URL to check")
	flag.DurationVar(&to, "duration", 3*time.Second, "Connection timeout")
	flag.IntVar(&okCode, "ok-code", http.StatusOK, "HTTP OK code")
	flag.BoolVar(&silent, "silent", false, "Silent mode (response body not printed")
	flag.Parse()

	if len(checkUrl) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "URL is missing")
		flag.PrintDefaults()
		os.Exit(1)
	}

	hc := http.Client{
		Timeout: to,
	}

	resp, err := hc.Get(checkUrl)
	if err != nil {
		log.Fatal(err)
	}

	if !silent {
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	if resp.StatusCode != okCode {
		log.Fatalf("Unexpected status code %d, wanted %d", resp.StatusCode, okCode)
	}
}
