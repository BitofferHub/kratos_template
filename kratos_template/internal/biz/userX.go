package biz

import (
	"context"
	"time"
)

// 表定义也需要放在biz层, 还是为了解耦biz与data层
type User struct {
	UserID     int `gorm:"column:id"`
	UserName   string
	Pwd        string
	Sex        int
	Age        int
	Email      string
	Contact    string
	Mobile     string
	IdCard     string
	CreateTime time.Time  `gorm:"column:create_time;default:null"`
	ModifyTime *time.Time `gorm:"column:modify_time;default:null"`
}

// TableName 表名
func (p *User) TableName() string {
	return "t_user_info"
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	FindByName(context.Context, string) (*User, error)
	Delete(context.Context, int64) error
}

// UserXUseCase is a User usecase.
type UserXUseCase struct {
	repo UserRepo
	tm   Transaction
}

// NewUserUseCase new a User usecase.
func NewUserXUseCase(repo UserRepo, tm Transaction) *UserXUseCase {
	return &UserXUseCase{repo: repo, tm: tm}
}

// CreateUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: creates a User, and returns the new User.
//	@Receiver uc
//	@param ctx
//	@param data
//	@param g
//	@return *User
//	@return error
func (uc *UserXUseCase) CreateUser(ctx context.Context, g *User) (*User, error) {
	var user *User
	var err error
	// 开启事务
	uc.tm.InTx(ctx, func(ctx context.Context) error {
		// 故意删除一条id 为4的记录, 用于测试事务回滚
		uc.repo.Delete(ctx, 4)

		user, err = uc.repo.Save(ctx, g)
		if err != nil {
			// 需要回滚就返回error
			return err
		}
		// 返回nil，提交事务
		return nil
	})
	return user, err
}

// GetUserInfo
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: get  User, and returns new User.
//	@Receiver uc
//	@param ctx
//	@param data
//	@param userID
//	@return *User
//	@return error
func (uc *UserXUseCase) GetUserInfo(ctx context.Context, userID int64) (*User, error) {
	return uc.repo.FindByID(ctx, userID)
}

// GetUserInfoByName
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: get  User, and returns new User.
//	@Receiver uc
//	@param ctx
//	@param data
//	@param userName
//	@return *User
//	@return error
func (uc *UserXUseCase) GetUserInfoByName(ctx context.Context, userName string) (*User, error) {
	return uc.repo.FindByName(ctx, userName)
}
