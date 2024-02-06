package Methods

type Solution struct {
	father []int
}

func NewSolution(n int) *Solution {
	father := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
	}
	return &Solution{father: father}
}

func (s *Solution) find(u int) int {
	if u == s.father[u] {
		return u
	}
	s.father[u] = s.find(s.father[u])
	return s.father[u]
}

func (s *Solution) isSame(u, v int) bool {
	return s.find(u) == s.find(v)
}

func (s *Solution) join(u, v int) {
	u, v = s.find(u), s.find(v)
	if u == v {
		return
	}
	s.father[v] = u
}

func (s *Solution) ValidPath(n int, edges [][]int, source, destination int) bool {
	for _, edge := range edges {
		s.join(edge[0], edge[1])
	}
	return s.isSame(source, destination)
}
