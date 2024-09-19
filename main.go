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
	if len(args) > 2{
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
	}
	// check for valid -flag
	if len(args) > 1 {
		switch args[1]{
		case "-wc":
			fmt.Println("Word count in file is:",wordCount)
		case "-cc":
			fmt.Printf("Their are %v characters on the page.\n",charCount)
		case "-help":
			fmt.Println("pending devlopment...wouldve added more features but im never using a character counter  but the flag for word count is -wc and the flag for character count is -cc...includes spaces.")
		default:
			fmt.Println("Invalid flag, use flag -help for manual on how to use command.")
		}
	}

	if len(args) == 1{
		fmt.Println("Use -help flag for manual on command.")
	}


}

func countLineLength (line string) int {
	return len(line)
}

func countWords(line string, wc *int){
	stringArr := strings.Split(line," ")
	 *wc += len(stringArr)
		
}
