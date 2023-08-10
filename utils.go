package dataparse

func ListToMap[K comparable](input []K) map[K]bool {
	m := make(map[K]bool, len(input))
	for _, member := range input {
		m[member] = true
	}
	return m
}

func FilterSlice[V comparable](in []V, removees ...V) []V {
	removeeMap := ListToMap(removees)

	ret := []V{}
	for _, member := range in {
		if _, ok := removeeMap[member]; ok {
			continue
		}
		ret = append(ret, member)
	}

	return ret
}
