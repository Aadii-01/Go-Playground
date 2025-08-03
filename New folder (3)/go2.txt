//go routines are basically threads (smallest part of a process)
//go routines never keep a track of the state of the other routine 
//and never know whether the other routine is throwing an error or not

//sequential execution
//parallel execution
//concurrent execution

//time.sleep
	//import "time"
	//func name(){
	//time.Sleep(100 * time.Millisecond)}
	//the above code snippet is used to give delay to the paticular block of code
	//to check the concurrency of the code
	//but its not an efficient way as child go routine depends on main routines
	//problems with this:-
	// go routines never return anything 
	//main routine doesnt care abt child go routine
	// package main
	// import "fmt"
	// import "time"
	// func printnum(){
	// 	for i:=1;i<=10;i++{
	// 		fmt.Println(i)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
	// func printchar(){
	// 	for i:='a';i<='e';i++{
	// 		fmt.Printf("%c \n",i)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
	// func main(){
	// 	go printnum()
	// 	go printchar()
	// 	time.Sleep(1500 * time.Millisecond)
	// }

//waitgroups
	//waits until all the child go routines are executed 
	// package main
	// import "fmt"
	// import "sync"
	// func printnum(wg *sync.WaitGroup){
	// 	defer wg.Done()
	// 	for i:=1;i<=10;i++{
	// 		fmt.Println(i)
	// 	}
	// }
	// func printchar(wg *sync.WaitGroup){
	// 	defer wg.Done()
	// 	for i:='a';i<='e';i++{
	// 		fmt.Printf("%c \n",i)
	// 	}
	// }
	// func main(){
	// 	var wg sync.WaitGroup
	// 	wg.Add(1)
	// 	go printnum(&wg)
	// 	wg.Add(1)
	// 	go printchar(&wg)
	// 	wg.Wait()

	// }

//channels
	//provide synchroinzation
	// package main
	// import "fmt"
	// import "time"
	// func sum(a,b int,ch chan int){
	// 	s:=0
	// 	for i:=a;i<=b;i++{
	// 		s+=1
	// 	}
	// 	ch <- s
	// 	time.Sleep(100* time.Millisecond)
	// 	fmt.Println("WLUG")
	// }
	// func main(){
	// 	ch:=make(chan int)
	// 	go sum(1,100,ch)
	// 	output:= <- ch
	// 	fmt.Println(output)
	// }
//mutexes
	// package main
	// import "fmt"
	// import "sync"
	// func increment(a *int,mu *sync.Mutex,wg *sync.WaitGroup){
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	*a++
	// }
	// func decrement(a *int,mu *sync.Mutex,wg *sync.WaitGroup){
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	*a--
	// }
	// func main(){
	// 	a:=8
	// 	var mu sync.Mutex
	// 	var wg sync.WaitGroup
	// 	wg.Add(1)
	// 	increment(&a,&mu,&wg)
	// 	wg.Add(1)
	// 	decrement(&a,&mu,&wg)
	// 	fmt.Println(a)
	// 	wg.Wait()
	// }
//