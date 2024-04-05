package main

import "fmt"

func Datatype() {
    // Numeric Types:
	name := "Sanaklp Haritash"
	fmt.Printf(name)

    // 1. int: Signed integer type, can be 32 or 64 bits.
    var x int = 10
    fmt.Println("x:", x)

    // 2. uint: Unsigned integer type, can be 32 or 64 bits.
    var y uint = 20
    fmt.Println("y:", y)

    // 3. float32: 32-bit floating-point numbers.
    var f1 float32 = 3.14
    fmt.Println("f1:", f1)

    // 4. float64: 64-bit floating-point numbers.
    var f2 float64 = 3.14159
    fmt.Println("f2:", f2)

    // 5. complex64: Complex numbers with float32 real and imaginary parts.
    var c1 complex64 = 3 + 4i
    fmt.Println("c1:", c1)

    // 6. complex128: Complex numbers with float64 real and imaginary parts.
    var c2 complex128 = 3.14 + 2.71i
    fmt.Println("c2:", c2)

    // 7. bool: Represents boolean values true or false.
    var b bool = true
    fmt.Println("b:", b)

    // 8. string: Represents a sequence of characters.
    var s string = "Hello, Go!"
    fmt.Println("s:", s)

    // 9. Array: Represents a fixed-size collection of elements of the same type.
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("arr:", arr)

    // 10. Slice: Represents a dynamically-sized sequence.
    var slice []int = []int{1, 2, 3, 4, 5}
    fmt.Println("slice:", slice)

    // 11. Map: Represents an unordered collection of key-value pairs.
    var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Println("m:", m)

    // 12. Pointer: Stores the memory address of another variable.
    var ptr *int = &x
    fmt.Println("ptr:", ptr)

    // 13. Struct: Represents a collection of fields.
    type Person struct {
        Name string
        Age  int
    }
    var p1 Person = Person{Name: "Alice", Age: 30}
    fmt.Println("p1:", p1)

    // 14. Function: Represents a function.
    fmt.Println("add(3, 5):", add(3, 5))

    // 15. Interface: Specifies a set of method signatures.
    var shape Shape
    fmt.Println("shape:", shape)
}

func add(x, y int) int {
    return x + y
}

type Shape interface {
    Area() float64
}
