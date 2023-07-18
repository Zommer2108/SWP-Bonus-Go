package main // package system

import "fmt" // format package
import "strconv"

// ///////////////////////////////////////////////////////////////////////
// Qualified import.
// Functions (types) starting with upper-case are exported.
func example1() {
	fmt.Printf("Hello %d \n", 5)
}

// ///////////////////////////////////////////////////////////////////////
// Keyword for variable and function declarations.
// First comes the name and then the type of the variable.
// No ";" required, one statement per line.
func example2() {
	var x int // In C written as "int x;"
	x = 5
	fmt.Printf("Hello %d \n", x)
}

// ///////////////////////////////////////////////////////////////////////
// Local type inference.
func example3() {
	x := 5 // In C++ written as "auto x = 5;"
	fmt.Printf("Hello %d \n", x)

}

// ///////////////////////////////////////////////////////////////////////
// Arrays, slices and loops.
func example4() {
	var x [3]int

	x[0] = 0
	x[1] = 1
	x[2] = 2

	y := [3]int{2, 1, 0}

	// One looping construct "for".
	for i := 0; i < 3; i++ {
		fmt.Printf("%d \n", y[i])
	}

	// range yields the current index position and element.
	for i, e := range x {
		fmt.Printf("%d %d\n", e, y[i])
	}

	// Slice, dynamically-sized, internally represented via fixed-size arrays.
	z := []int{0, 1, 2}
	z = append(z, 5)

	// We don't care about the index position, indicated by "_".
	for _, e := range z {
		fmt.Printf("%d \n", e)
	}
}

/////////////////////////////////////////////////////////////////////////
// Functions.
// Return type (if any) comes last. Multiple return types possible.

func add(x int, y int) int { // In C written as "int addExample(int x, int y) { ... }
	return x + y
}

func div(x int, y int) (int, bool) {
	if y == 0 {
		return 0, false
	}

	return x / y, true
}

func example5() {
	r1, status1 := div(add(3, 5), 3)
	if status1 {
		fmt.Printf("%d \n", r1)
	} else {
		fmt.Printf("failure\n")
	}

	r2, status2 := div(add(3, 5), 0)
	if status2 {
		fmt.Printf("%d \n", r2)
	} else {
		fmt.Printf("failure\n")
	}

}

/////////////////////////////////////////////////////////////////////////
// Higher-order functions.

// Check if any element in xs satisfies some predicate p.
func any(p func(int) bool, xs []int) (bool, int) {
	for _, x := range xs {
		if p(x) {
			return true, x
		}
	}
	return false, -1
}

func gt5(x int) bool {
	return x > 5
}

func example6() {
	xs := []int{3, 5, 6, 1}
	b, r := any(gt5, xs)
	fmt.Printf("%b %d \n", b, r)

}

/////////////////////////////////////////////////////////////////////////
// Lambda functions (anonymous functions).

func sum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}

}

func example7() {
	xs := []int{3, 5, 6, 1}
	p := func(x int) bool {
		return x > 5
	}
	b, r := any(p, xs)
	fmt.Printf("%b %d \n", b, r)

	inc := sum(1)
	fmt.Printf("%d \n", inc(3))
}

/////////////////////////////////////////////////////////////////////////
// Structs.

type rectangle struct {
	x int
	y int
}

func scaleRectangle(p rectangle, s int) rectangle {
	return rectangle{p.x * s, p.y * s}
}

func example8() {
	p := rectangle{1, 2}
	q := scaleRectangle(p, 3)
	fmt.Printf(" (%d, %d) \n", q.x, q.y)
}

/////////////////////////////////////////////////////////////////////////
// Pointers.
// No pointer arithmetic.
// Automatic garbage collection. Stack allocated objects can be moved to the heap.

func stackToHeap() *int {
	x := 1
	return &x
}

func scaleRectangleRef(p *rectangle, s int) {
	p.x = p.x * s
	(*p).y = (*p).y * s
}

func example9() {
	int_pointer := stackToHeap()
	fmt.Printf(" %d \n", *int_pointer)

	p := rectangle{1, 2}
	fmt.Printf(" (%d, %d) \n", p.x, p.y)
	scaleRectangleRef(&p, 3)
	fmt.Printf(" (%d, %d) \n", p.x, p.y)

}

/////////////////////////////////////////////////////////////////////////
// Methods and interfaces.
//
// Methods operate on "objects" (whose type is defined via "type").
// Interfaces specify a collection of methods that operate on the same object.
// In Go speak, we refer to such an object as the receiver.

type square struct {
	x int
}

// Overloaded methods.
func (r rectangle) area() int {
	return r.x * r.y
}

/*
   // Each method declaration is uniquely identified by the receiver type and method name.
   // If we uncomment the below, we get an error.
   func (r rectangle) area(z int) int {
       return r.x * r.y * z
   }
*/

func (s square) area() int {
	return s.x * s.x
}

func example10() {

	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	x := r.area()
	y := s.area()

	fmt.Printf("%d %d \n", x, y)

}

// Common interface.
type shape interface {
	area() int
}

func sumArea(x, y shape) int {
	return x.area() + y.area()
}

func example11() {

	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	x2 := sumArea(r, s)  // applied on (rectangle, square)
	x2b := sumArea(r, r) // applied on (rectangle, rectangle)

	// Go uses structural subtyping.
	// rectangle <= shape because
	//    (1) rectangle implements the area method
	//    (2) interface shape demands that shape method must be present.

	fmt.Printf("%d %d \n", x2, x2b)

}

// The "any" interface.

func swap(x *interface{}, y *interface{}) {
	var tmp interface{}

	tmp = *x
	*x = *y
	*y = tmp

}

func example12() {
	var x, y int
	x = 3
	y = 2

	var x2, y2 interface{}
	x2 = x // int <= interface{}
	y2 = y

	swap(&x2, &y2)

	x = x2.(int) // Type assertion
	y = y2.(int)

	fmt.Printf("%d %d", x, y)

}

/////////////////////////////////////////////////////////////////////////
// Simple expression language

type Exp interface {
	pretty() string
	eval() Val
}

// Values

type Kind int

const (
	ValueInt  Kind = 0
	ValueBool Kind = 1
	Undefined Kind = 2
)

type Val struct {
	flag Kind
	valI int
	valB bool
}

func mkInt(x int) Val {
	return Val{flag: ValueInt, valI: x}
}
func mkBool(x bool) Val {
	return Val{flag: ValueBool, valB: x}
}
func mkUndefined() Val {
	return Val{flag: Undefined}
}

func showVal(v Val) string {
	var s string
	switch {
	case v.flag == ValueInt:
		s = Num(v.valI).pretty()
	case v.flag == ValueBool:
		s = Bool(v.valB).pretty()
	case v.flag == Undefined:
		s = "Undefined"
	}
	return s
}

// Cases

type Bool bool
type Num int
type Mult [2]Exp
type Plus [2]Exp
type And [2]Exp
type Or [2]Exp

// pretty print

func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}

}

func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}

func (e Mult) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "*"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Plus) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "+"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e And) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "&&"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Or) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "||"
	x += e[1].pretty()
	x += ")"

	return x
}

// Evaluator

func (x Bool) eval() Val {
	return mkBool((bool)(x))
}

func (x Num) eval() Val {
	return mkInt((int)(x))
}

func (e Mult) eval() Val {
	n1 := e[0].eval()
	n2 := e[1].eval()
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}

func (e Plus) eval() Val {
	n1 := e[0].eval()
	n2 := e[1].eval()
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}

func (e And) eval() Val {
	b1 := e[0].eval()
	b2 := e[1].eval()
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval() Val {
	b1 := e[0].eval()
	b2 := e[1].eval()
	switch {
	case b1.flag == ValueBool && b1.valB == true:
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}

// Helper functions to build ASTs by hand

func number(x int) Exp {
	return Num(x)
}

func boolean(x bool) Exp {
	return Bool(x)
}

func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})

	// The type Plus is defined as the two element array consisting of Exp elements.
	// Plus and [2]Exp are isomorphic but different types.
	// We first build the AST value [2]Exp{x,y}.
	// Then cast this value (of type [2]Exp) into a value of type Plus.

}

func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}

func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}

func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}

func example13() {

	run := func(e Exp) {
		fmt.Printf("\n ******* ")
		fmt.Printf("\n %s", e.pretty())
		fmt.Printf("\n %s", showVal(e.eval()))
	}

	{
		ast := plus(mult(number(1), number(2)), number(0))

		run(ast)
	}

	{
		ast := and(boolean(true), number(0))
		run(ast)
	}

	{
		ast := or(boolean(false), number(0))
		run(ast)
	}

}

func main() {

	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
	example8()
	example9()
	example10()
	example11()
	example12()
	example13()
}
