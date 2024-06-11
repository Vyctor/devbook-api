package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repo Posts) CreatePost(post models.Post) (uint64, error) {
	stmt, err := repo.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(post.Title, post.Body, post.AuthorId)
	if err != nil {
		return 0, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertedId), nil
}

func (repo Posts) FindPost(id uint64) (models.Post, error) {
	stmt, err := repo.db.Query("select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?", id)
	if err != nil {
		return models.Post{}, err
	}
	defer stmt.Close()
	var post models.Post
	if stmt.Next() {
		if err = stmt.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorId, &post.AuthorNick); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

func (repo Posts) FindPosts(userId uint64) ([]models.Post, error) {
	rows, err := repo.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorId, &post.AuthorNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo Posts) UpdatePost(post models.Post) error {
	stmt, err := repo.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(post.Title, post.Body, post.ID); err != nil {
		return err
	}
	return nil
}

func (repo Posts) DeletePost(postId uint64) error {
	stmt, err := repo.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(postId); err != nil {
		return err
	}
	return nil
}

func (repo Posts) FindPostsByUserId(userId uint64) ([]models.Post, error) {
	rows, err := repo.db.Query(`
		select p.*, u.nick from publicacoes p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorId, &post.AuthorNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo Posts) LikePost(postId uint64) error {
	stmt, err := repo.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(postId); err != nil {
		return err
	}
	return nil
}

func (repo Posts) Unlike(postId uint64) error {
	stmt, err := repo.db.Prepare("update publicacoes set curtidas = curtidas - 1 where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(postId); err != nil {
		return err
	}
	return nil
}
