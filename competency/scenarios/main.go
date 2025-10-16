package scenarios

//this is the Rebase Conflict branch
import "fmt"

func f1(flavor string) string {
	return "I love the flavor of " + flavor
}

func f2(flavor string) string {
	return "I do not like the flavor of " + flavor
}

func f3(flavor string) string {
	return "I am allergic the flavor of " + flavor
}

func main() {
	fmt.Println("This is the main test file for these github conflict scenarios.")
	var flavor1 string = "blueberry"
	var flavor2 string = "mint"
	var flavor3 string = "chocolate"
	fmt.Println(f1(flavor1))
	fmt.Println(f2(flavor2))
	fmt.Println(f3(flavor3))
}
