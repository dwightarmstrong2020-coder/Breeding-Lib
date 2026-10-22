package breedable

import (
	"strings"
)

type Breedable interface {
	breed(mate Breedable) (child Breedable)
}

type StringBreedable struct {
	value string
	points int
}

func (str *StringBreedable) score() {
	str.points++
}

func (str *StringBreedable) reset() {
	str.points = 0
}

func NewStringBreedable(val string) (*StringBreedable, error) {
	if val == "" {
		return nil, fmt.Errorf("Can't be empty")
	}
	return &StringBreedable{val, 0}, nil
}

func (str StringBreedable) breed(mate StringBreedable) (child StringBreedable) {
	for i, v := range []rune(str.value) {
		if i > ((len([]rune(str.value)) - 1) / 2) {
			break
		}
		child.value = strings.Join([]string{child.value, string(v)}, "")
	}
	mateRunes := []rune(mate.value)
	if len(mateRunes) > 1 {
		mid := (len(mateRunes) - 1) / 2
		for _, v := range mateRunes[mid : len(mateRunes)-1] {
			child.value = strings.Join([]string{child.value, string(v)}, "")
		}
	}
	return
}