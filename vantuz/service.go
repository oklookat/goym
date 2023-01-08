package vantuz

// status >= 400
func isHttpError(status int) bool {
	return status >= 400
}

// status >= 200 and <= 299
func isHttpSuccess(status int) bool {
	return status >= 200 && status <= 299
}
