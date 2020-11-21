// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build arm64

package kvm

import (
	"syscall"
	"unsafe"

	"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/platform/ring0"
)

// fpsimdPtr returns a fpsimd64 for the given address.
//
//go:nosplit
func fpsimdPtr(addr *byte) *arch.FpsimdContext {
	return (*arch.FpsimdContext)(unsafe.Pointer(addr))
}

// dieArchSetup initialies the state for dieTrampoline.
//
// The arm64 dieTrampoline requires the vCPU to be set in R1, and the last PC
// to be in R0. The trampoline then simulates a call to dieHandler from the
// provided PC.
//
//go:nosplit
func dieArchSetup(c *vCPU, context *arch.SignalContext64, guestRegs *userRegs) {
	// If the vCPU is in user mode, we set the stack to the stored stack
	// value in the vCPU itself. We don't want to unwind the user stack.
	if guestRegs.Regs.Pstate&ring0.PsrModeMask == ring0.UserFlagsSet {
		regs := c.CPU.Registers()
		context.Regs[0] = regs.Regs[0]
		context.Sp = regs.Sp
		context.Regs[29] = regs.Regs[29] // stack base address
	} else {
		context.Regs[0] = guestRegs.Regs.Pc
		context.Sp = guestRegs.Regs.Sp
		context.Regs[29] = guestRegs.Regs.Regs[29]
		context.Pstate = guestRegs.Regs.Pstate
	}
	context.Regs[1] = uint64(uintptr(unsafe.Pointer(c)))
	context.Pc = uint64(dieTrampolineAddr)
}

// bluepillArchFpContext returns the arch-specific fpsimd context.
//
//go:nosplit
func bluepillArchFpContext(context unsafe.Pointer) *arch.FpsimdContext {
	return &((*arch.SignalContext64)(context).Fpsimd64)
}

// getHypercallID returns hypercall ID.
//
// On Arm64, the MMIO address should be 64-bit aligned.
//
//go:nosplit
func getHypercallID(addr uintptr) int {
	if addr < arm64HypercallMMIOBase || addr >= (arm64HypercallMMIOBase+_AARCH64_HYPERCALL_MMIO_SIZE) {
		return _KVM_HYPERCALL_MAX
	} else {
		return int(((addr) - arm64HypercallMMIOBase) >> 3)
	}
}

// bluepillStopGuest is reponsible for injecting sError.
//
//go:nosplit
func bluepillSetIrq(c *vCPU, irqType, irq, level uint32) {
	kvmIrq := (irqType << _KVM_ARM_IRQ_TYPE_SHIFT) | irq
	cpuIdx1 := uint32(c.id % 256)
	cpuIdx2 := uint32(c.id / 256)

	kvmIrq |= (cpuIdx1 << _KVM_ARM_IRQ_VCPU_SHIFT) | (cpuIdx2 << _KVM_ARM_IRQ_VCPU2_SHIFT)

	irqLevel := irqLevel{
		irq:   kvmIrq,
		level: level,
	}
	if _, _, errno := syscall.RawSyscall( // escapes: no.
		syscall.SYS_IOCTL,
		uintptr(c.machine.fd),
		_KVM_IRQ_LINE,
		uintptr(unsafe.Pointer(&irqLevel))); errno != 0 {
		throw("irq injection failed")
	}
}

// bluepillStopGuest is reponsible for injecting sError.
//
//go:nosplit
func bluepillStopGuest(c *vCPU) {
	/*
	   int kvm_irq = (irqtype << KVM_ARM_IRQ_TYPE_SHIFT) | irq;
	   int cpu_idx1 = cpu % 256;
	   int cpu_idx2 = cpu / 256;

	   kvm_irq |= (cpu_idx1 << KVM_ARM_IRQ_VCPU_SHIFT) |
	              (cpu_idx2 << KVM_ARM_IRQ_VCPU2_SHIFT);

	   return kvm_set_irq(kvm_state, kvm_irq, !!level);
	*/

	/*	if _, _, errno := syscall.RawSyscall( // escapes: no.
			syscall.SYS_IOCTL,
			uintptr(c.fd),
			_KVM_SET_VCPU_EVENTS,
			uintptr(unsafe.Pointer(&vcpuSErrBounce))); errno != 0 {
			throw("sErr injection failed")
		}
	*/
	bluepillSetIrq(c, _KVM_ARM_IRQ_TYPE_CPU, _KVM_ARM_IRQ_CPU_FIQ, 1)
	bluepillSetIrq(c, _KVM_ARM_IRQ_TYPE_CPU, _KVM_ARM_IRQ_CPU_FIQ, 0)
}

// bluepillSigBus is reponsible for injecting sError to trigger sigbus.
//
//go:nosplit
func bluepillSigBus(c *vCPU) {
	if _, _, errno := syscall.RawSyscall( // escapes: no.
		syscall.SYS_IOCTL,
		uintptr(c.fd),
		_KVM_SET_VCPU_EVENTS,
		uintptr(unsafe.Pointer(&vcpuSErrNMI))); errno != 0 {
		throw("sErr injection failed")
	}
}

// bluepillReadyStopGuest checks whether the current vCPU is ready for sError injection.
//
//go:nosplit
func bluepillReadyStopGuest(c *vCPU) bool {
	return true
}
