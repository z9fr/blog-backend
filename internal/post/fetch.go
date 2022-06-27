package post

import "github.com/z9fr/blog-backend/internal/types"

// return all the posts in the database
func (s *Service) GetAllPosts() []*types.Post {
	var posts []*types.Post
	s.DB.Debug().Order("created_at DESC").Where("is_public = ?", true).Find(&posts)

	return posts
}

// return a question based on the slug given
func (s *Service) GetPostsBySlug(slug string) types.Post {
	var post types.Post
	s.DB.Debug().Where("slug = ?", slug).Find(&post)
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
