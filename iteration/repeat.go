package iteration

func Repeat(st string) string {
	var res string
	for i := 0; i < 5; i++ {
		res += st
	}

	return res
}
