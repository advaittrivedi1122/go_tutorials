package main
import "fmt"
import "go_tutorials/tutorials/tutorial_1"

func sayHello() string {
	return "hello"
}

func main() {
	var myNum int = 12
	fmt.Println(myNum)

	var pi float32 = 3.14
	fmt.Println(pi)

	newNum := pi*2
	fmt.Println(newNum)

	var a int = 1
	var b float32 = 2
	fmt.Println("a + b =",float32(a) + b)
	fmt.Println("a * b =", a * int(b))
	fmt.Println("b - a =", int(b) - a)
	fmt.Println("b / a =", int(b) / int(a))

	var myStr string = `hello there`
	fmt.Println(myStr)


	var myNewNumber int = 11;
	var numPtr *int = &myNewNumber;
	fmt.Println(*numPtr)

	myStr = "hello"
	fmt.Println(len(myStr))

	myBool := false
	fmt.Println(myBool)

	const myConst int = 11
	fmt.Println(myConst)

	if (myConst == 11) {
		fmt.Println("in if")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	type user struct {
		age int;
		name string;
	}

	var myObject user;
	fmt.Println(myObject)

	myObject.age = 11;
	myObject.name = "user1";
	fmt.Println(myObject.name);

	var myArr = [4]int{1,2,3};
	fmt.Println(myArr)
	fmt.Println("myArr = ",myArr[:]);

	myStr = tutorial_1.Greet()
	fmt.Print(myStr + " there changed")
}
