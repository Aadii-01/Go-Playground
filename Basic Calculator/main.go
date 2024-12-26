package main
import "fmt"

func main(){
	var operand1 , operand2 float64;
	var operator string

	fmt.Print("Enter first operand : ")
	fmt.Scanln(&operand1)
	fmt.Print("Enter second operand : ")
	fmt.Scanln(&operand2)
	fmt.Print("Enter operator (/, *, +, -) : ")
	fmt.Scanln(&operator)
	
	switch (operator)	{
	case "/":
		if  operand2 != 0 {
            fmt.Printf("Result: %.2f\n", operand1 / operand2)
        } else {
            fmt.Println("Error: Division by zero!")
        }
	case "+":
		fmt.Printf("%.2f", operand1 + operand2)
	case "-":
		fmt.Printf("%.2f", operand1 - operand2)
	case "*":
		fmt.Printf("%.2f", operand1 * operand2)
	default:
		fmt.Println("Invalid operator passed")
	}	
}