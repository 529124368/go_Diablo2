package joystick

import (
	"bytes"
	"embed"
	"game/status"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var JoyStick_o, JoyStick_base *ebiten.Image = nil, nil

// Sprite represents an image.
type Sprite struct {
	image                *ebiten.Image
	x                    int
	y                    int
	dir                  float64
	JoyStickx, JoySticky int
	R                    float64
}

func New(m *ebiten.Image, x, y int, r float64) *Sprite {
	s := &Sprite{
		image:     m,
		x:         x,
		y:         y,
		JoyStickx: x,
		JoySticky: y,
		R:         r,
	}
	return s
}
func (s *Sprite) In(x, y int) bool {
	return s.image.At(x-s.x, y-s.y).(color.RGBA).A > 0
}

// MoveBy moves the sprite by (x, y).
func (s *Sprite) MoveBy(x, y int) (int, int) {
	ddx := x
	ddy := y
	dx := float64(s.x + x - s.JoyStickx)
	dy := float64(s.y + y - s.JoySticky)
	rad := math.Atan2(dy, dx)
	s.dir = rad * 180 / math.Pi
	max := s.R * math.Cos(rad)
	may := s.R * math.Sin(rad)
	if math.Abs(dx) > math.Abs(max) {
		ddx = int(max) + s.JoyStickx - s.x
	}
	if math.Abs(dy) > math.Abs(may) {
		ddy = int(may) + s.JoySticky - s.y
	}
	return ddx, ddy
}

func (s *Sprite) GetDir() float64 {
	if s.dir >= -270 && s.dir < -90 {
		s.dir = 270 + 180 + s.dir
	} else {
		s.dir = s.dir + 90
	}
	return s.dir
}

func (s *Sprite) Back() {
	s.x = s.JoyStickx
	s.y = s.JoySticky
}

// Draw draws the sprite.
func (s *Sprite) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	xx, yy := s.MoveBy(x, y)
	op.GeoM.Translate(float64(s.x+xx), float64(s.y+yy))
	screen.DrawImage(s.image, op)
}

// StrokeSource represents a input device to provide strokes.
type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

// MouseStrokeSource is a StrokeSource implementation of mouse.
type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

// TouchStrokeSource is a StrokeSource implementation of touch.
type TouchStrokeSource struct {
	ID ebiten.TouchID
}

func (t *TouchStrokeSource) Position() (int, int) {
	return ebiten.TouchPosition(t.ID)
}

func (t *TouchStrokeSource) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}

// Stroke manages the current drag state by mouse.
type Stroke struct {
	source StrokeSource

	// initX and initY represents the position when dragging starts.
	initX int
	initY int

	// currentX and currentY represents the current position
	currentX int
	currentY int

	released bool

	// draggingObject represents a object (sprite in this case)
	// that is being dragged.
	draggingObject interface{}
}

func NewStroke(source StrokeSource) *Stroke {
	cx, cy := source.Position()
	return &Stroke{
		source:   source,
		initX:    cx,
		initY:    cy,
		currentX: cx,
		currentY: cy,
	}
}

func (s *Stroke) Update() {
	if s.released {
		return
	}
	if s.source.IsJustReleased() {
		s.released = true
		return
	}
	x, y := s.source.Position()
	s.currentX = x
	s.currentY = y
}

func (s *Stroke) IsReleased() bool {
	return s.released
}

func (s *Stroke) Position() (int, int) {
	return s.currentX, s.currentY
}

func (s *Stroke) PositionDiff() (int, int) {
	dx := s.currentX - s.initX
	dy := s.currentY - s.initY
	return dx, dy
}

func (s *Stroke) DraggingObject() interface{} {
	return s.draggingObject
}

func (s *Stroke) SetDraggingObject(object interface{}) {
	s.draggingObject = object
}

//main
type JoyStickBase struct {
	touchIDs             []ebiten.TouchID
	strokes              map[*Stroke]struct{}
	JoyStickM, JoyStickB *Sprite
	Dir                  float64
	jmx, jbx, jmy, jby   int
	R                    float64
}

func NewJoyStick(asset *embed.FS, x1, y1, x2, y2 int, r float64) *JoyStickBase {
	// Initialize the sprites.
	if JoyStick_o == nil {
		ss, _ := asset.ReadFile("resource/UI/stick_o.png")
		img, _, err := image.Decode(bytes.NewReader(ss))
		if err != nil {
			log.Fatal(err)
		}

		JoyStick_o = ebiten.NewImageFromImage(img)
		ss, _ = asset.ReadFile("resource/UI/stick_base.png")
		img, _, err = image.Decode(bytes.NewReader(ss))
		if err != nil {
			log.Fatal(err)
		}
		JoyStick_base = ebiten.NewImageFromImage(img)
	}

	// Initialize the game.
	return &JoyStickBase{
		strokes:   map[*Stroke]struct{}{},
		JoyStickM: New(JoyStick_o, x2, y2, r),
		JoyStickB: New(JoyStick_base, x1, y1, 0),
		Dir:       -1,
		jmx:       x1,
		jmy:       y1,
		jbx:       x2,
		jby:       y2,
		R:         r,
	}
}
func (j *JoyStickBase) spriteAt(x, y int) *Sprite {
	if j.JoyStickM.In(x, y) {
		return j.JoyStickM
	}
	return nil
}

func (j *JoyStickBase) updateStroke(stroke *Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(*Sprite)
	if s == nil {
		return
	}
	stroke.SetDraggingObject(nil)
}

func (j *JoyStickBase) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && j.JoyStickM.In(ebiten.CursorPosition()) {
		s := NewStroke(&MouseStrokeSource{})
		s.SetDraggingObject(j.spriteAt(s.Position()))
		j.strokes[s] = struct{}{}
	}
	//触摸屏幕
	j.touchIDs = inpututil.AppendJustPressedTouchIDs(j.touchIDs[:0])
	for _, id := range j.touchIDs {
		s := NewStroke(&TouchStrokeSource{ID: id})
		s.SetDraggingObject(j.spriteAt(s.Position()))
		j.strokes[s] = struct{}{}
	}

	if len(j.strokes) > 0 {
		status.Config.IsTakeJoyStick = true
	} else {
		status.Config.IsTakeJoyStick = false
	}
	for s := range j.strokes {
		j.updateStroke(s)
		if s.IsReleased() {
			j.JoyStickM.Back()
			j.Dir = -1
			delete(j.strokes, s)
		}
	}
	return nil
}

func (j *JoyStickBase) Draw(screen *ebiten.Image) {
	draggingSprites := map[*Sprite]struct{}{}
	for s := range j.strokes {
		if sprite := s.DraggingObject().(*Sprite); sprite != nil {
			draggingSprites[sprite] = struct{}{}
		}
	}
	//摇杆背景
	j.JoyStickB.Draw(screen, 0, 0)
	//
	if _, ok := draggingSprites[j.JoyStickM]; !ok {
		j.JoyStickM.Draw(screen, 0, 0)
	}
	for s := range j.strokes {
		dx, dy := s.PositionDiff()
		if sprite := s.DraggingObject().(*Sprite); sprite != nil {
			j.Dir = sprite.GetDir()
			sprite.Draw(screen, dx, dy)
		}
	}
}
