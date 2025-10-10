package user

import (
	"context"
	"main/internal/entity"
	"main/internal/services/user"
	"time"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data user.Create) (entity.User, error) {
	var detail entity.User

	detail.FullName = data.FullName
	detail.Email = data.Email
	detail.Password = &data.Password
	detail.RegionId = data.RegionId
	detail.DistrictId = data.DistrictId
	detail.Avatar = data.Avatar
	detail.CreatedBy = data.CreatedBy

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var detail entity.User
	err := r.NewSelect().Model(&detail).Where("email = ?", email).Scan(ctx)
	return detail, err
}

func (r Repository) Update(ctx context.Context, data user.Update, updated_by int) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", *data.Id).Scan(ctx)
	if data.FullName != nil {
		detail.FullName = *data.FullName
	}
	if data.Password != nil {
		detail.Password = data.Password
	}
	if data.Email != nil {
		detail.Email = *data.Email
	}
	if data.RegionId != nil {
		detail.RegionId = *data.RegionId
	}
	if data.DistrictId != nil {
		detail.DistrictId = *data.DistrictId
	}
	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}
	if data.Status != nil {
		detail.Status = *data.Status
	}
	if data.Role != nil {
		detail.Role = *data.Role
	}

	detail.UpdatedBy = updated_by
	detail.UpdatedAt = time.Now()
	_, err = r.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)

	return detail, err
}

func (r Repository) UpdatePassword(ctx context.Context, email, password string) error {
	_, err := r.NewUpdate().
		Model(&entity.User{}).
		Set("password = ?", password).
		Where("email = ?", email).
		Exec(ctx)
	return err
}

func (r Repository) GetAll(ctx context.Context, filter user.Filter, order string) ([]entity.User, int, error) {
	var list []entity.User

	q := r.NewSelect().Model(&list)

	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}

	if filter.Status != nil {
		q.WhereGroup(" and ", func(query *bun.SelectQuery) *bun.SelectQuery {
			query.Where("status = ?", *filter.Status)
			return query
		})
	}

	if order == "asc" {
		q.Order("id asc")
	} else {
		q.Order("id desc")
	}
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {
	_, err := r.NewDelete().Table("users").Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) GetByEmailWithLocation(ctx context.Context, id int, lang string) (user.UserWithLocation, error) {
	var result user.UserWithLocation
	var u entity.User
	err := r.NewSelect().Model(&u).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return result, err
	}

	result.FullName = u.FullName
	result.Email = u.Email
	result.Avatar = u.Avatar
	result.Role = u.Role

	var regionName string
	err = r.NewSelect().
		Table("regions").
		ColumnExpr("name ->> ?", lang).
		Where("id = ?", u.RegionId).
		Scan(ctx, &regionName)
	if err == nil {
		result.RegionName = regionName
	}

	var districtName string
	err = r.NewSelect().
		Table("districts").
		ColumnExpr("name ->> ?", lang).
		Where("id = ?", u.DistrictId).
		Scan(ctx, &districtName)
	if err == nil {
		result.DistrictName = districtName
	}
	return result, nil
}

func (r *Repository) UpdateCabinet(ctx context.Context, data user.UpdateCabinet, id int) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return entity.User{}, err
	}

	if data.Email != nil {
		detail.Email = *data.Email
	}

	if data.RegionId != nil {
		detail.RegionId = *data.RegionId
	}

	if data.DistrictId != nil {
		detail.DistrictId = *data.DistrictId
	}

	if data.Fullname != nil {
		detail.FullName = *data.Fullname
	}

	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}

	detail.UpdatedAt = time.Now()
	detail.UpdatedBy = id

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return entity.User{}, err
	}

	detail.Password = nil

	return detail, nil
}
