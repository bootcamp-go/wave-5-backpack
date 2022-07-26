package users

import "fmt"

type NotFound struct {
	fileName    string
	searchValue string
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("error: attribute %s doesn´t have value %s", n.fileName, n.searchValue)
}
