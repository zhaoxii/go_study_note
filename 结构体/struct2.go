package main

import (
	"fmt"
)

type Student struct {
	Name  string
	age   int
	score int
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (s Student) test() {
	fmt.Println("test......")
}

func (s *Student) SetScore(score int) {
	s.score = score
}

func main() {
	fmt.Println("hello world")
	s1 := Pupil{}
	s1.Student.Name = "zx"
	s1.Student.age = 111
	s1.Name = "zxzxzx"
	s1.test()
	s1.Student.SetScore(100)
	s2 := Graduate{}
	s2.Student.Name = "zxzx"
	s2.Student.age = 12
	s2.test()
	s2.SetScore(200)
	fmt.Println(s1, s2)

	s3 := Pupil{Student{Name: "xzzx"}}
	fmt.Println(s3)

}
