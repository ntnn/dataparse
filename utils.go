package dataparse

// ListToAny returns a copy of the passed list as []any.
func ListToAny[V any](input []V) []any {
	ret := make([]any, len(input))
	for i := range input {
		ret[i] = any(input[i])
	}
	return ret
}

// ListToMap returns a map where each of the list members is a key and
// each value is true.
// Utilized to take a list of options and transform it into a map for
// lookups.
func ListToMap[K comparable](input []K) map[K]bool {
	m := make(map[K]bool, len(input))
	for _, member := range input {
		m[member] = true
	}
	return m
}

// FilterSlice returns a copy of the passed slice with the removees
// removed.
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
