package main 
import (
	"bufio"
	"os"
	"fmt"

)

func main(){
	args := os.Args
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		countLineLength(line)
		for i, arg := range args {
			fmt.Printf("Arg %d: %s\n", i, arg)
		}

		
	}
}

func countLineLength (line string)  {
	// fmt.Println(reflect.TypeOf(line))
	fmt.Println(line)
	fmt.Println("Length of line is:",len(line),"\n")
}