// package demo

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func polar1(ch2 chan Polar) chan float64 {
// 	ch1 := make(chan float64, 2)
// 	p := Polar{}
// 	i := 1
// 	go func() {
// 		for {
// 			v, ok := <-ch1
// 			if ok {
// 				if i == 1 {
// 					p.T = v
// 				} else if i == 2 {
// 					p.R = v
// 					ch2 <- p
// 				}
// 				i++
// 			}
// 		}
// 	}()
// 	return ch1
// }

// func cartensian() chan Polar {
// 	ch2 := make(chan Polar)
// 	p := Point{}
// 	go func() {
// 		for {
// 			v, ok := <-ch2
// 			if ok {
// 				x := v.R * math.Cos(v.T)
// 				y := v.R * math.Sin(v.T)
// 				p = Point{X: x, Y: y}
// 				fmt.Printf("value is: %v", p)
// 			}
// 		}
// 	}()
// 	return ch2
// }

// func format2(ch chan float64) {
// 	reader := bufio.NewReader(os.Stdin)
// 	// i := 0
// 	for {
// 		buf, err := reader.ReadBytes('\n')
// 		if *numberline {
// 			str, err := strconv.Atoi(strings.ReplaceAll(string(buf), "\n", ""))
// 			if err == nil {
// 				fmt.Println("输入：", str)
// 				ch <- float64(str)
// 				// i++
// 				// if i == 2 {
// 				// 	break
// 				// }
// 			}
// 		} else {
// 			fmt.Fprintf(os.Stdout, "%s", buf)
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// func InitPolarToCartesian() {
// 	flag.PrintDefaults()
// 	flag.Parse()
// 	ch2 := cartensian()
// 	defer close(ch2)
// 	ch1 := polar1(ch2)
// 	defer close(ch1)
// 	format2(ch1)
// }

package demo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	prompt = fmt.Sprintf(prompt, "Ctrl+D")
}

func InitPolarToCartesian() {
	questions := make(chan Polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan Polar) chan Point {
	answers := make(chan Point)
	go func() {
		for {
			polar := <-questions
			polar.T = polar.T * math.Pi / 180.0
			x := polar.R * math.Cos(polar.T)
			y := polar.R * math.Sin(polar.T)
			answers <- Point{X: x, Y: y}
		}
	}()
	return answers
}

func interact(questions chan Polar, answers chan Point) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		if numbers := strings.Fields(line); len(numbers) == 2 {
			fmt.Println(numbers)
			floats, err := floatsForStrings(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			questions <- Polar{R: floats[0], T: floats[1]}
			coord := <-answers
			fmt.Printf(result, floats[0], floats[1], coord.X, coord.Y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
	fmt.Println()
}

func floatsForStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
