//Author: Garrett Becker

package main

import "fmt"
import "regexp"
import "bufio"
import "log"
import "os"
import "strings"
import "sort"






func main() {
	var i int = 0 //counter for loop
	var g int = 0 //counter for length of line loop
	var j int = 0 //counter for looping through key
	var words1 string
	var words3 string
	scanner := bufio.NewScanner(os.Stdin)
	holdWords := make(map[string]int)
	//Loop through the txt doc and cleanup the lines of strings
	for scanner.Scan(){
		i=i+1

		words1=scanner.Text(); //get the text from line
		words1=strings.TrimSpace(words1) //trim the trailing and leading whitespace
		words1=strings.ToLower(words1)
		reg, err:=  regexp.Compile("^[a-zA-Z]+") //regex to look for letters/spaces
		
		if err != nil{
			log.Fatal(err)
		}
		
		var words2 = strings.Fields(words1)
		
		for g < len(words2){
		words3 = reg.FindString(words2[g])
		g=g+1
		holdWords[words3] = i //stores the line number with each string

	}
	g=0 //reset the counter
		
	}
	keys := make([]string, 0, len(holdWords))
    for key := range holdWords {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    
    
    for j < len(keys){
    	fmt.Println(keys[j],holdWords[keys[j]])
    	j=j+1

    }
	

}

//1. cleanup the inputted string except for letters and `
//2. loop through the new string and store each invidual string into a map along with it's corresponding line number
//3. then pull out each of the strings and sort them and then display them with their line number(s)