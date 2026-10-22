//TODO: Test on Termux
//FIXME: Single char lengths
package main

import (
	"fmt"
	//_ "breedable"
	"strings"
	"math/rand"
	"time"
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

func randomize(pool *[10]StringBreedable, pass string) {
	var temp string
	var err error
	for i := 0; i < 10; i++ {
		for j := 0; j <= len(pass); j++ {
			temp = strings.Join([]string{string(rune(rand.Int31() % 127))}, "")
		}
		var sb *StringBreedable
		sb, err = NewStringBreedable(temp)
		if err != nil {
			panic(err)
		}
		pool[i] = *sb
		temp = ""
	}
}

func breed(prePool [10]StringBreedable, pass string, thresh int) [10]StringBreedable {
	var pool [10]StringBreedable
	for _, str := range prePool {
		for _, char := range []rune(pass) {
			if strings.Contains(str.value, string(char)) {
				str.score()
				break
			}
		}
	}
	for i, str := range prePool {
		if str.points == thresh {
			pool[i] = str
		}
	}

	for i, str := range pool {
		pool[i] = str.breed(pool[rand.Int() % 10])
	}

	return pool
}

func main() {
	var thresh int
	rand.Seed(time.Now().UnixNano())
	var pool [10]StringBreedable
	var secretPasscode string
	fmt.Print("Enter The Secret Code: ")
	fmt.Scan(&secretPasscode)
	length := len(secretPasscode)
	randomize(&pool, secretPasscode)
	var i int
	for thresh != length {
		if i > 4294967296 {break}
		for thresh = 0; thresh <= length; thresh++ {
			if pool == [10]StringBreedable{} {
				randomize(&pool, secretPasscode)
				break
			}
			pool = breed(pool, secretPasscode, thresh)
			fmt.Println(pool)
			for i,_ = range pool {
				pool[i].reset()
			}
		}
		i++
	}
}