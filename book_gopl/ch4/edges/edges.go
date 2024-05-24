package main

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("a", "c")
	addEdge("c", "d")
	addEdge("c", "a")
	addEdge("a", "a")

	println(hasEdge("a", "b"))
	println(hasEdge("c", "d"))
	println(hasEdge("a", "c"))
	println(hasEdge("c", "a"))
	println(hasEdge("a", "a"))
	println(hasEdge("b", "a"))
	println(hasEdge("d", "c"))
}
