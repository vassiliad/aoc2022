package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Jet int8
type Rock int8

const (
	JetLeft  Jet = 0
	JetRight Jet = 1

	RockA Rock = 0
	RockB Rock = 1
	RockC Rock = 2
	RockD Rock = 3
	RockE Rock = 4

	RockNumber int = 5
)

func (r *Rock) Width() int {
	if *r == RockA {
		return 4
	} else if *r == RockB {
		return 3
	} else if *r == RockC {
		return 3
	} else if *r == RockD {
		return 1
	} else if *r == RockE {
		return 2
	} else {
		panic(*r)
	}
}

func (r *Rock) Height() int {
	if *r == RockA {
		return 1
	} else if *r == RockB {
		return 3
	} else if *r == RockC {
		return 3
	} else if *r == RockD {
		return 4
	} else if *r == RockE {
		return 2
	} else {
		panic(*r)
	}
}

func (r *Rock) GetColliderIndex(index, width int) []int {
	x, y := index%width, index/width

	return r.GetCollider(x, y, width)
}

func (r *Rock) GetCollider(x, y, width int) []int {
	ret := []int{}

	if *r == RockA {
		w := r.Width()
		for dx := 0; dx < w; dx++ {
			ret = append(ret, x+dx+y*width)
		}
	} else if *r == RockB {
		w := r.Width()
		for i := 0; i < w; i++ {
			if i != 1 {
				ret = append(ret, x+i+(y+1)*width, x+1+(y+i)*width)
			} else {
				ret = append(ret, x+i+(i+y)*width)
			}
		}
	} else if *r == RockC {
		w := r.Width()
		for i := 0; i < w; i++ {
			if i != 2 {
				ret = append(ret, x+w-1+(y+i)*width, x+i+y*width)
			} else {
				ret = append(ret, x+i+(y+i)*width)
			}
		}
	} else if *r == RockD {
		h := r.Height()
		for dy := 0; dy < h; dy++ {
			ret = append(ret, x+(y+dy)*width)
		}
	} else if *r == RockE {
		h := r.Height()
		for dy := 0; dy < h; dy++ {
			for dx := 0; dx < h; dx++ {
				ret = append(ret, x+dx+(y+dy)*width)
			}
		}
	} else {
		panic(*r)
	}

	return ret
}

func ReadScanner(scanner *bufio.Scanner) ([]Jet, error) {
	ret := []Jet{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		for _, c := range line {
			if c == '>' {
				ret = append(ret, JetRight)
			} else if c == '<' {
				ret = append(ret, JetLeft)
			} else {
				return ret, fmt.Errorf("unknown character %c", c)
			}
		}
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Jet, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Jet, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
