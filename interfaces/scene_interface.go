package mangaI

type Scene interface {
	Initialize()
	Update()
	Render()
}
