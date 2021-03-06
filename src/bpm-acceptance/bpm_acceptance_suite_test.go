// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package bpm_acceptance_test

import (
	"flag"
	"log"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBpmAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BpmAcceptance Suite")
}

var (
	agentURI string
	client   *http.Client
)

func init() {
	flag.StringVar(&agentURI, "agent-uri", "", "http address of the test-server")
	flag.Parse()

	if agentURI == "" {
		log.Fatal("Agent URI must be provided")
	}
}

var _ = BeforeSuite(func() {
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
})
