package handler

import (
	"alta-cookit-be/features/users"
	"mime/multipart"
)

type RegisterReq struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateProfileReq struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Username       string `json:"username" form:"username"`
	Bio            string `json:"bio" form:"bio"`
	FileHeader     multipart.FileHeader
}

type UpdatePasswordReq struct {
	Password             string `json:"old_password" form:"old_password"`
	NewPassword          string `json:"new_password" form:"new_password"`
	ConfirmationPassword string `json:"confirmation_password" form:"confirmation_password"`
}

type ApprovementReq struct {
	Approvement string `json:"approvement" form:"approvement"`
	Role        string `json:"role" form:"role"`
}

func ReqToCore(data interface{}) *users.Core {
	res := users.Core{}

	switch data.(type) {
	case RegisterReq:
		cnv := data.(RegisterReq)
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	case LoginReq:
		cnv := data.(LoginReq)
		res.Username = cnv.Username
		res.Password = cnv.Password
	case UpdateProfileReq:
		cnv := data.(UpdateProfileReq)
		res.ProfilePicture = cnv.ProfilePicture
		res.Username = cnv.Username
		res.Bio = cnv.Bio
	case UpdatePasswordReq:
		cnv := data.(UpdatePasswordReq)
		res.Password = cnv.Password
		res.NewPassword = cnv.NewPassword
		res.ConfirmationPassword = cnv.ConfirmationPassword
	case ApprovementReq:
		cnv := data.(ApprovementReq)
		res.Approvement = cnv.Approvement
		res.Role = cnv.Role
	default:
		return nil
	}
	return &res
}
