package func_option_mode

type MoreParameters struct {
	POne   int
	PTwo   int
	PThree float32
}
type InitOption func(parameters *MoreParameters)

// WithPOne 用于创建一个函数对象，该对象是用于初始化该对象
func WithPOne(v int) InitOption {
	return func(item *MoreParameters) {
		item.POne = v
	}
}

// WitchPTwo 用于创建一个函数对象，该对象是用于初始化该对象
func WitchPTwo(v int) InitOption {
	return func(item *MoreParameters) {
		item.PTwo = v
	}
}

// NewMoreParameters 入参是 初始化对象属性的函数列表.
func NewMoreParameters(flist ...InitOption) *MoreParameters {
	item := &MoreParameters{}
	for _, op := range flist {
		op(item)
	}
	return item
}

func GetOptMode() *MoreParameters {
	return NewMoreParameters(WithPOne(1), WitchPTwo(200))
}
