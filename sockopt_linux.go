// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"os"
	"syscall"
)

func (opt *opt) setKeepAliveProbes(max int) error {
	fd, err := opt.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, max))
}

func (opt *opt) setCork(on bool) error {
	fd, err := opt.sysfd()
	if err != nil {
		return err
	}
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_CORK, boolint(on)))
}