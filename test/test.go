package main

func main() {

}

func append0(tmp []int) {
	ans := make([][]int, len(tmp))
	for idx := range tmp {
		ans[0][idx] = tmp[idx]
	}
}

func append1(tmp []int) {
	ans := make([][]int, 0)
	ans = append(ans, append([]int{}, tmp...))
}

func append2(tmp []int) {
	ans := make([][]int, 0)
	tmpCopy := make([]int, len(tmp))
	copy(tmpCopy, tmp)
	ans = append(ans, tmpCopy)
}
