package dao

import (
	"context"
	"database/sql"
	"fmt"
	utils "projects/chatterbox/server/pkgs/utilities"
	"strings"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    int64       `sql:"user_id" json:"user_id"`
	Username  string      `json:"username" validate:"required"`
	Email     string      `json:"email" validate:"required"`
	Password  string      `json:"password" validate:"required"`
	Role      string      `json:"role" validate:"required"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
	UpdatedAt interface{} `sql:"updated_at" json:"updated_at"`
}

type Profile struct {
	UserId    int64       `sql:"user_id" json:"user_id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Following int         `json:"following"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
}

type DAO struct {
	PgClient *sql.DB
	MgClient *mongo.Client
	RdClient *redis.Client
}

func (dao DAO) CheckExistingUser(userData User) (isExisting bool, user User, err error) {
	query := fmt.Sprintf("select user_id, username, role, password from chatterbox.user where username='%v' limit 1;", userData.Username)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.UserId, &user.Username, &user.Role, &user.Password); err != nil {
			return
		} else {
			isExisting = true
			return
		}
	}

	return
}

func (dao DAO) AddUserToDatabase(userData User) error {
	// pCost, err := bcrypt.Cost([]byte(userData.Password))
	// fmt.Println("pCost", pCost, err)

	if passwordHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		query := fmt.Sprintf("insert into chatterbox.user (username, email, password) VALUES ('%s','%s','%s');", userData.Username, userData.Email, passwordHash)
		if err = dao.execQuery(query); err != nil {
			if strings.Contains(err.Error(), "user_email_key") {
				return fmt.Errorf("email already in use")
			} else {
				return err
			}
		}
		return nil
	}
}

func (dao DAO) GetAllUsersFromDatabase() (users []User, err error) {
	query := "select user_id, username, email, password, role, created_at from chatterbox.user;"
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func (dao DAO) GetAllUserProfilesFromDatabase(auth_user_id string) (profiles []Profile, err error) {
	query := fmt.Sprintf("select u.user_id, u.username, u.email, COALESCE(f.user_id, 0) as following, u.created_at from chatterbox.user u left join (select user_id, follower_id from chatterbox.follower where follower_id=%v) f on u.user_id=f.user_id where u.user_id<>%v;", auth_user_id, auth_user_id)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		var profile Profile
		if err = rows.Scan(&profile.UserId, &profile.Username, &profile.Email, &profile.Following, &profile.CreatedAt); err != nil {
			return
		}
		profiles = append(profiles, profile)
	}

	return
}

func (dao DAO) execQuery(queryString string) error {
	dao.PgClient.Exec(queryString)
	_, err := dao.PgClient.Exec(queryString)
	if err != nil {
		return err
	}

	return nil
}

func (dao DAO) GetFollowers(userId string) (users []User, err error) {
	query := fmt.Sprintf("select u.user_id, user.username, u.created_at from chatterbox.user u left join follower f on u.user_id=f.user_id where u.user_id=%v;", userId)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.UserId, &user.Username, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func (dao DAO) GetFollowing(userId string) (users []User, err error) {
	query := fmt.Sprintf("select u.user_id, u.username, u.created_at from chatterbox.user u left join chatterbox.follower f on u.user_id=f.user_id where f.follower_id=%v;", userId)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.UserId, &user.Username, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func (dao DAO) FollowUser(followingId, followerId string) error {
	followQuery := fmt.Sprintf("insert into followers (user_id, follower_id) values (%v,%v) ON CONFLICT DO NOTHING;", followingId, followerId)
	return dao.execQuery(followQuery)
}

func (dao DAO) UnfollowUser(followingId, followerId string) error {
	unfollowQuery := fmt.Sprintf("delete from followers where user_id=%v and follower_id=%v;", followingId, followerId)
	return dao.execQuery(unfollowQuery)
}

func (dao DAO) SetRedisValue(ctx context.Context, key string, value interface{}) error {
	rdResp := dao.RdClient.Set(ctx, key, value, utils.RedisChatExpiry)
	fmt.Println(rdResp)
	return rdResp.Err()
}

func (dao DAO) GetRedisValue(ctx context.Context, key string) error {
	rdResp := dao.RdClient.Get(ctx, key)
	fmt.Println(rdResp)
	return rdResp.Err()
}

func (dao DAO) UnsetRedisValue(ctx context.Context, key string, value interface{}) error {
	rdResp := dao.RdClient.Set(ctx, key, value, 1)
	fmt.Println(rdResp)
	return rdResp.Err()
}
