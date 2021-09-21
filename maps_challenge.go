/* Alta3 Research | RZFeeser
   Map - hosts to IP:port resolution  */

package main

import "fmt"

func main(){ 
	var fileExtensions = map[string]string{
	"python": ".py",
	"c++": ".cpp",
	"Java": ".java",
	"Golang": ".go",
	"Ansible": ".yml",
	}
	fmt.Println(fileExtensions)

}

