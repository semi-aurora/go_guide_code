package main

import "syscall"

const LinuxRebootMagic1 uintptr = 0xfee1dead
const LinuxRebootMagic2 uintptr = 672274793
const LinuxRebootCmdRestart uintptr = 0x1234567

func main() {
	_, _, _ = syscall.Syscall(syscall.SYS_REBOOT,
		LinuxRebootMagic1,
		LinuxRebootMagic2,
		LinuxRebootCmdRestart)
}