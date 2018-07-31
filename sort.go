package main
import "fmt"
import "sort"

func main() {
	//数组排序
	strs := []string{"c","b","a"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	nums := []int{7,9,8}
	sort.Ints(nums)
	fmt.Println("Ints: ", nums)
}
