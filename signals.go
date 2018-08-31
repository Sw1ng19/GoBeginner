package main
import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	sigs := make(chan os.Signal, 1) //定义通道接受消息为系统信号
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("waiting for signs!")
	//阻塞等待消息
	<- done
	fmt.Println("exit")
}
