package _134_加油站


func canCompleteCircuit(gas []int, cost []int) int {
	total, curr := 0, 0
	n := len(gas)
	start := 0
	for i:=0; i<n; i++{
		total += gas[i] - cost[i]
		curr += gas[i] - cost[i]
		if curr < 0{
			start = i+1
			curr = 0
		}
	}
	if total >= 0{
		return start
	}else {
		return -1
	}
}
