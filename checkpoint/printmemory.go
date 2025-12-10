package checkpoint 

import "fmt"

func PrintMemory(arr [10]byte) {

	for i:=0; i <10 ; i++{
		fmt.Printf("%02x", arr[i])
		if i ==3 || i ==7{
			fmt.Println()
		}else if i == 9{
			//fmt.Print(" ")
			continue
		}else{
			fmt.Print(" ")
		}
	}

	fmt.Println()

	for _ , r := range arr{
		if r >= 32 && r < 127{
			fmt.Printf("%c",r)		
		}else{
			fmt.Print(".")
		}
	}
	fmt.Println()
}
