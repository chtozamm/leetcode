// 2096. Step-By-Step Directions From a Binary Tree Node to Another
package directions

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	startPath := []string{}
	destPath := []string{}

	// Find paths from root to start and destination nodes
	findPath(root, startValue, &startPath)
	findPath(root, destValue, &destPath)

	directions := []string{}
	commonPathLength := 0

	// Find the length of the common path
	for commonPathLength < len(startPath) && commonPathLength < len(destPath) && startPath[commonPathLength] == destPath[commonPathLength] {
		commonPathLength++
	}

	// Add "U" for each step to go up from start to common ancestor
	for i := 0; i < len(startPath)-commonPathLength; i++ {
		directions = append(directions, "U")
	}

	// Add directions from common ancestor to destination
	directions = append(directions, destPath[commonPathLength:]...)

	return stringJoin(directions)
}

func findPath(node *TreeNode, target int, path *[]string) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		return true
	}

	// Try left subtree
	*path = append(*path, "L")
	if findPath(node.Left, target, path) {
		return true
	}
	*path = (*path)[:len(*path)-1] // Remove last character

	// Try right subtree
	*path = append(*path, "R")
	if findPath(node.Right, target, path) {
		return true
	}
	*path = (*path)[:len(*path)-1] // Remove last character

	return false
}

func stringJoin(arr []string) string {
	sb := strings.Builder{}
	for _, s := range arr {
		sb.WriteString(s)
	}
	return sb.String()
}
