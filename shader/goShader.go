package shader

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var Shader *ebiten.Shader

func init() {
	Shader, _ = ebiten.NewShader([]byte(`
	package main
	var IsOutLine float
	func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
		pos := texCoord
		offset :=1/imageSrcTextureSize()
		if IsOutLine == 1 && imageSrc0UnsafeAt(vec2(pos.x,pos.y)).a==0 && (imageSrc0UnsafeAt(vec2(pos.x+offset.x,pos.y)).a!=0 || imageSrc0UnsafeAt(vec2(pos.x-offset.x,pos.y)).a!=0 || imageSrc0UnsafeAt(vec2(pos.x,pos.y+offset.y)).a!=0 || imageSrc0UnsafeAt(vec2(pos.x+offset.x,pos.y-offset.y)).a!=0){
			return vec4(1,0,0,1)
		}
		return imageSrc0UnsafeAt(texCoord)
	}
`))
}
