package controller

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func MouseOnceLeftPress() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func MouseOnceRightPress() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
}

func MouseRightPress() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

func MouseleftPress() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func MousePressF1() bool {
	return ebiten.IsKeyPressed(ebiten.KeyF1)
}
