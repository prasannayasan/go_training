package main

import (
	"fmt"
)

// var number int = 5

func main() {
	// fmt.Println("Hello Go World")
	// fmt.Println(number)

	// name := "Check"
	// fmt.Println(name)
	// age := 5
	// _ = age

	// fmt.Println("my age %d", 25)
	// fmt.Printf("my age %yd", 25)

	// s := fmt.Sprintf("a is %d, b is %f, c is %s \n", 5, 85.68, "g ahuih")
	// fmt.Println(s)

	// const (
	// 	min1 = 98
	// 	min2 = 8
	// 	min3
	// )
	// fmt.Println(min1, min2, min3)

	// const (
	// 	min4 = iota * 9
	// 	min5
	// 	_
	// 	min6
	// )
	// fmt.Println(min4, min5, min6)

	// var cities = [4]string{"chennai", "banglore", "Coorg"}
	// fmt.Println(cities)

	// maps := map[int]string{
	// 	1: "ONE",
	// 	2: "TWO",
	// }
	// fmt.Println(maps)

	// var x = 2
	// ptr := &x
	// fmt.Println(ptr)

	// type Person struct {
	// 	name string
	// 	age  int
	// }
	// var myself Person
	// fmt.Printf("Type is %T and value is %v", myself, myself)
	// myself.name, myself.age = "Prasanna", 26
	// fmt.Printf("Type is %T and value is %v", myself, myself)

	// var i int = 5
	// var j float64 = 3.46
	// fmt.Println(float64(i) * j)
	// fmt.Println(i * int(j))

	// str := fmt.Sprintf("%d", i)
	// fmt.Println(str, "")
	// str = fmt.Sprintf("%f", j)
	// fmt.Println(str)

	// k := strconv.Itoa(i)
	// fmt.Println(k)

	// y, err := strconv.ParseFloat("df", 69)
	// _, _ = y, err

	// fmt.Println("OS.Args : ", os.Args[0])

	// if k, er := strconv.Atoi(os.Args[1]); k >= 5 {
	// 	fmt.Println(k + 1)
	// 	fmt.Println(er)
	// } else {
	// 	fmt.Println(k * 6)
	// }

	// var arr1 [5]int
	// var arr2 = []float64{}
	// arr3 := [...]string{"Ram", "Sam", "Yam", "Ham"}
	// fmt.Printf("%v\n %v\n %v\n", arr1, arr2, arr3)

	// for i, v := range arr3 {
	// 	fmt.Printf("index is %d and value is %s\n", i, v)
	// }

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Printf("index is %d and value is %s\n", i, arr3[i])
	// }

	// var sli = []int{}
	// fmt.Printf("%#v\n", sli)
	// // sli[0] = 7

	// sli1 := make([]int, 10)
	// fmt.Printf("%#v", sli1)

	// a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// var eq bool = true
	// if len(a) != len(b) {
	// 	eq = false
	// } else {
	// 	for i, valueA := range a {
	// 		if valueA != b[i] {
	// 			eq = false // don't check further, break!
	// 			break
	// 		}
	// 	}
	// }

	// if eq {
	// 	fmt.Println("a and b slices are equal")
	// } else {
	// 	fmt.Println("a and b slices are not equal")
	// }

	// s1 := []int{1, 2, 3, 4, 5, 6}
	// s1 = append(s1[:4], 100)
	// fmt.Println(s1)

	// s2, s3 := s1[:4], s1[3:]
	// s4 := append(s1, 85, 54)
	// s1[3] = 678
	// fmt.Println(s1, s2, s3, s4)
	// fmt.Printf("%p\n", &s1)

	// years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	// newYears := append(years[0:3], years[len(years)-3:]...)
	// fmt.Printf("%#v\n", newYears)

	// var1 := 'a'
	// fmt.Printf("Type: %T, Value:%v \n", var1, var1)

	// s2 := "中文维基是世界上"
	// r1 := []rune(s2)
	// fmt.Println(string(r1[0:4]))

	// str2 := "56151.sdfuhsf.sgerg.sdfgsdfg.sdfv.esrfs.f.se.f.s=ef=sef.d:hthd.erh"
	// fmt.Println(strings.ReplaceAll(str2, ".", "->"))
	// fmt.Println(strings.Trim(str2, ".:="))

	// var ru = 'ă'
	// fmt.Printf("%T/n", ru)
	// ma, m := "ma", "m"
	// str3 := ma + m + string(ru)
	// fmt.Println(str3)

	// s1 := "țară means country in Romanian"
	// for _, v := range s1 {
	// 	fmt.Printf("%d --- %q\n", v, v)
	// }

	// var employees map[string]string
	// employees = make(map[string]string)
	// employees["lead"] = "rajesh"
	// fmt.Printf("%#v\n", employees)
	// fmt.Println(employees)

	// os.Create("a.txt")
	// os.Truncate("a.txt", 0)
	// os.Remove("a.txt")

	// type company = struct {
	// 	name         string
	// 	startingYear int
	// }
	// type employee = struct {
	// 	name    string
	// 	age     int
	// 	company company
	// }
	// accenture := company{"Accenture", 1992}
	// prasanna := employee{
	// 	"Prasanna",
	// 	25,
	// 	accenture,
	// }
	// fmt.Println(prasanna)

	// x := increment(16)
	// fmt.Println(x)
	// fmt.Println(x())
	// fmt.Println(x())

	// mycar := Car{"Audi", 2018}
	// mycar.changeCar("BMW", 1997)
	// fmt.Println(mycar)

	// l1 := Lion{"Sher"}
	// m1 := Man{"Prasanna"}
	// print(l1)
	// print(m1)

	// fmt.Println("CPUs available : ", runtime.NumCPU())
	// fmt.Println("Threads available : ", runtime.GOMAXPROCS(0))

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func1(&wg)
	// wg.Wait()

	// var url = "https://www.google.com"
	// res, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	defer res.Body.Close()
	// 	fmt.Println("Writing res into file")
	// 	byteSlice, _ := ioutil.ReadAll(res.Body)
	// 	file := strings.Split(url, "//")[1] + ".txt"
	// 	fmt.Println(file)
	// 	err = ioutil.WriteFile(file, byteSlice, 0664)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// var ch = make(chan int)
	// go f5(ch, 67)
	// val := <-ch
	// fmt.Println(val)
	// defer close(ch)

	// for i := 1; i <= 3; i++ {
	// 	go func(c chan int, n int) {
	// 		k := 1
	// 		for j := 1; j <= i; j++ {
	// 			k *= j
	// 		}
	// 		c <- k
	// 	}(ch, i)
	// }
	// for l := 1; l <= 20; l++ {
	// 	fmt.Printf("The factorial of %d is %d\n", l, <-ch)
	// }

	u1 := IsEven(2)
	fmt.Printf("is number %d even :%v\n", 2, u1)
}

// func f5(c1 chan int, n int) {
// 	c1 <- n
// }

// func func1(wg *sync.WaitGroup) {
// 	for i := 1; i > 5; i++ {
// 		fmt.Println("From func1 : ", i)
// 	}
// 	wg.Done()
// }

// type Animal interface {
// 	run() int
// 	speak()
// }

// type Lion struct {
// 	name string
// }

// type Man struct {
// 	name string
// }

// func (ln Lion) run() int {
// 	return 50
// }
// func (ln Lion) speak() {
// 	fmt.Println("Lion can't speak")
// }
// func (m Man) run() int {
// 	return 50
// }
// func (m Man) speak() {
// 	fmt.Println("Lion can't speak")
// }
// func print(an Animal) {
// 	an.run()
// 	an.speak()
// }

// type Car struct {
// 	name     string
// 	makeYear int
// }

// func (car *Car) changeCar(name1 string, make1 int) {
// 	(*car).name = name1
// 	(*car).make = make1
// }

// func increment(x int) func() int {
// 	return func() int {
// 		x = x * 5
// 		return x
// 	}
// }
