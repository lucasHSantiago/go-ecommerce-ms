package api

import "embed"

//go:embed "swagger/*"
var StaticSwaggerFS embed.FS
