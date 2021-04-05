package sqlstore

import (
	"database/sql"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *model.User) (int, error) {

	createUserSql := `INSERT INTO "user" (
		first_name,
		middle_name,
		last_name,
		mobile,
		email,
		password,
		admin,
		vendor,
		intro,
		profile)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`
	err := r.store.db.QueryRow(createUserSql,
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Mobile,
		u.Email,
		u.PasswordHash,
		u.Admin,
		u.Vendor,
		u.Intro,
		u.Profile,
	).Scan(&u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.store.db.QueryRow(`SELECT max(id) FROM "user"`).Scan(&u.ID)
		} else {
			logrus.Error(err.Error())
		}
		if err != nil {
			return 0, err
		}
	}

	return u.ID, nil
}

func (r *UserRepo) GetUser(userID int) (*model.User, error) {

	querySql := `SELECT 
		id, 
		first_name,
		middle_name,
		last_name,
		mobile,
		email,
		password,
		admin,
		vendor,
		intro,
		profile,
		registered_at,
		last_login
	FROM "user" 
	WHERE id=$1`

	row := r.store.db.QueryRow(querySql, userID)

	u := &model.User{}
	var lastLoginNullable sql.NullTime
	err := row.Scan(&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Mobile,
		&u.Email,
		&u.PasswordHash,
		&u.Admin,
		&u.Vendor,
		&u.Intro,
		&u.Profile,
		&u.RegisteredAt,
		&lastLoginNullable,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if lastLoginNullable.Valid {
		u.LastLogin = lastLoginNullable.Time
	}

	return u, nil
}

func (r *UserRepo) DeleteUser(userID int) error {

	deleteUserSql := `DELETE FROM "user" WHERE id=$1`

	_, err := r.store.db.Exec(deleteUserSql, userID)
	if err != nil {
		return err
	}

	return nil
}
