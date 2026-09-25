// Package openapi embeds the bundled OpenAPI documents for the MaxMind web
// services, so Go programs can validate requests and responses against them.
package openapi

import "embed"

// Bundled holds one self-contained OpenAPI document per product, for example
// "bundled/geoip.yaml".
//
//go:embed bundled/*.yaml
var Bundled embed.FS
