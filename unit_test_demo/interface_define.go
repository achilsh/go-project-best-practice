package unit_test_demo

type UserEr interface {
	SetName(string)
	GetName(vType int) string
}
type UserWrapper struct {
	OPEr UserEr
}

func (g *UserWrapper) DoCall(v1 string) string {
	g.OPEr.SetName(v1)
	v := g.OPEr.GetName(123)
	return v
}
