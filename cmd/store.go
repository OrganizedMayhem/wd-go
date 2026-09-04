package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WarpPoint is a named directory destination.
type WarpPoint struct {
	Name string
	Path string
}

// Store persists warp points in the user's .warprc file.
type Store struct {
	path string
}

func newStore() (Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Store{}, fmt.Errorf("find home directory: %w", err)
	}
	return Store{path: filepath.Join(home, ".warprc")}, nil
}

func (s Store) Load() ([]WarpPoint, error) {
	file, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []WarpPoint
	indexes := make(map[string]int)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return nil, fmt.Errorf("invalid warp point at %s:%d", s.path, lineNumber)
		}

		point := WarpPoint{Name: parts[0], Path: parts[1]}
		if index, ok := indexes[point.Name]; ok {
			points[index] = point
			continue
		}
		indexes[point.Name] = len(points)
		points = append(points, point)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read %s: %w", s.path, err)
	}
	return points, nil
}

func (s Store) Get(name string) (string, error) {
	points, err := s.Load()
	if err != nil {
		return "", err
	}
	for _, point := range points {
		if point.Name == name {
			return point.Path, nil
		}
	}
	return "", fmt.Errorf("warp point %q not found", name)
}

func (s Store) Put(name, path string) error {
	if name == "" || strings.Contains(name, ":") {
		return errors.New("warp point names must be non-empty and cannot contain ':'")
	}
	if path == "" {
		return errors.New("warp point paths must be non-empty")
	}

	points, err := s.Load()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	for index, point := range points {
		if point.Name == name {
			points[index].Path = path
			return s.Save(points)
		}
	}
	return s.Save(append(points, WarpPoint{Name: name, Path: path}))
}

func (s Store) Remove(name string) error {
	points, err := s.Load()
	if err != nil {
		return err
	}

	filtered := make([]WarpPoint, 0, len(points))
	for _, point := range points {
		if point.Name != name {
			filtered = append(filtered, point)
		}
	}
	return s.Save(filtered)
}

func (s Store) Clean() error {
	points, err := s.Load()
	if err != nil {
		return err
	}

	filtered := make([]WarpPoint, 0, len(points))
	for _, point := range points {
		if _, err := os.Stat(point.Path); err == nil {
			filtered = append(filtered, point)
		} else if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("check warp point %q: %w", point.Name, err)
		}
	}
	return s.Save(filtered)
}

func (s Store) Save(points []WarpPoint) error {
	directory := filepath.Dir(s.path)
	file, err := os.CreateTemp(directory, ".warprc-")
	if err != nil {
		return fmt.Errorf("create temporary config: %w", err)
	}
	temporaryPath := file.Name()
	defer os.Remove(temporaryPath)

	if err := file.Chmod(0600); err != nil {
		file.Close()
		return fmt.Errorf("set temporary config permissions: %w", err)
	}
	for _, point := range points {
		if _, err := fmt.Fprintf(file, "%s:%s\n", point.Name, point.Path); err != nil {
			file.Close()
			return fmt.Errorf("write temporary config: %w", err)
		}
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close temporary config: %w", err)
	}
	if err := os.Rename(temporaryPath, s.path); err != nil {
		return fmt.Errorf("replace config: %w", err)
	}
	return nil
}
