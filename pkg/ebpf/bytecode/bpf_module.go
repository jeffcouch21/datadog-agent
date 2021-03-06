// +build linux_bpf

package bytecode

import (
	"fmt"
)

// ReadBPFModule from the asset file
func ReadBPFModule(bpfDir string, debug bool) (AssetReader, error) {
	file := "pkg/ebpf/c/tracer-ebpf.o"
	if debug {
		file = "pkg/ebpf/c/tracer-ebpf-debug.o"
	}

	ebpfReader, err := GetReader(bpfDir, file)
	if err != nil {
		return nil, fmt.Errorf("couldn't find asset: %s", err)
	}

	return ebpfReader, nil
}

// ReadOffsetBPFModule from the asset file
func ReadOffsetBPFModule(bpfDir string, debug bool) (AssetReader, error) {
	file := "pkg/ebpf/c/offset-guess.o"
	if debug {
		file = "pkg/ebpf/c/offset-guess-debug.o"
	}

	ebpfReader, err := GetReader(bpfDir, file)
	if err != nil {
		return nil, fmt.Errorf("couldn't find asset: %s", err)
	}

	return ebpfReader, nil
}
