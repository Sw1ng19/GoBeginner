package main
import "fmt"
import "os/exec"

func main() {
	cmd := exec.Command("date")

	date, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(date))
}
