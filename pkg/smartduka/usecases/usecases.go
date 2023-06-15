package usecases

import (
	"github.com/oryx-systems/smartduka/pkg/smartduka/usecases/otp"
	"github.com/oryx-systems/smartduka/pkg/smartduka/usecases/user"
)

// Smartduka manages the usecases intrefaces
type Smartduka struct {
	User user.UseCasesUser
	OTP  otp.UseCasesOTP
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewSmartdukaUsecase(
	user user.UseCasesUser,
	otp otp.UseCasesOTP,
) *Smartduka {
	m := &Smartduka{
		User: user,
		OTP:  otp,
	}

	return m
}
