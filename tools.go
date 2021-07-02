// +build tools

// This file uses the recommended method for tracking developer tools in a go module.
//
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/golangci/golangci-lint"
)
