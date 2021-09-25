package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Reference
	//(https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)

	//step-1
	// cmd := exec.Command("ls", "-lah")
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("tasklist")
	// }
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	//step-2
	// cmd := exec.Command("ls", "-lah")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	//step - 2 (modified)
	// var b bytes.Buffer
	// cmd := exec.Command("ls", "-lah")
	// cmd.Stdout = &b
	// cmd.Stderr = &b
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// bs := b.Bytes()
	// fmt.Println(string(bs))

	//step-3
	// cmd := exec.Command("ls", "-lah")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// fmt.Printf("combined out:\n%s\n", string(out))

	//data, err := ioutil.ReadAll(os.Stdin)
	// data, err := io.ReadAll(os.Stdin)
	// fmt.Println(data, err)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
