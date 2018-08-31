package main
import "fmt"
import "net/url"
//import "strings"

func main() {
	s := "postgres://admin:admin@host.com:5432/path?k=v#f"

	//parse url
	u,err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	
	//output: postgres
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	//自动 try-catch
	password, _ := u.User.Password()
	fmt.Println(password)
}
