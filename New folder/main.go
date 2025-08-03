package main
import (
	"fmt"
	"os"
	features "student-db/Features"
)
func DisplayChoices(){
	fmt.Println("1.Add Students")
	fmt.Println("2.Delete Students")
	fmt.Println("3.Update Students")
	fmt.Println("4.Show all Students")
	fmt.Println("5.Exit")
}
func main(){
	fmt.Println("Hello DB")
	var myclass features.Class
	myclass.NewClass()	

	for{
		DisplayChoices()
		var input int
		fmt.Scanln(&input)

		switch input{
		case 1:
			myclass.AddStudent()
		case 2:
			myclass.DeleteStudent()
		case 3:
			myclass.UpdateStudent()
		case 4:
			myclass.ShowStudents()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invlaid Choice")
		}
	}
}