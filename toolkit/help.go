package toolkit

func MergeSlices[T any](slices ...[]T) (Group [][]T) {
	Group = make([][]T, 0, len(slices))

	for _, slice := range slices {
		Group = append(Group, slice)
	}
	return Group
}
