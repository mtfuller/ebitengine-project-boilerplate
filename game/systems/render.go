package systems

import (
	//"fmt"
	"image/color"
	"github.com/yourname/yourgame/engine/ecs"
	"github.com/yourname/yourgame/game/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Render struct {
	ecs.BaseSystem
	viewport *ebiten.Image
	layers map[int][]*ecs.Entity
}

func NewRender() *Render {
	render := Render{
		viewport: ebiten.NewImage(640, 480),
		layers: make(map[int][]*ecs.Entity),
	}

	return &render
}

func (b *Render) AddEntity(e *ecs.Entity)  {
	if b.Entities == nil {
		b.Entities = make(map[*ecs.Entity]struct{})
	}

	b.Entities[e] = struct{}{}

	component := e.GetComponent("render")
	render, _ := component.(*components.Render)

	b.layers[render.Z] = append(b.layers[render.Z], e)
}

func (r Render) Render(screen *ebiten.Image)  {
	r.viewport.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	
	op := &ebiten.DrawImageOptions{}

	for i := 0; i < 16; i++ {
		entities := r.layers[i]

		for _, ptr := range entities {
			e := *ptr;

			if e.HasAllComponents([]string{"position","render"}) {
				component := e.GetComponent("position")
				position, ok := component.(*components.Position)
				if !ok {
					continue
				}
	
				component = e.GetComponent("render")
				render, ok := component.(*components.Render)
				if !ok {
					continue
				}

				sprite := render.Spritesheet.GetSprite(render.CurrentSprite)

				if render.Countdown > 0 {
					render.Countdown -= 1
				} else {
					nextFrame := (render.CurrentFrame + 1)

					if nextFrame >= len(sprite) {
						nextFrame = len(sprite) % nextFrame
					}

					render.CurrentFrame = nextFrame

					render.Countdown = render.MaxCountdown
				}

				//fmt.Println("render.CurrentFrame=", render.CurrentFrame)
	
				op.GeoM.Reset()
				op.GeoM.Translate(position.X, position.Y)
				r.viewport.DrawImage(sprite[render.CurrentFrame], op)
			}
		}
	}

	screen.DrawImage(r.viewport, nil)
}