package util

import (
	"os"
	"syscall"
)

// CheckProcessExist 检查进程pid是否存在，如果存在的话，返回true
// linux有效
func CheckProcessExist(pid int) bool {
	// 查询这个pid
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// 给进程发送signal 0, 如果返回nil，代表进程存在, 否则进程不存在
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}
