package controller

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

func Touch(minx, miny, maxx, maxy int) bool {
	ids := ebiten.TouchIDs()
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	for _, id := range ids {
		idx, idy := ebiten.TouchPosition(id)
		if idx >= minx && idx <= maxx && idy >= miny && idy <= maxy {
			return true
		}
	}
	return false
}

func IsTouch() bool {
	return len(ebiten.TouchIDs()) > 0
}
