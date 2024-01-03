package Methods

func EvalRPN(Arr []byte) int {
	tmp := make([]byte, 0)

	for len(Arr) > 1 {
		for i := range Arr {
			if Arr[i] == '+' {
				Arr[i] = Arr[i-2] + Arr[i-1]
				tmp = append(tmp, Arr[:i-2]...)
				tmp = append(tmp, Arr[i:]...)
				Arr = tmp
				tmp = make([]byte, 0)
				break

			} else if Arr[i] == '-' {
				Arr[i] = Arr[i-2] - Arr[i-1]
				tmp = append(tmp, Arr[:i-2]...)
				tmp = append(tmp, Arr[i:]...)
				Arr = tmp
				tmp = make([]byte, 0)
				break

			} else if Arr[i] == '*' {
				Arr[i] = Arr[i-2] * Arr[i-1]
				tmp = append(tmp, Arr[:i-2]...)
				tmp = append(tmp, Arr[i:]...)
				Arr = tmp
				tmp = make([]byte, 0)
				break

			} else if Arr[i] == '-' {
				Arr[i] = Arr[i-2] / Arr[i-1]
				tmp = append(tmp, Arr[:i-2]...)
				tmp = append(tmp, Arr[i:]...)
				Arr = tmp
				tmp = make([]byte, 0)
				break

			}
		}

	}

	return int(Arr[0])
}
