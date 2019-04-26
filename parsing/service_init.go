package parsing

func NewPageService(store PageGetter) *service {
	return &service{store}
}
