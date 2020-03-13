package just_for_func

func Ints(vs []int)int  {
	if len(vs) == 0 {
		return 0
	}
	return Ints(vs[1:])+vs[0]
}

