package data

import "fmt"

func (s *Student) ToString() string {
	return fmt.Sprint(s.Base) + s.Name + fmt.Sprint(s.Age) + s.School
}
