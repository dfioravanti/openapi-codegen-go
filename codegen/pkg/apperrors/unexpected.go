package apperrors

import "errors"

var (
	ThisShouldNotHappenErr = errors.New("the condition that lead to this error should never have happen, please open a bug request on GitHub and attach the file you were trying to edit")
)
