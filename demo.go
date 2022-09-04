package main

import (
	"fmt"
)

func main() {
	basicAssertion()
}

func basicAssertion() {
	var a interface{}
	// go version 1.19 เป็นต้นมา เราสามารถใช้ any แทย interface{} ได้, ตัวอย่างเช่น: `var a any`
	a = 10
	fmt.Printf("%T %v\n", a, a)

	if i, ok := a.(int); ok {
		fmt.Println(i)
	}
	// การทำ type assertion สามารถรับค่าที่ 2 ได้ เป็นค่า boolean โดยจะได้ค่า true เมื่อ type ที่ assert นั้นตรงตามที่ระบุ แต่ถ้าไม่ตรงจะได้ค่า false

}

func customTypeAssertion() {
	var a, b Obj

	a = rectangle{w: 4, l: 4}
	b = triangle{w: 4, h: 4}

	fmt.Println(a.Area())
	fmt.Println(b.Area())

	if v, ok := b.(triangle); ok {
		v.h = 5
	}
	fmt.Println(b.Area())

	// เนื่องจากตัวแปร b มี type เป็น interface ของ Obj เราจึงไม่สามารถเห็น field ของ triangle ผ่านตัวแปร b ได้
	// หากต้องการ access ตรงไปที่ field ของ triangle จำเป็นต้องทำ assertion แบบนี้ก่อน
}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (r rectangle) Area() float64 {
	return r.l * r.w
}

type triangle struct {
	w, h float64
}

func (t triangle) Area() float64 {
	return t.w * t.h * 0.5
}
