package crsc

import (
	cs "CRSC/internal/CS"
	ecs "CRSC/internal/ECS"
	ee "CRSC/internal/EE"
	ge "CRSC/internal/GE"
	gl "CRSC/internal/GL"

	"github.com/veandco/go-sdl2/sdl"
)

func Init() {
	ge.Init("My Game", 800, 600)

	player := ecs.CreateEntity()
	ecs.AddComponent(player, cs.CreateTransform(100, 100, 32, 32), cs.COMPONENT_TRANSFORM)
	ecs.AddComponent(player, &cs.Sprite{}, cs.COMPONENT_SPRITE)
	ecs.AddComponent(player, &cs.KeyboardControlle{}, cs.COMPONENT_KEYBOARDCONTROLLER)

	//mapa := ecs.CreateEntity()
	//ecs.AddComponent(mapa, nil, cs.COMPONENT_MAP)

	ecs.Init()
	ee.Init()

	ecs.GetComponent(player, cs.COMPONENT_SPRITE).(*cs.Sprite).
		AddAnimation("Idle", 0, 6, 300).
		AddAnimation("Right", 1, 6, 100).
		AddAnimation("Down", 2, 6, 100).
		AddAnimation("Up", 3, 6, 100).
		AddAnimation("DownRight", 4, 6, 100).
		SetTex("./data/fox.png")

	player.Active(true)

	config := gl.Config{
		TargetFPS:     60,
		IdleThreshold: 1,
		CurrentTimeFunc: func() float64 {
			return float64(sdl.GetTicks()) * 0.001
		},
		ProcessInputFunc: func() (quit bool) {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				if event.GetType() == sdl.QUIT {
					return true
				}
			}
			return false
		},
		UpdateFunc: func(dt float64) {
			ecs.Update()
		},
		RenderFunc: func() {
			ge.Render.SetDrawColor(0, 0, 0, 255)
			ge.Render.Clear()
			ecs.Render()
			ge.Render.Present()
		},
	}

	runLoop := gl.Create(config)

	runLoop()
	ge.Quit()
}
