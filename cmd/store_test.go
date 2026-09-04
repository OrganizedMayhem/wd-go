package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStorePutReplacesExistingPoint(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), ".warprc")
	store := Store{path: configPath}

	if err := store.Put("project", "/first"); err != nil {
		t.Fatal(err)
	}
	if err := store.Put("project", "/second"); err != nil {
		t.Fatal(err)
	}
	if err := store.Put("windows", `C:\workspace`); err != nil {
		t.Fatal(err)
	}

	points, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	want := []WarpPoint{
		{Name: "project", Path: "/second"},
		{Name: "windows", Path: `C:\workspace`},
	}
	if len(points) != len(want) {
		t.Fatalf("loaded %d points, want %d: %#v", len(points), len(want), points)
	}
	for index := range want {
		if points[index] != want[index] {
			t.Errorf("point %d = %#v, want %#v", index, points[index], want[index])
		}
	}

	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := info.Mode().Perm(), os.FileMode(0600); got != want {
		t.Errorf("config permissions = %04o, want %04o", got, want)
	}
}

func TestStoreLoadRejectsMalformedConfig(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), ".warprc")
	if err := os.WriteFile(configPath, []byte("valid:/tmp\nnot-a-warp-point\n"), 0600); err != nil {
		t.Fatal(err)
	}

	_, err := (Store{path: configPath}).Load()
	if err == nil {
		t.Fatal("Load() error = nil, want malformed config error")
	}
	if !strings.Contains(err.Error(), "invalid warp point") {
		t.Errorf("Load() error = %q, want invalid warp point error", err)
	}
}

func TestStoreRemoveAndCleanWriteConsistentConfig(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), ".warprc")
	store := Store{path: configPath}
	livePath := t.TempDir()

	if err := store.Save([]WarpPoint{
		{Name: "live", Path: livePath},
		{Name: "missing", Path: filepath.Join(livePath, "missing")},
	}); err != nil {
		t.Fatal(err)
	}
	if err := store.Clean(); err != nil {
		t.Fatal(err)
	}
	if err := store.Remove("live"); err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(contents); got != "" {
		t.Errorf("config after clean and remove = %q, want empty config", got)
	}
}
