package repository

import (
	"giapps/servisin/domain/entity"
	"giapps/servisin/infrastructure/database"
	"giapps/servisin/infrastructure/exception"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	FindUserByUsername(username string) (user entity.UserEntity, err error)
	FindUserByUsernameOrEmail(username string, email string) (user entity.UserEntity, err error)
	Insert(user entity.UserEntity)
}

func NewUserRepository(database *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		db: database,
	}
}

type userRepositoryImpl struct {
	db *pgxpool.Pool
}

func (repo *userRepositoryImpl) FindUserByUsernameOrEmail(username string, email string) (user entity.UserEntity, err error) {
	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, "SELECT user_id,username,email, password FROM auth.users WHERE username = $1 or email = $2", username, email)
	exception.PanicIfNeeded(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password)
		exception.PanicIfNeeded(err)
		return user, nil
	} else {
		return user, exception.ErrResponse{
			Code:    http.StatusUnauthorized,
			Message: "pengguna tidak ditemukan",
		}
	}
}

func (repo *userRepositoryImpl) FindUserByUsername(username string) (user entity.UserEntity, err error) {
	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	rows, err := repo.db.Query(ctx, "SELECT user_id,username,email, password FROM auth.users WHERE username = $1", username)
	exception.PanicIfNeeded(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password)
		exception.PanicIfNeeded(err)
		return user, nil
	} else {
		return user, exception.ErrResponse{
			Code:    http.StatusUnauthorized,
			Message: "pengguna tidak ditemukan",
		}
	}
}

func (repo *userRepositoryImpl) Insert(user entity.UserEntity) {
	ctx, cancel := database.NewDatabaseContext()
	defer cancel()

	row, err := repo.db.Exec(ctx, "INSERT INTO auth.users (user_id, username, email, password) VALUES ($1, $2, $3, $4)", user.UserId, user.Username, user.Email, user.Password)
	exception.PanicIfNeeded(err)
	row.Insert()
}
