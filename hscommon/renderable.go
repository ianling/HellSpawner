package hscommon

type Renderable interface {
	Render()
	Cleanup()
}
