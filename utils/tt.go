package utils

type Work struct {
	Who, DoWhat string
}

type DemoM string

func (m *DemoM) DoWork(w *Work, whoT *string) error {
	*whoT = "是谁：" + w.Who + "，在做什么---" + w.DoWhat
	return nil
}

type DemoM1 string

func (m1 *DemoM1) DoWork1(w *Work, whoT *string) error {
	*whoT = "是谁1：" + w.Who + "，在做什么1---" + w.DoWhat
	return nil
}
