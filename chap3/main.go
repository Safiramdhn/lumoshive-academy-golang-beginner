package main

import (
	"fmt"
)

/*
	1/buat struct dengan properti dan inisiasi sebanyak 3 data
	2/ buat struct dengan properti nama dan score. buatkan fungsi untuk merubah score
	3/ buat 3 variable yang dapat mengakses 1 variable
	4/ 1 function 1 parameter, dapat menerima 3 object bangun datar, function untuk menghitung keliling bangun datar
	5/ buat struct lalu inisiasi. lalu bikin varible dengan tipe stuct yang bisa mengubah nilai inisiasi pertama
	6/ buat struct buat 3 data, buat function buat filter object struct tersebut, berdasarkan tahun
*/
// soal 1
type User struct {
	first_name string
	last_name  string
	age        int
}

// soal 2
type Student struct {
	name  string
	score int
}

func (s *Student) updateScore(new_score int) {
	s.score = new_score
}

// soal no 4
type Shape interface {
	Perimeter() float64
}

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func CalculatePerimeter(shape Shape) float64 {
	return shape.Perimeter()
}

func filterBook(year int, book []Book) {
	fmt.Printf("Filtered Book : above year %d\n", year)
	for _, value := range book {
		if value.Year > 2005 {
			fmt.Printf("Result %v\n", value)
		}
	}
}

// soal 6
type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	users := []User{
		{"John", "Doe", 30},
		{"Safira", "Ramadhani", 17},
		{"Bagas", "Arisandi", 20},
	}

	for _, value := range users {
		fmt.Printf("Nama: %s %s, Umur: %d\n", value.first_name, value.last_name, value.age)
	}

	fmt.Println("---------------------")

	// soal no 3
	var p1 *int
	var p2 **int
	var p3 ***int
	var mainVar int = 27

	p1 = &mainVar
	p2 = &p1
	p3 = &p2

	fmt.Println("value mainVar")
	fmt.Println("Before: ", mainVar)
	*p1 = 10
	fmt.Println("After 1st change: ", mainVar)
	**p2 = 13
	fmt.Println("After 2nd change: ", mainVar)
	***p3 = 99
	fmt.Println("After 3rd change: ", mainVar)
	fmt.Println("---------------------")

	// soal no 2
	student := Student{"Bagas", 80}
	fmt.Printf("Student %s score is changed\n", student.name)
	fmt.Println("Before : ", student.score)
	student.updateScore(100)
	fmt.Println("After : ", student.score)

	fmt.Println("---------------------")

	// soal no 5
	triangle := Triangle{A: 3, B: 4, C: 5}
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 4, Width: 5}

	fmt.Printf("Triangle perimeter: %.2f\n", CalculatePerimeter(triangle))
	fmt.Printf("Circle perimeter: %.2f\n", CalculatePerimeter(circle))
	fmt.Printf("Rectangle perimeter: %.2f\n", CalculatePerimeter(rectangle))

	fmt.Println("---------------------")

	books := []Book{
		{"Harry Potter", "J.K. Rowling", 2006},
		{"The Lord of the Rings", "J.R.R. Tolkien", 1954},
		{"The Hunger Games", "Suzanne Collins", 2008},
	}
	filterBook(2005, books)
}
