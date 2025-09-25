package inmemory

import (
	"toDoList/internal/domain/user/user_errors"
	"toDoList/internal/domain/user/user_models"
)

// TODO: протестить локальное хранилище

func (storage *Storage) GetAllUsers() ([]user_models.User, error) {
	var users []user_models.User

	for _, user := range storage.users {
		users = append(users, user)
	}
	return users, nil
}

func (storage *Storage) SaveUser(user user_models.User) (user_models.User, error) {
	for _, userInMemory := range storage.users {
		if user.Email == userInMemory.Email {
			return user_models.User{}, user_errors.ErrorUserIsAlreadyExist
		}
	}

	storage.users[user.Uuid] = user

	return user, nil
}

func (storage *Storage) GetUserByID(userID string) (user_models.User, error) {
	user, ok := storage.users[userID]
	if !ok {
		return user_models.User{}, user_errors.ErrorUserNotExist
	}

	return user, nil
}

func (storage *Storage) GetUserByEmail(email string) (user_models.User, error) {
	var user user_models.User
	for _, userInMemory := range storage.users {
		if userInMemory.Email == email {
			user = userInMemory
			return user, nil
		}
	}
	return user_models.User{}, user_errors.ErrorUserNotExist
}

func (storage *Storage) UpdateUser(user user_models.User) (user_models.User, error) {
	for _, userInMemory := range storage.users {
		if userInMemory.Uuid == user.Uuid {
			userInMemory.Name = user.Name
			userInMemory.Email = user.Email
			userInMemory.Password = user.Password

			storage.users[user.Uuid] = userInMemory
			return user, nil
		}
	}
	return user_models.User{}, user_errors.ErrorUserNotExist
}

func (storage *Storage) DeleteUser(userID string) error {
	_, ok := storage.users[userID]
	if !ok {
		return user_errors.ErrorUserNotExist
	}
	delete(storage.users, userID)
	return nil
}
