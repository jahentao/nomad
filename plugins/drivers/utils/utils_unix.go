// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package utils

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/unix"
)

// isolateCommand sets the setsid flag in exec.Cmd to true so that the process
// becomes the process leader in a new session and doesn't receive signals that
// are sent to the parent process.
func isolateCommand(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Setsid = true
}

// IsUnixRoot returns true if system is unix and user running is effectively root
func IsUnixRoot() bool {
	return unix.Geteuid() == 0
}
