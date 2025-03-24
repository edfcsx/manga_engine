package mangaI

type Scene interface {
	Initialize()
	Update(float64)
	Render()
}
