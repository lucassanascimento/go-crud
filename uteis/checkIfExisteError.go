package uteis

func CheckIfExisteError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
