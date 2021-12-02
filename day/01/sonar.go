package main

func diffs(vs []int) []int {
	nd := len(vs) - 1
	if nd <= 0 {
		return nil
	}

	ds := make([]int, nd)
	for i, vi := range vs {
		if i == 0 {
			continue
		}

		j := i - 1
		vj := vs[j]
		ds[j] = vi - vj
	}
	return ds
}

func deeperCount(vs []int) (count int) {
	ds := diffs(vs)
	for _, delta := range ds {
		if delta > 0 {
			count++
		}
	}
	return
}

func slidingSums(vs []int, width int) (sums []int) {
	var window []int
	for _, vi := range vs {
		window = append(window, vi)
		if len(window) < width {
			continue
		}
		if len(window) > width {
			window = window[1:]
		}

		var s int
		for _, wi := range window {
			s += wi
		}
		sums = append(sums, s)
	}
	return
}
