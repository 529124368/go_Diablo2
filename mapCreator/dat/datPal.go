package dat

import (
	"fmt"
	"game/interfaces"
)

const (
	numColors = 256
)

// DATPalette represents a 256 color palette.
type DATPalette struct {
	colors [numColors]interfaces.Color
}

// New creates a new dat palette
func NewP() *DATPalette {
	result := &DATPalette{}
	for i := range result.colors {
		result.colors[i] = &DATColor{}
	}

	return result
}

// NumColors returns the number of colors in the palette
func (p *DATPalette) NumColors() int {
	return len(p.colors)
}

// GetColors returns the slice of colors in the palette
func (p *DATPalette) GetColors() [numColors]interfaces.Color {
	return p.colors
}

// GetColor returns a color by index
func (p *DATPalette) GetColor(idx int) (interfaces.Color, error) {
	if color := p.colors[idx]; color != nil {
		return color, nil
	}

	return nil, fmt.Errorf("cannot find color index '%d in palette'", idx)
}
