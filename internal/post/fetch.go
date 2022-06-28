package post

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/types"
)

// return all the posts in the database
func (s *Service) GetAllPosts() []*types.Post {
	var posts []*types.Post

	// check cache first
	cachedPosts, err := s.RedisClient.Get("posts").Bytes()

	if err != nil {
		s.DB.Debug().Order("created_at DESC").Where("is_public = ?", true).Find(&posts)
		cachedPosts, err = json.Marshal(posts)

		if err != nil {
			logrus.Warn(err)
		}

		if err := s.RedisClient.Set("posts", cachedPosts, 20*time.Second).Err(); err != nil {
			logrus.Warn(err)
		}

		return posts
	}

	if err := json.Unmarshal(cachedPosts, &posts); err != nil {
		logrus.Warn(err)
	}

	return posts
}

// return a question based on the slug given
func (s *Service) GetPostsBySlug(slug string) types.Post {
	var post types.Post

	cachedPost, err := s.RedisClient.Get(slug).Bytes()

	if err != nil {
		s.DB.Debug().Where("slug = ?", slug).Find(&post)
		cachedPost, err := json.Marshal(post)

		if err != nil {
			logrus.Warn(err)
		}

		if err := s.RedisClient.Set(slug, cachedPost, 20*time.Second).Err(); err != nil {
			logrus.Warn(err)
		}

		return post
	}

	if err := json.Unmarshal(cachedPost, &post); err != nil {
		logrus.Warn(err)
	}

	return post
}

// return all unpublished posts
func (s *Service) GetAllUnPublishedPosts() []*types.Post {
	var posts []*types.Post
	s.DB.Debug().Order("created_at DESC").Where("is_public = ?", false).Find(&posts)
	return posts
}

func (s *Service) TotalPostCount() int64 {
	var count int64
	s.DB.Debug().Model(&types.Post{}).Count(&count)

	return count
}

func (s *Service) TotalPublishedPostCount() int64 {
	var count int64
	s.DB.Debug().Model(&types.Post{}).Where("is_public = ?", true).Count(&count)

	return count
}

func (s *Service) GetUnpublishedPosts() []*types.Post {
	var posts []*types.Post
	s.DB.Debug().Order("created_at DESC").Where("is_public = ?", false).Find(&posts)

	return posts
}
