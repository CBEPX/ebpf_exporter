package exporter

import (
	"fmt"

	"github.com/iovisor/gobpf/bcc"
)

// attachKprobes attaches functions to their kprobles in provided module
func attachKprobes(module *bcc.Module, kprobes map[string]string) error {
	for kprobeName, targetName := range kprobes {
		target, err := module.LoadKprobe(targetName)
		if err != nil {
			return fmt.Errorf("failed to load target %q: %s", targetName, err)
		}

		err = module.AttachKprobe(kprobeName, target)
		if err != nil {
			return fmt.Errorf("failed to attach kprobe %q to %q: %s", kprobeName, targetName, err)
		}
	}

	return nil
}

// attachKretprobes attaches functions to their kretprobles in provided module
func attachKretprobes(module *bcc.Module, kretprobes map[string]string) error {
	for kretprobeName, targetName := range kretprobes {
		target, err := module.LoadKprobe(targetName)
		if err != nil {
			return fmt.Errorf("failed to load target %q: %s", targetName, err)
		}

		err = module.AttachKretprobe(kretprobeName, target)
		if err != nil {
			return fmt.Errorf("failed to attach kretprobe %q to %q: %s", kretprobeName, targetName, err)
		}
	}

	return nil
}

// attachTracepoints attaches functions to their tracepoints in provided module
func attachTracepoints(module *bcc.Module, tracepoints map[string]string) error {
	for tracepointName, targetName := range tracepoints {
		target, err := module.LoadTracepoint(targetName)
		if err != nil {
			return fmt.Errorf("failed to load target %q: %s", targetName, err)
		}

		err = module.AttachTracepoint(tracepointName, target)
		if err != nil {
			return fmt.Errorf("failed to attach tracepoint %q to %q: %s", tracepointName, targetName, err)
		}
	}

	return nil
}

// attachRawTracepoints attaches functions to their tracepoints in provided module
func attachRawTracepoints(module *bcc.Module, tracepoints map[string]string) error {
	for tracepointName, targetName := range tracepoints {
		target, err := module.LoadRawTracepoint(targetName)
		if err != nil {
			return fmt.Errorf("failed to load target %q: %s", targetName, err)
		}

		err = module.AttachRawTracepoint(tracepointName, target)
		if err != nil {
			return fmt.Errorf("failed to attach raw tracepoint %q to %q: %s", tracepointName, targetName, err)
		}
	}

	return nil
}
