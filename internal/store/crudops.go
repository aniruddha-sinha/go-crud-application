package store

import "context"

func (s *UserStore) CreateUser(ctx context.Context, name, email string) (int, error) {
	var id int
	query := `INSERT INTO app_data.users (username, email) VALUES ($1, $2) RETURNING id`

	err := s.DB.QueryRowContext(ctx, query, name, email).Scan(&id)
	return id, err
}

func (s *UserStore) GetUser(ctx context.Context, id int) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email FROM app_data.users where id = $1`
	err := s.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
