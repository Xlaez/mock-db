package input

import "fmt"

/* a function that checks if the number of students data
to be inputed on the database is not more than 10 and not
 equals 0 but it can be less than 10 and equals to 1 */
func NumChecker(num int) (int, error) {
	if num > 10 || num <= 0 {
		return 0, fmt.Errorf("number of students shouldn't be more than 10 ")
	}
	return num, nil
}
