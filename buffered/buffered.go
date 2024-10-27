package buffered

import (
	"bufio"
	"fmt"
	"strings"
)

func BufioExample() {

	csv := `
		ID,Name,Age,Salary,Department
		1,Name_1,39,40210,Marketing
		2,Name_2,27,56792,Sales
		3,Name_3,40,53286,Engineering
		4,Name_4,28,54496,Engineering
		5,Name_5,27,108013,Sales
		6,Name_6,25,109421,Engineering
		7,Name_7,24,73668,HR
		8,Name_8,41,97826,Engineering
		9,Name_9,40,87709,Sales
		10,Name_10,41,114566,Engineering
	`

	scanner := bufio.NewScanner(strings.NewReader(csv))

	for scanner.Scan() {

		fmt.Println(scanner.Text())

	}

}
