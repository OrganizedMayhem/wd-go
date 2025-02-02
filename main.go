package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OrganizedMayhem/wd-go/utils"
)

type WarpPoint struct {
	Name      string
	Directory string
}

func readConfigFile(filename string) ([]WarpPoint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var warpPoints []WarpPoint

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			warpPoint := WarpPoint{
				Name:      parts[0],
				Directory: parts[1],
			}
			warpPoints = append(warpPoints, warpPoint)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return warpPoints, nil
}

func addWarpPoint(warpPoints []WarpPoint, name, directory string) ([]WarpPoint, error) {
	// Check if the warp point already exists
	for _, wp := range warpPoints {
		if wp.Name == name {
			return warpPoints, fmt.Errorf("warp point '%s' already exists", name)
		}
	}

	// Add the warp point to the slice
	newWarpPoint := WarpPoint{
		Name:      name,
		Directory: directory,
	}
	warpPoints = append(warpPoints, newWarpPoint)

	return warpPoints, nil
}

func removeWarpPoint(warpPoints []WarpPoint, name string) ([]WarpPoint, error) {
	// Find the index of the warp point to remove
	var index int
	var found bool
	for i, wp := range warpPoints {
		if wp.Name == name {
			index = i
			found = true
			break
		}
	}

	if !found {
		return warpPoints, fmt.Errorf("warp point '%s' not found", name)
	}

	// Remove the warp point from the slice
	warpPoints = append(warpPoints[:index], warpPoints[index+1:]...)

	return warpPoints, nil
}

func main() {
	// Read the configuration file
	warpPoints, err := readConfigFile("/home/sevans/.warprc")
	if err != nil {
		fmt.Println("Error reading configuration file:", err)
		return
	}

	// Add a new warp point
	warpPoints, err = addWarpPoint(warpPoints, "home", "/home/user")
	if err != nil {
		fmt.Println("Error adding warp point:", err)
		return
	}

	// Remove a warp point
	warpPoints, err = removeWarpPoint(warpPoints, "home")
	if err != nil {
		fmt.Println("Error removing warp point:", err)
		return
	}

	// Print the warp points
	for _, wp := range warpPoints {
		fmt.Printf("Warp point: %s -> %s\n", wp.Name, wp.Directory)
	}
}
