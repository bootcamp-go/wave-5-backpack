package transactions

import "fmt"

type NotAllowedAmountZeroOrNegative struct {
}

func (n *NotAllowedAmountZeroOrNegative) Error() string {
	return "error: amount is zero or below"
}

type NotFound struct {
	fieldName   string
	searchValue string
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("error: attribute %s doesn`t have value %s", n.fieldName, n.searchValue)
}
