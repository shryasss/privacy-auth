package utils

const NumPublicInputs = 3

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
