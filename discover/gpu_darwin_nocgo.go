//go:build darwin && !cgo

package discover

import (
	"encoding/binary"
	"fmt"
	"runtime"

	"golang.org/x/sys/unix"
)

// Darwin exposes both 32-bit and 64-bit unsigned sysctl counters.
func sysctlMemoryValue(name string) (uint64, error) {
	b, err := unix.SysctlRaw(name)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	switch len(b) {
	case 4:
		return uint64(binary.NativeEndian.Uint32(b)), nil
	case 8:
		return binary.NativeEndian.Uint64(b), nil
	default:
		return 0, fmt.Errorf("%s: unexpected sysctl size %d", name, len(b))
	}
}

func GetCPUMem() (memInfo, error) {
	total, err := sysctlMemoryValue("hw.memsize")
	if err != nil {
		return memInfo{}, err
	}
	pageSize, err := sysctlMemoryValue("hw.pagesize")
	if err != nil {
		return memInfo{}, err
	}
	free, err := sysctlMemoryValue("vm.page_free_count")
	if err != nil {
		return memInfo{}, err
	}
	speculative, _ := sysctlMemoryValue("vm.page_speculative_count")
	// Inactive pages have no portable Darwin sysctl. Report a conservative
	// available-memory estimate; dynamic swap is omitted, as in the cgo build.
	available := (free + speculative) * pageSize
	if available > total {
		available = total
	}
	return memInfo{TotalMemory: total, FreeMemory: available}, nil
}

func GetCPUDetails() []CPU {
	physical, _ := unix.SysctlUint32("hw.physicalcpu")
	logical, _ := unix.SysctlUint32("hw.logicalcpu")
	efficiency, _ := unix.SysctlUint32("hw.perflevel1.physicalcpu")
	if logical == 0 {
		logical = uint32(runtime.NumCPU())
	}
	if physical == 0 {
		physical = logical
	}
	return []CPU{{CoreCount: int(physical), ThreadCount: int(logical), EfficiencyCoreCount: int(efficiency)}}
}

func IsNUMA() bool { return false }
