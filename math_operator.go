package main

import "fmt"

func main() {
	var number1 int = 10
	number2 := 3
	var number3, number4 int = 5, 9 //สามารถกำหนดตัวแปรในบรรทัดเดียวมากกว่าหนึ่งตัวแปรได้ดังรูป
	fmt.Println("ผลบวก = ", number1+number2)
	fmt.Println("ผลบวก = ", number1-number2)
	fmt.Println("ผลบวก = ", number1*number2)
	fmt.Println("ผลบวก = ", number1/number2)
	fmt.Println("ผลบวก = ", number1%number2)
	//การดำเนินการเปรียบเทียบ
	fmt.Println("number3>number4 =", number3 > number4)
	fmt.Println("number3<number4 =", number3 < number4)
	fmt.Println("number3==number4 =", number3 == number4)
	fmt.Println("number3!=number4 =", number3 != number4)
	fmt.Println("number3>=number4 =", number3 >= number4)
	fmt.Println("number3<=number4 =", number3 <= number4)

}
