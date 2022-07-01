package main


import ("fmt"
		"time")


func main(){

	c := make(chan int)


	for i:=0;i<10;i++{
		go proccess(i,c)
	}

	for i :=0;i<10;i++{
		fmt.Println("teminado el programa en", <- c)
	}

}

func proccess(i int, c chan int){
	fmt.Println(i,"-Inicia")
	time.Sleep(1000*time.Millisecond)
	fmt.Println(i,"-Termina")
	c <- i
}



