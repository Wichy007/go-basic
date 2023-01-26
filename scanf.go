package main

import "fmt"

func main() {
	var name string
	//var surname string
	fmt.Println("ป้อนชื่อนักเรียน = ")
	fmt.Scanf("%s", &name)
	//fmt.Println("ป้อนนามสกุล =")
	//fmt.Scanf("%s", &surname)
	// เก็บใว้หาวิธีแก้ว่าทำอย่างไรถึงจะสามารถรับสองค่าได้

	fmt.Println("สวัสดี ", name)
	//fmt.Println(surname)
}
