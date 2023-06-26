package database

import (
	"database/sql"
	"mark-server/internal/domain"
	"time"
)

type UserRepository interface {
	GetUsers(page, limit int) ([]*domain.User, error)
	GetTotalUsers() (int, error)
	UpdateUser(user *domain.User, userID int) error
	DeleteUser(userID int) error
	CreateUser(*domain.User) (*domain.User, error)
}

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

func (r *MySQLUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO t_user (name, age, gender, phone, create_time) VALUES (?,?,?,?,?)"

	result, err := r.db.Exec(query, user.Name, user.Age, user.Gender, user.Phone, time.Now().UTC())
	if err != nil {
		return &domain.User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return &domain.User{}, err
	}

	user.ID = int(userID)
	return user, nil
}

func (r *MySQLUserRepository) UpdateUser(user *domain.User, userID int) error {
	query := "UPDATE t_user SET name = ?, age = ?, gender = ?, phone = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Age, user.Gender, user.Phone, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLUserRepository) DeleteUser(userID int) error {
	query := "UPDATE t_user SET is_deleted = 1 WHERE id = ?"
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLUserRepository) GetTotalUsers() (int, error) {
	query := "SELECT COUNT(*) FROM t_user WHERE is_deleted = 0"
	var total int
	if err := r.db.QueryRow(query).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MySQLUserRepository) GetUsers(offset int, size int) ([]*domain.User, error) {
	query := "SELECT id, name, create_time, gender, phone, age FROM t_user WHERE is_deleted = 0 LIMIT ?, ?"
	rows, err := r.db.Query(query, offset, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.CreateTime, &user.Gender, &user.Phone, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
