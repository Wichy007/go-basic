package main

import "fmt"

func main() {
	name := "naravit"          //วิธีการกำหนดรูปแบบของข้อมูลแบบ manual type declaration
	var age int = 25           //วิธีการกำหนดรูปแบบของข้อมูลแบบ type inference
	const score float32 = 95.8 //การนิยามตัวแปรที่เป็นค่าคงที่เปลี่ยนแปลงไม่ได้
	// score = 20 จะไม่สามรถเปลี่ยนแปลงค่าของตัวแปรได้เพระกำหนดเป็นค่าคงที่ไปแล้ว
	var ispass bool = true

	fmt.Println("my name is ", name)
	fmt.Println("age ", age)
	fmt.Println("score ", score)
	fmt.Println("pass exam ", ispass)
}
