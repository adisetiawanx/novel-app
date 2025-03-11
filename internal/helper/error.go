package helper

func HandlerIfError(err error) {
	if err != nil {
		panic(err)
	}
}
