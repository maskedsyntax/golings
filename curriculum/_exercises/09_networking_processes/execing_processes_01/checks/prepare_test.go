package prepare

import (
	"errors"
	"slices"
	"testing"
)

func fakeLookPath(name string) (string, error) {
	if name == "ls" {
		return "/bin/ls", nil
	}
	return "", errors.New("executable file not found in $PATH")
}

func TestPrepare(t *testing.T) {
	environ := []string{"PATH=/bin", "LANG=C", "HOME=/home/ada"}
	extra := map[string]string{"LANG": "en_US.UTF-8", "DEBUG": "1", "APP": "x"}
	path, argv, env, err := Prepare("ls", []string{"-a", "-l"}, extra, environ, fakeLookPath)
	if err != nil {
		t.Fatalf("Prepare returned an error: %v", err)
	}
	if path != "/bin/ls" {
		t.Errorf("path = %q, want \"/bin/ls\"", path)
	}
	if want := []string{"ls", "-a", "-l"}; !slices.Equal(argv, want) {
		t.Errorf("argv = %q, want %q", argv, want)
	}
	if want := []string{"PATH=/bin", "LANG=en_US.UTF-8", "HOME=/home/ada", "APP=x", "DEBUG=1"}; !slices.Equal(env, want) {
		t.Errorf("env = %q, want %q", env, want)
	}
	if !slices.Equal(environ, []string{"PATH=/bin", "LANG=C", "HOME=/home/ada"}) {
		t.Errorf("Prepare modified the caller's environ: %q", environ)
	}
}

func TestPrepareNoExtra(t *testing.T) {
	_, argv, env, err := Prepare("ls", nil, nil, []string{"A=1"}, fakeLookPath)
	if err != nil || !slices.Equal(argv, []string{"ls"}) || !slices.Equal(env, []string{"A=1"}) {
		t.Fatalf("Prepare = %q, %q, %v", argv, env, err)
	}
}

func TestPrepareNotFound(t *testing.T) {
	if _, _, _, err := Prepare("nope", nil, nil, nil, fakeLookPath); err == nil {
		t.Fatal("Prepare of a missing tool returned no error")
	}
}
