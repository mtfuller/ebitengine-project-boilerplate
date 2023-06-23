package systems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

type Render struct {
	ecs.BaseSystem
	viewport *ebiten.Image
	layers   map[int][]*ecs.Entity
}

func NewRender() *Render {
	render := Render{
		viewport: ebiten.NewImage(640, 480),
		layers:   make(map[int][]*ecs.Entity),
	}

	return &render
}

func (r Render) GetName() string {
	return "System::Render"
}

func (b *Render) HandleEntityCreated(e *ecs.Entity) {
	component := e.GetComponent("render")
	render, _ := component.(*components.Render)

	b.layers[render.Z] = append(b.layers[render.Z], e)
}

func (r Render) Render(screen *ebiten.Image) {
	r.viewport.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	for i := 0; i < 16; i++ {
		entities := r.layers[i]

		for _, ptr := range entities {
			e := *ptr

			if e.HasAllComponents([]string{"position", "render"}) {
				component := e.GetComponent("position")
				position, _ := component.(*components.Position)

				component = e.GetComponent("render")
				render, _ := component.(*components.Render)

				sprite := render.Spritesheet.Sprites[render.EntityName][render.SpriteName]

				if render.FrameCount < sprite.Durations[render.CurrentFrame] {
					render.FrameCount++
				} else {
					nextFrame := (render.CurrentFrame + 1)

					if nextFrame >= len(sprite.Frames) {
						nextFrame = len(sprite.Frames) % nextFrame
					}

					render.CurrentFrame = nextFrame
					render.FrameCount = 0
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Reset()
				op.GeoM.Translate(position.X, position.Y)

				r.viewport.DrawImage(sprite.Frames[render.CurrentFrame], op)
			}
		}
	}

	screen.DrawImage(r.viewport, nil)
}
