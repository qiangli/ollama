//go:build darwin && !cgo

package discover

import "testing"

func TestNoCGODarwinDiscovery(t *testing.T) {
	memory, err := GetCPUMem()
	if err != nil {
		t.Fatal(err)
	}
	if memory.TotalMemory == 0 || memory.FreeMemory > memory.TotalMemory {
		t.Fatalf("invalid memory discovery: %+v", memory)
	}
	cpus := GetCPUDetails()
	if len(cpus) != 1 || cpus[0].CoreCount <= 0 || cpus[0].ThreadCount < cpus[0].CoreCount || cpus[0].EfficiencyCoreCount > cpus[0].CoreCount {
		t.Fatalf("invalid CPU discovery: %+v", cpus)
	}
	if IsNUMA() {
		t.Fatal("Darwin must not enable NUMA")
	}
}
