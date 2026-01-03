// cmd/repositories/userDb.go
package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func GetUsers() ([]models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `
    UPDATE users
    SET name = $2, email = $3, password = $4, updated_at = $5
    WHERE id = $1
    RETURNING id`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = id
	return user, nil
}
