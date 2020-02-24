// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}