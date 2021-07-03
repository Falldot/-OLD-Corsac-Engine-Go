package gl

type Config struct {
	TargetFPS        uint16
	IdleThreshold    float64
	CurrentTimeFunc  func() float64
	ProcessInputFunc func() (quit bool)
	UpdateFunc       func(dt float64)
	RenderFunc       func()
}

func Create(c Config) func() {
	secsPerUpdate := 1 / float64(c.TargetFPS)
	var current, elapsed float64

	return func() {
		previous := c.CurrentTimeFunc()
		lag := 0.0

		for quit := false; !quit; {
			current = c.CurrentTimeFunc()
			elapsed = current - previous
			previous = current

			if elapsed > c.IdleThreshold {
				continue
			}

			lag += elapsed

			quit = c.ProcessInputFunc()

			for lag >= secsPerUpdate {
				c.UpdateFunc(secsPerUpdate)
				lag -= secsPerUpdate
			}

			c.RenderFunc()
		}
	}
}
