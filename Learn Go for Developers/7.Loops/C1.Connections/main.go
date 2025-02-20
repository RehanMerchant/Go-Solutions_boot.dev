package main

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		connections += i
	}
	return connections
}
