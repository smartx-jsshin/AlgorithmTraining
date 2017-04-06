package main

import "fmt"

type Node struct {
	id        int
	buildtime int
	adjacent  []int
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var nodecnt, edgecnt int
		fmt.Scan(&nodecnt, &edgecnt)

		var priqueue = make([]*Node, nodecnt+1)
		for j := 1; j <= nodecnt; j++ {
			priqueue[j] = new(Node)
			priqueue[j].id = j
			fmt.Scan(&priqueue[j].buildtime)
			fmt.Println(priqueue[j])
		}

		fmt.Println("")
		for j := 0; j < edgecnt; j++ {
			var from, to int
			fmt.Scan(&from, &to)
			priqueue[from].adjacent = append(priqueue[from].adjacent, to)
			fmt.Println(priqueue[from])
		}

		var tgt int
		fmt.Scan(&tgt)

		fmt.Println(tgt)
	}
}

func calBuildTime(queue []*Node, tgt int) {

}

/*
func makeTree(times []int, rules [][2]int, target_id int) *Node {
	var nodeList []*Node = make([]*Node, len(times))

	for i := 0; i < len(times); i++ {
		var tmp *Node = new(Node)
		tmp.id = i + 1
		tmp.buildtime = times[i]
		nodeList[i] = tmp
	}

	for i := 0; i < len(rules); i++ {
		var from_id int = rules[i][0]
		var to_id int = rules[i][1]
		var to_node *Node = nodeList[to_id-1]
		to_node.child = append(to_node.child, nodeList[from_id-1])
	}
	return nodeList[target_id-1]
}

func traverseTree(n *Node) int {
	var maxval int = -1
	fmt.Println(n.id)
	if n.child == nil {
		return n.buildtime
	} else {
		for i := 0; i < len(n.child); i++ {
			var tmp int = traverseTree(n.child[i])
			if tmp >= maxval {
				maxval = tmp
			}
		}
		return n.buildtime + maxval
	}
}
*/
