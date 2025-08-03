package main
import "fmt"
import "os"
//Println("hello")
	//func main(){
	////print(hello)
	// fmt.Println("Hello World")

//To store values in variables
	// var name string="wlug"
	// sname:="tree"
	// fmt.Println(name,sname)

//if..else
	// var i int=34
	// if(i<5){
	// 	fmt.Println("its smaller than 5")
	// } else{
	// 	fmt.Println("its bigger than 5")
	// }

//for loop
	// for i:=0;i<10;i++{
	// 	fmt.Println(i)
	// }

//Switch---case
	// var j int;
	// fmt.Scanln(&j)
	// switch j{
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// default:
	// 	fmt.Println("enter correctly")
	// }

//arrayss
	// var name[5] int;
	// name[0] = 12
	// name[1] = 13
	// name[2] = 14
	// fmt.Println(name)

//array input using for loop
	// var l[5] int;
	// for i:=0;i<5;i+=1{
	// 	l[i]=i+1
	// }
	// fmt.Println(l)

//array slicing/sizing
	// name:=make([]int,4)
	// fmt.Println(name)
	// name=append(name,2,3,4)
	// fmt.Println(name)

//maps
	// map1:=make(map[int]int)
	// map1[0]=12
	// map1[1]=14
	// map1[2]=15
	// fmt.Println(map1)


//Funtions
	//func function_name(parameter data_type)datatype{}
	//}
	
//functions
	// func add(a int , b int)int{
	// 	res:=a+b
	// 	return res
	// }
	// func main(){
	// 	fmt.Println(add(2,3))
	// }

//Pointer
	// func main(){
	// 	var x int=40
	// 	fmt.Println(x)
	// 	var myptr *int
	// 	myptr=&x
	// 	*myptr=50
	// 	fmt.Println(*myptr)
	// }

//Call by value/reference
	// func swap(a , b *int){
	// 	var c int
	// 	c=*a
	// 	*a=*b
	// 	*b=c
	// }
	// func main(){
	// 	var x,y int = 10,20
	// 	fmt.Println(x,y)
	// 	swap(&x,&y)
	// 	fmt.Println(x,y)
	// }

// File Handling
	// func main(){
	// 	var text="hello world"
	// 	file err:=	os.Create("test1.txt")
	// 		if (err!= nil){
	// 			panic(err)
	// 		}
		
	// 	}
