package utils

import (
	"fmt"
)

func AddWarpPoint(warpPoints []WarpPoint, name, directory string) ([]WarpPoint, error) {
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
