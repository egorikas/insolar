package tests

type Cases struct {
	Input         string
	ExpectedError TestError
}

type CasesInt struct {
	Input         int32
	ExpectedError TestError
}

type TestError struct {
	Code    int32
	Message string
}
