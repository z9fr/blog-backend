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
