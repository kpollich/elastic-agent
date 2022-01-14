// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by elastic-agent/internals/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/elastic-agent-poc/internal/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// internal/spec/apm-server.yml
	// internal/spec/endpoint.yml
	// internal/spec/filebeat.yml
	// internal/spec/fleet-server.yml
	// internal/spec/heartbeat.yml
	// internal/spec/metricbeat.yml
	// internal/spec/osquerybeat.yml
	// internal/spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}