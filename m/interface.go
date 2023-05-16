package m

type CouldBeVisible interface {
	ToVo() Visible
}

type Visible interface {
	Visible()
}
