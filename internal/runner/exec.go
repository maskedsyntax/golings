package runner

import (
	"context"
	"errors"
	"os/exec"
	"time"
)

// outputLimit bounds each captured stream (spec §3.4: "last 64 KiB").
const outputLimit = 64 << 10

// testJSONLimit bounds `go test -json` stdout, which is parsed before display
// and so may need more room than displayed output.
const testJSONLimit = 4 << 20

// tailBuffer keeps only the last limit bytes written to it.
type tailBuffer struct {
	limit     int
	buf       []byte
	truncated bool
}

func newTailBuffer(limit int) *tailBuffer { return &tailBuffer{limit: limit} }

func (b *tailBuffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	if len(b.buf) > 2*b.limit {
		b.buf = append([]byte(nil), b.buf[len(b.buf)-b.limit:]...)
		b.truncated = true
	}
	return len(p), nil
}

func (b *tailBuffer) Bytes() []byte {
	if len(b.buf) > b.limit {
		b.truncated = true
		return b.buf[len(b.buf)-b.limit:]
	}
	return b.buf
}

func (b *tailBuffer) Truncated() bool {
	b.Bytes()
	return b.truncated
}

type execResult struct {
	ExitCode  int
	Stdout    []byte
	Stderr    []byte
	Truncated bool
	TimedOut  bool // ctx deadline exceeded
	Canceled  bool // ctx canceled for another reason (superseded, Ctrl+C)
	StartErr  error
	Duration  time.Duration
}

// execCommand runs bin with args in dir using exactly env. The whole process
// tree is killed when ctx ends.
func execCommand(ctx context.Context, dir string, env []string, stdoutLimit int, bin string, args ...string) execResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Dir = dir
	cmd.Env = env
	stdout, stderr := newTailBuffer(stdoutLimit), newTailBuffer(outputLimit)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	cmd.WaitDelay = 2 * time.Second

	g := newProcGroup()
	g.prepare(cmd)
	cmd.Cancel = func() error { return g.kill(cmd) }
	defer g.close()

	res := execResult{ExitCode: -1}
	if err := cmd.Start(); err != nil {
		res.StartErr = err
		res.Duration = time.Since(start)
		return res
	}
	g.started(cmd)
	err := cmd.Wait()
	res.Duration = time.Since(start)
	res.Stdout, res.Stderr = stdout.Bytes(), stderr.Bytes()
	res.Truncated = stdout.Truncated() || stderr.Truncated()
	if cmd.ProcessState != nil {
		res.ExitCode = cmd.ProcessState.ExitCode()
	}
	if ctxErr := ctx.Err(); ctxErr != nil {
		if errors.Is(ctxErr, context.DeadlineExceeded) {
			res.TimedOut = true
		} else {
			res.Canceled = true
		}
		return res
	}
	var exitErr *exec.ExitError
	if err != nil && !errors.As(err, &exitErr) && !errors.Is(err, exec.ErrWaitDelay) {
		res.StartErr = err
	}
	return res
}
