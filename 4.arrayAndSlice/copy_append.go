package main
import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)//size 10

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	s6 := 12
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	for i:=0;i< len(sl3);i++  {
		fmt.Printf("%x\n",&sl3[i])
	}
	//slice内存中是连续的
	fmt.Printf("%x\n",&s6)
}