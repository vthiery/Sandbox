package main

func bubbleSort(list []int) {
	N := len(list)
	for i := 0; i < N; i++ {
		if !swap(list, i) {
			return
		}
	}
}

func swap(list []int, prevPasses int) bool {
	var (
		N       = len(list)
		i       = 0
		j       = 1
		didSwap = false
	)
	for j < (N - prevPasses) {
		listI := list[i]
		listJ := list[j]
		if listI > listJ {
			list[i] = listJ
			list[j] = listI
			didSwap = true
		}
		i++
		j++
	}
	return didSwap
}
