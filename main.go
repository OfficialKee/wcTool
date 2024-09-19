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
	const maxArgs  = 3
	var charCount int

	//limit amount of args on command line
	 if len(args) - 1 > maxArgs {
        fmt.Printf("Error: Too many arguments. Maximum allowed is %d\n", maxArgs)
        os.Exit(1)
    }
	 
	//check for valid file
	file, err := os.Open(args[2])
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		countWords(line,&wordCount)
		charCount += countLineLength(line)
	}

	switch args[1]{
	case "wc":
		fmt.Println("Word count in file is:",wordCount)
	case "cc":
		fmt.Printf("Their are %v characters on the page.\n",charCount)
	default:
		fmt.Println("use flag -help for info on how to use command")
	}
	
}

func countLineLength (line string) int {
	return len(line)
}

func countWords(line string, wc *int){
	stringArr := strings.Split(line," ")
	 *wc += len(stringArr)
		
}
