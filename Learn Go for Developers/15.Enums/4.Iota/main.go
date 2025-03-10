package main

type emailStatus int

const (
	emailBounced   emailStatus = iota // 0
	emailInvalid                      // 1
	emailDelivered                    // 2
	emailOpened                       // 3
)
