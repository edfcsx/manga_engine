package manga

import "time"

type fpsCounter struct {
	lastTime time.Time
	frames   int
	fps      float64
}

func makeFpsCounter() *fpsCounter {
	return &fpsCounter{
		lastTime: time.Now(),
		frames:   0,
		fps:      0,
	}
}

func (f *fpsCounter) Update() {
	now := time.Now()
	elapsed := now.Sub(f.lastTime).Seconds()
	f.frames++

	if elapsed >= 1.0 {
		f.fps = float64(f.frames) / elapsed
		f.frames = 0
		f.lastTime = now
	}
}

func (f *fpsCounter) GetFPS() float64 {
	return f.fps
}
