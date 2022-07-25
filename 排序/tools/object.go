package tools

import "fmt"

type Object struct {
	Score int
	Age   int
	Name  string
}

func NewObject(score, age int, name string) *Object {
	return &Object{
		score, age, name,
	}
}
func (i Object) Println() string {
	return fmt.Sprintf("Object [score=%d, name=%s]", i.Score, i.Name)
}
func (i Object) CompareTo(o *Object) int {
	return i.Age - o.Age
}
