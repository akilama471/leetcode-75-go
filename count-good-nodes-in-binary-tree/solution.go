/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxSoFar int) int {
    if node == nil {
        return 0
    }

    good := 0
    if node.Val >= maxSoFar {
        good = 1
    }

    // Update the maximum value along the path
    newMax := max(maxSoFar, node.Val)

    // Recurse on left and right subtrees
    good += dfs(node.Left, newMax)
    good += dfs(node.Right, newMax)

    return good
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
