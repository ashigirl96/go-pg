package main

import (
	"fmt"
	"github.com/ashigirl96/go-pg/pkg"
	"math"
	"time"
)

type MyFloat64 float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (f MyFloat64) Abs() float64 {
	return float64(-f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 直で書き換える場合は、ポイントレシーバーを使う
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat64) Minus() MyFloat64 {
	return -f
}

func methods() {
	pkg.FuncName()
	v := Vertex{1, 1}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
	w := &Vertex{2, 2}
	w.Scale(10)
	fmt.Println(w.Abs())
}

func methodsContinued() {
	pkg.FuncName()
	f := MyFloat64(10)
	fmt.Println(f.Minus())
}

func interfaces() {
	pkg.FuncName()
	var a Abser
	f := MyFloat64(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	var b Abser = &Vertex{1, 1}
	fmt.Println(b.Abs())
}

func emptyInterface() {
	pkg.FuncName()
	describe := func(i interface{}) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
}

func typeSwitches() {
	pkg.FuncName()
	do := func(v interface{}) {
		switch i := v.(type) {
		case int:
			fmt.Println(i * 2)
		case string:
			fmt.Println("hello,", i)
		default:
			fmt.Println(i)
		}
	}
	do(nil)
	do(1)
	do("world")
}

func stringers() {
	pkg.FuncName()
	p1 := Person{"Yamada", 10}
	p2 := Person{"Tanaka", 20}
	fmt.Println(p1)
	fmt.Println(p2)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func exerciseStringers() {
	pkg.FuncName()
	hosts := map[string]IPAddr{
		"lookback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func errors() {
	pkg.FuncName()
	run := func() error {
		return &MyError{
			time.Now(),
			"it didn't work",
		}
	}
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (f ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v cannot use", float64(f))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
}

func exerciseErrors() {
	pkg.FuncName()
	if _, err := Sqrt(2); err != nil {
		fmt.Println(err)
	}
	if _, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))
}

func main() {
	methods()
	methodsContinued()
	interfaces()
	emptyInterface()
	typeSwitches()
	stringers()
	exerciseStringers()
	errors()
	exerciseErrors()
}
