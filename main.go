package main 
import (
	"bufio"
	"os"
	"fmt"
	"reflect"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()

		fmt.Println(reflect.TypeOf(line))
		fmt.Println(line)
		fmt.Println("Length of line is:",countLineLength(line))
		
	}
}

func countLineLength (line string) int {
	return len(line)
}