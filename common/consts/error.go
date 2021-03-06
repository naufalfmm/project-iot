package consts

import "errors"

var (
	NotTrxError             error = errors.New("Not TRX")
	Unauthorized            error = errors.New("Unauthorized")
	UnclaimedToken          error = errors.New("Unclaimed Token")
	WrongClaimsFormat       error = errors.New("Wrong claims format")
	UndefinedLoginDataError error = errors.New("login data is not defined by token")
	UniqueError             error = errors.New("Unique error")
	NotFoundError           error = errors.New("Not Found")
)
