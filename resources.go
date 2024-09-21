package main

import "fmt"

type resources struct {
	gold       int
	materials  int
	experience int
	level      int
}

func viewResources(resources resources) string {
	resourceBar := fmt.Sprintf("Gold: %s | Materials: %s | Level: %s",
		resources.gold,
		resources.materials,
		resources.level)
	return resourceBar
}
