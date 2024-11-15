// Copyright 2017 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong64 && linux
// +build loong64,linux

package tcp

const (
	sysSIOCINQ  = 0x541b
	sysSIOCOUTQ = 0x5411
)
