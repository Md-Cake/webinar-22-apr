package main

import "fmt"

type PostSource interface {
	NewPost(author *User, text string) (*Post, error)
	GetPostsByAuthors(ids []int) []*Post
}

type postSource struct {
	posts []*Post
}

func NewPostSource() PostSource {
	return &postSource{}
}

func (ps *postSource) NewPost(author *User, text string) (*Post, error) {
	p, err := NewPost(author, text)
	if err != nil {
		return nil, err
	}

	ps.posts = append(ps.posts, p)
	return p, err
}

func (ps *postSource) GetPostsByAuthors(ids []int) []*Post {
	var posts []*Post
	for _, post := range ps.posts {
		for _, id := range ids {
			if post.Author.Id == id {
				posts = append(posts, post)
				break
			}
		}
	}
	return posts
}

type Post struct {
	Author *User
	Text   string
}

func NewPost(author *User, text string) (*Post, error) {
	if len(text) < 5 {
		return nil, fmt.Errorf("Post is too small: %d symbols (min 5)", len(text))
	}

	if len(text) > 500 {
		return nil, fmt.Errorf("Post is too big: %d symbols (max 500)", len(text))
	}

	return &Post{
		Author: author,
		Text:   text,
	}, nil
}

func (p *Post) String() string {
	return fmt.Sprintf("Post[by=%s\n  %s\n]", p.Author, p.Text)
}
