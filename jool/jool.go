package jool

import (
	"bytes"
	"context"
	"errors"
	"os/exec"

	"github.com/gocarina/gocsv"
)

// Framework is a constant from https://github.com/NICMx/Jool/blob/main/src/common/config.c#L165
type Framework uint

const (
	FrameworkNetfilter Framework = 1 << 2
	FrameworkIPTables  Framework = 1 << 3
)

func (f Framework) String() string {
	switch f {
	case FrameworkNetfilter:
		return "netfilter"
	case FrameworkIPTables:
		return "iptables"
	default:
		return ""
	}
}

func (f *Framework) UnmarshalText(text []byte) error {
	switch string(text) {
	case "netfilter":
		*f = FrameworkNetfilter
	case "iptables":
		*f = FrameworkIPTables
	default:
		return errors.New("netfilter and iptables are the only available instance frameworks")
	}
	return nil
}

type Instance struct {
	Name      string    `csv:"Name"`
	Namespace string    `csv:"Namespace"`
	Framework Framework `csv:"Framework"`
}

func Instances(ctx context.Context) ([]Instance, error) {
	var out []Instance
	return out, joolCmd(ctx, &out, "instance", "display")
}

func Stats(ctx context.Context, instance string) (map[Statistic]uint64, error) {
	var out []struct {
		Stat  string `csv:"Stat"`
		Value uint64 `csv:"Value"`
	}
	err := joolCmd(ctx, &out, "-i", instance, "stats", "display")
	if err != nil {
		return nil, err
	}
	stats := make(map[Statistic]uint64, len(out))
	for _, stat := range out {
		id := ParseStatistic(stat.Stat)
		if id == 0 {
			panic(stat.Stat)
		}
		stats[id] = stat.Value
	}
	return stats, nil
}

func joolCmd(ctx context.Context, into any, args ...string) error {
	cmd := exec.CommandContext(ctx, "jool", append(args, "--csv")...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = gocsv.Unmarshal(stdout, into)
	if err != nil {
		_ = cmd.Cancel()
		return err
	}

	err = cmd.Wait()
	if err != nil {
		// Annotate the ExitError with stderr
		var procErr *exec.ExitError
		if errors.As(err, &procErr) {
			procErr.Stderr = stderr.Bytes()
		}
	}
	return err
}
