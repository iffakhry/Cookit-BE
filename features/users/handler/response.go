package handler

import (
	"alta-cookit-be/features/users"
	"errors"
)

type ProfileResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
	Bio            string `json:"bio"`
}

func ToProfileResponse(data users.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Username:       data.Username,
		Bio:            data.Bio,
	}
}

func ConvertUpdateResponse(input users.Core) (interface{}, error) {
	ResponseFilter := users.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.ProfilePicture != "" {
		result["profile_picture"] = ResponseFilter.ProfilePicture
	}
	if ResponseFilter.Username != "" {
		result["username"] = ResponseFilter.Username
	}
	if ResponseFilter.Bio != "" {
		result["bio"] = ResponseFilter.Bio
	}

	if len(result) <= 1 {
		return users.Core{}, errors.New("no data was change")
	}
	return result, nil
}
