package dao

import (
	"context"
	"database/sql"
	"fmt"
	utils "projects/chatterbox/server/pkgs/utilities"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64       `sql:"id" json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Role      string      `json:"role"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
	UpdatedAt interface{} `sql:"updated_at" json:"updated_at"`
}

type Home struct {
	Id        int64       `sql:"id" json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Following int         `json:"following"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
}

type ShortUser struct {
	Id        int64       `sql:"id" json:"id"`
	Username  string      `json:"username"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
}

type DAO struct {
	PgClient *sql.DB
	MgClient *mongo.Client
	RdClient *redis.Client
}

func (dao DAO) CheckExistingUser(userData User) (bool, User, error) {
	var (
		isExisting bool
		users      []User
		user       User
	)

	query := fmt.Sprintf("select id, username, password from users where username='%v' limit 1;", userData.Username)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return isExisting, user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return isExisting, user, err
		}
		users = append(users, user)
	}

	if len(users) > 0 {
		isExisting = true
	}

	fmt.Println("isExisting", isExisting)
	return isExisting, user, nil
}

func (dao DAO) AddUserToDatabase(userData User) error {
	pCost, err := bcrypt.Cost([]byte(userData.Password))
	fmt.Println("pCost", pCost, err)

	pHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT into users (username, email, password) VALUES ('%s','%s','%s') ON CONFLICT DO NOTHING;;", userData.Username, userData.Email, pHash)
	return dao.execQuery(query)
}

func (dao DAO) GetAllUsersFromDatabase() ([]User, error) {
	var (
		users []User
		user  User
	)

	query := "select id, username, email, password, created_at from users;"
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		fmt.Println("user:", user)
		users = append(users, user)
	}

	return users, err
}

func (dao DAO) GetAllUserProfilesFromDatabase(follower_id string) ([]Home, error) {
	var (
		user  Home
		users []Home
	)

	query := fmt.Sprintf("select id, username, email, COALESCE(f.user_id, -1) as following, created_at from users u left join (select user_id, follower_id from followers where follower_id=%v) f on u.id=f.user_id where u.id<>%v;", follower_id, follower_id)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Following, &user.CreatedAt); err != nil {
			return nil, err
		}
		fmt.Println("user:", user)
		users = append(users, user)
	}

	return users, err
}

func (dao DAO) execQuery(queryString string) error {
	dao.PgClient.Exec(queryString)
	_, err := dao.PgClient.Exec(queryString)
	if err != nil {
		return err
	}

	return nil
}

func (dao DAO) GetFollowers(userId string) ([]ShortUser, error) {
	var (
		user  ShortUser
		users []ShortUser
	)

	query := fmt.Sprintf("select users.id, users.username, users.created_at from users left join followers on users.id=followers.user_id where user_id=%v;", userId)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.CreatedAt); err != nil {
			return nil, err
		}
		fmt.Println("user:", user)
		users = append(users, user)
	}

	return users, err
}

func (dao DAO) GetFollowing(userId string) ([]ShortUser, error) {
	var (
		user  ShortUser
		users []ShortUser
	)

	query := fmt.Sprintf("select users.id, users.username, users.created_at from users left join followers on users.id=followers.user_id where followers.follower_id=%v;", userId)
	rows, err := dao.PgClient.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows, err)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.CreatedAt); err != nil {
			return nil, err
		}
		fmt.Println("user:", user)
		users = append(users, user)
	}

	return users, err
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
