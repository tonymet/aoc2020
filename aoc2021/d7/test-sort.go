// custom sort
type funny []int

func (i funny) Len() int {
	return len(i)
}
func (s funny) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s funny) Less(i, j int) bool {
	return funnyDelta(s[i]) < funnyDelta(s[j])
}
