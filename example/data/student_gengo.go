package data

import "fmt"

func (s *Student) ToString() string {
	return s.Name + fmt.Sprint(s.Age) + s.School
}
