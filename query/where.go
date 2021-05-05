package query

// TODO: Corresponding to "AND", "OR" and "BETWEEN" etc..
func makeWhereFunc(targetIndex int, q *Query) execFunc {
	comparator := makeComparator(q.Statement, q.Value)

	return func(data [][]string) ([][]string, error) {
		results := make([][]string, 0, len(data))
		for _, row := range data {
			if comparator(row[targetIndex]) {
				results = append(results, row)
			}
		}
		return results, nil
	}
}

func makeComparator(state, value string) func(string) bool {
	switch state {
	case "=":
		return func(target string) bool {
			if target == value {
				return true
			}
			return false
		}
	case "<":
		return func(target string) bool {
			if target < value {
				return true
			}
			return false
		}
	case ">":
		return func(target string) bool {
			if target > value {
				return true
			}
			return false
		}
	case "<=":
		return func(target string) bool {
			if target <= value {
				return true
			}
			return false
		}
	case ">=":
		return func(target string) bool {
			if target >= value {
				return true
			}
			return false
		}
	}
	return nil
}
