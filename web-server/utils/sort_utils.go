package utils

// [] int {9,8,7,6,5}
// [] int {5,6,7,8,9}
func BubbleSort(elems []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elems)-1; i++ {
			if elems[i] > elems[i+1] {
				elems[i], elems[i+1] = elems[i+1], elems[i]
				keepRunning = true
			}
		}
	}
}
