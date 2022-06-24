package helper

import (
	"fmt"
	"github.com/z9fr/blog-backend/internal/user"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/utils"
)

func Createpostdetails(data types.PostCreateRequest, userdetails user.User) (types.Post, error) {

	if data.Title == "" || data.Body == "" || data.Description == "" || data.HeaderImage == "" {
		return types.Post{}, fmt.Errorf("Missing fields")
	}

	if !utils.IsUrl(data.HeaderImage) {
		logrus.Warn(fmt.Sprintf("Invalid url (%s) parsed as the header image by user %s", data.HeaderImage, userdetails.Email))
		return types.Post{}, fmt.Errorf("Invalid url for Header Image (%s)", data.HeaderImage)
	}

	return types.Post{
		Title:       data.Title,
		Body:        data.Body,
		Descrption:  data.Description,
		IsPublic:    false,
		Tags:        data.Tags,
		HeaderImage: data.HeaderImage,
		CreatedBy:   userdetails.UserName,
	}, nil
}
