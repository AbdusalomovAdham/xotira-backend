package user

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/services/user"
	"mime/multipart"
	"time"
)

type UseCase struct {
	user User
	auth Auth
	file File
}

func NewUseCase(user User, auth Auth, file File) *UseCase {
	return &UseCase{user: user, auth: auth, file: file}
}

func (uc UseCase) AdminCreateUser(ctx context.Context, data user.Create, authHeader string) (entity.User, error) {

	tokenDetail, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.User{}, err
	}

	hashPassword, err := uc.auth.HashPassword(data.Password)
	if err != nil {
		return entity.User{}, err
	}

	emailCheck, err := uc.user.GetByEmail(ctx, data.Email)
	if err == nil && emailCheck.Id != 0 {
		return entity.User{}, errors.New("email already exists")
	}
	data.Password = hashPassword
	data.CreatedBy = tokenDetail.Id

	detail, err := uc.user.Create(ctx, data)
	if err != nil {
		return entity.User{}, err
	}

	detail.Password = nil
	return detail, nil
}

func (uc UseCase) AdminGetUserList(ctx context.Context, filter user.Filter, order string) ([]entity.User, int, error) {
	return uc.user.GetAll(ctx, filter, order)
}

func (uc UseCase) AdminGetUserDetail(ctx context.Context, id int) (entity.User, error) {
	detail, err := uc.user.GetById(ctx, id)
	if err != nil {
		return entity.User{}, errors.New("user not found")
	}
	return detail, nil
}

func (uc UseCase) AdminDeleteUser(ctx context.Context, id int) error {
	if err := uc.user.Delete(ctx, id); err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (uc UseCase) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	return uc.file.Upload(ctx, file, folder)
}
func (uc UseCase) GetByEmail(ctx context.Context, authHeader string) (entity.User, error) {
	tokenDetail, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.User{}, err
	}
	detail, err := uc.user.GetByEmail(ctx, tokenDetail.Email)
	if err != nil {
		return entity.User{}, err
	}

	return detail, nil
}

func (uc UseCase) GetByEmailWithLocation(ctx context.Context, authHeader, lang string) (user.UserWithLocation, error) {
	tokenStr, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return user.UserWithLocation{}, err
	}
	return uc.user.GetByEmailWithLocation(ctx, tokenStr.Id, lang)
}

func (uc UseCase) AdminUpdateUser(ctx context.Context, data user.Update, avatarUpdated bool, authHeader string) (entity.User, error) {
	oldUser, err := uc.user.GetById(ctx, *data.Id)
	if err != nil {
		return entity.User{}, err
	}

	tokenDetail, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.User{}, err
	}

	if avatarUpdated && oldUser.Avatar != "" {
		if err := uc.file.Delete(ctx, oldUser.Avatar); err != nil {
			return entity.User{}, err
		}
	}

	if data.Password != nil {
		hashedPassword, err := uc.auth.HashPassword(*data.Password)
		if err != nil {
			return entity.User{}, err
		}
		data.Password = &hashedPassword
	}

	data.UpdatedAt = time.Now()

	detail, err := uc.user.Update(ctx, data, tokenDetail.Id)

	if err != nil {
		return entity.User{}, err
	}

	detail.Password = nil

	return detail, nil
}

func (uc UseCase) UpdateCabinet(ctx context.Context, data user.UpdateCabinet, authHeader string, avatarUpdated bool) (entity.User, error) {

	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.User{}, err
	}

	oldUser, err := uc.user.GetById(ctx, token.Id)
	if err != nil {
		return entity.User{}, err
	}

	if avatarUpdated && oldUser.Avatar != "" {
		if err := uc.file.Delete(ctx, oldUser.Avatar); err != nil {
			return entity.User{}, err
		}
	}

	detail, err := uc.user.UpdateCabinet(ctx, data, token.Id)

	return detail, err
}
