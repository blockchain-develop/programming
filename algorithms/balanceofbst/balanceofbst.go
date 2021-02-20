package balanceofbst

import (
	"github.com/programming/utils"
)

func Balance(root *utils.BSTNode) int {
	if root == nil {
		return 0
	}
	leftBalance := Balance(root.Left)
	if leftBalance == -1 {
		return -1
	}
	rightBalance := Balance(root.Right)
	if rightBalance == -1 {
		return -1
	}
	if leftBalance > rightBalance + 1 || rightBalance > leftBalance + 1 {
		return -1
	}
	return utils.Max(leftBalance, rightBalance) +  1
}

func Height(root *utils.BSTNode) int {
	if root == nil {
		return 0
	}
	return utils.Max(Height(root.Left), Height(root.Right)) + 1
}

func IsBalance_Down2Up(root *utils.BSTNode) bool {
	balance := Balance(root)
	return balance != -1
}

func IsBalance_Up2Down(root *utils.BSTNode) bool {
	if root == nil {
		return true
	}
	diff := utils.Abs(Height(root.Left) - Height(root.Right))
	if diff > 1 {
		return false
	}
	return IsBalance_Up2Down(root.Left) && IsBalance_Up2Down(root.Right)
}
