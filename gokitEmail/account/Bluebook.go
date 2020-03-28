package account

import (
	"fmt"
)

func main() {
	var bluebookAdress [100]string
	bluebookAdress[0] = "Henry@here.com"
	bluebookAdress[1] = "Josh@here.com"
	bluebookAdress[2] = "AJ@here.com"

	bluebookAdress[3] = "David@there.com"
	bluebookAdress[4] = "Marie@there.com"
	bluebookAdress[5] = "Scott@there.com"

	fmt.Println(bluebookAdress)

}

func address(address string, idx int) {
	fmt.Println(address)
	fmt.Println("the value of the index is: ", idx)
}
