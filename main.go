package main 
import (
	"bufio"
	"os"
	"fmt"
	"strings"

)

func main(){
	 wordCount := 0 
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
		 countWordLength(line,&wordCount)
		for i, arg := range args {
			fmt.Printf("Arg %d: %s\n", i, arg)
		}
		
	}
	fmt.Println("word count",wordCount)
}

func countLineLength (line string)  {
	// fmt.Println(reflect.TypeOf(line))
	fmt.Println(line)
	fmt.Println("Length of line is:",len(line),"\n")
}

func countWordLength(line string, wc *int){
	stringArr := strings.Split(line," ")
	 *wc += len(stringArr)
		// for _,word := range stringArr {
		// 	fmtPrintln(word)
		// }
}