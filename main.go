import (
	"sort"
)

type aa struct {
	a, b, idx int
}

func cmp(a, b aa) bool {
	return a.a > b.a
}

func cmp2(a, b pair) bool {
	if a.first == b.first {
		return b.second < a.second
	}
	return a.first > b.first
}

type pair struct {
	first, second int
}

var v []aa

func initialize(scores [][]int) {
	for i := 0; i < len(scores); i++ {
		v = append(v, aa{scores[i][0], scores[i][1], i})
	}
}

func solution(scores [][]int) int {
	initialize(scores)
	sort.Slice(v, func(i, j int) bool {
		return cmp(v[i], v[j])
	})
	var ans []pair
	for i := 0; i < len(v); i++ {
		flag := true
		r := v[i]
		for t := 0; t < i; t++ {
			l := v[t]
			if l.a > r.a && l.b > r.b {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		ans = append(ans, pair{r.a + r.b, r.idx})
	}
	sort.Slice(ans, func(i, j int) bool {
		return cmp2(ans[i], ans[j])
	})
	answer := 0
	rank := 1
	for i := 0; i < len(ans); i++ {
		a := ans[i].first
		cnt := 0
		for i < len(ans) && a == ans[i].first {
			if ans[i].second == 0 {
				answer = rank
				rank++
				break
			}
			i++
			cnt++
		}
		rank += cnt
		i--
		if answer != 0 {
			break
		}
	}
	if answer == 0 {
		return -1
	}
	return answer
}
