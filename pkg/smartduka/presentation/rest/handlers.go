package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/dto"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/usecases"
)

// AcceptedContentTypes is a list of all the accepted content types
var AcceptedContentTypes = []string{"application/json", "application/x-www-form-urlencoded"}

// PresentationHandlers represents all the REST API logic
type PresentationHandlers interface {
	HandleLoginByPhone() gin.HandlerFunc
	HandleRegistration() gin.HandlerFunc
	SetUserPIN() gin.HandlerFunc
	GetUserProfileByPhoneNumber() gin.HandlerFunc
}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	usecases usecases.Smartduka
}

// NewPresentationHandlers initializes a new rest handlers usecase
func NewPresentationHandlers(usecases usecases.Smartduka) PresentationHandlers {
	return &PresentationHandlersImpl{usecases: usecases}
}

// HandleIncomingMessages handles and processes data posted by AIT to its callback URL
func (p PresentationHandlersImpl) HandleRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		c.Accepted = append(c.Accepted, AcceptedContentTypes...)

		payload := &dto.RegisterUserInput{}
		utils.DecodeJSONToTargetStruct(c.Writer, c.Request, payload)

		err := p.usecases.User.RegisterUser(ctx, payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Successfully registered user",
		})
	}
}

// HandleRegistration handles the user registration
func (p PresentationHandlersImpl) HandleLoginByPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		c.Accepted = append(c.Accepted, AcceptedContentTypes...)

		payload := &dto.LoginInput{}
		utils.DecodeJSONToTargetStruct(c.Writer, c.Request, payload)
		if payload.PhoneNumber == "" {
			err := fmt.Errorf("phone number is required")
			utils.ReportErr(c.Writer, err, http.StatusBadRequest)
			return
		}

		response, err := p.usecases.User.Login(ctx, payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "Successfully logged in user",
			"response": response,
		})
	}
}

// SetUserPIN handles the setting of the user PIN
func (p PresentationHandlersImpl) SetUserPIN() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		c.Accepted = append(c.Accepted, AcceptedContentTypes...)

		payload := &dto.UserPINInput{}
		utils.DecodeJSONToTargetStruct(c.Writer, c.Request, payload)
		if payload.UserID == "" {
			err := fmt.Errorf("user ID is required")
			utils.ReportErr(c.Writer, err, http.StatusBadRequest)
			return
		}

		ok, err := p.usecases.User.SetUserPIN(ctx, payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"setPIN": ok,
			"status": "Successfully set user PIN",
		})
	}
}

// GetUserProfileByPhoneNumberuser search
func (p PresentationHandlersImpl) GetUserProfileByPhoneNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		c.Accepted = append(c.Accepted, AcceptedContentTypes...)

		// get phone from query params and validate
		phone := c.Query("phoneNumber")
		if phone == "" {
			err := fmt.Errorf("phone number is required")
			utils.ReportErr(c.Writer, err, http.StatusBadRequest)
			return
		}

		users, err := p.usecases.User.SearchUserByPhoneNumber(ctx, phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Successfully searched for users",
			"users":  users,
		})
	}
}
