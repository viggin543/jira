package collections

type Pie struct {
	Id string
	Number int
}

//go:generate pie Pies.*
type Pies []Pie

