// internal/app/user/user_service.go

package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"mark-server/internal/domain"
	"mark-server/internal/helpers"
	"mark-server/internal/infrastructure/database"
	_ "mark-server/internal/infrastructure/database"
	"time"
)

type UserService struct {
	repo database.UserRepository
}

func NewUserService(repo database.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers(page int, size int) ([]*domain.UserList, int, error) {
	// Calculate the offset based on the page and limit
	offset := (page - 1) * size

	// Call the repository to retrieve users with pagination
	records, err := s.repo.GetUsers(offset, size)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.GetTotalUsers()
	if err != nil {
		return nil, 0, err
	}

	users := make([]*domain.UserList, 0)
	for _, record := range records {
		user := &domain.UserList{
			ID:         record.ID,
			Name:       record.Name,
			Age:        int(record.Age.Int32),
			Gender:     int(record.Gender.Int32),
			Phone:      record.Phone.String,
			CreateTime: time.Unix(timestamppb.New(record.CreateTime.Time).GetSeconds(), 0).Format("2006-01-02 15:04"),
		}

		users = append(users, user)
	}
	return users, total, nil
}

func (s *UserService) CreateUser(user *domain.UpdateUser) (*domain.CreateUserResp, error) {
	resp := &domain.CreateUserResp{}
	_user := &domain.User{
		Name:   user.Name,
		Age:    helpers.NewNullInt32(user.Age),
		Gender: helpers.NewNullInt32(user.Gender),
		Phone:  helpers.NewNullString(user.Phone),
	}
	createdUser, err := s.repo.CreateUser(_user)
	if err != nil {
		return resp, err
	}

	resp.Id = createdUser.ID
	return resp, nil
}

func (s *UserService) UpdateUser(user *domain.UpdateUser, userID int) error {
	_user := &domain.User{
		Name:   user.Name,
		Age:    helpers.NewNullInt32(user.Age),
		Gender: helpers.NewNullInt32(user.Gender),
		Phone:  helpers.NewNullString(user.Phone),
	}
	err := s.repo.UpdateUser(_user, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(userID int) error {
	err := s.repo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}
