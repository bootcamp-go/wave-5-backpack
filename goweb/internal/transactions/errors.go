package transactions

type NotAllowedAmountZeroOrNegative struct {
}

func (n *NotAllowedAmountZeroOrNegative) Error() string {
	return "error: amount is zero or below"
}
