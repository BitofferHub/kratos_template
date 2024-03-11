package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/log"
	"github.com/bitstormhub/bitstorm/userX/internal/biz"
)

type userXRepo struct {
	data *Data
}

func (r *userXRepo) Delete(ctx context.Context, id int64) error {
	return r.data.DB(ctx).Where("id = ?", id).Delete(&biz.User{}).Error
}

// NewUserRepo
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@param data
//	@return biz.UserRepo
func NewUserXRepo(data *Data) biz.UserRepo {
	return &userXRepo{
		data: data,
	}
}

// Save
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@param g
//	@return *biz.User
//	@return error
func (r *userXRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	// 开启事务的话, 需要调用r.data.DB(ctx) 而不是r.data.db
	err := r.data.DB(ctx).Create(g).Error
	return g, err
}

// Update
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@param g
//	@return *biz.User
//	@return error
func (r *userXRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	err := r.data.db.WithContext(ctx).Updates(g).Error
	return g, err
}

// FindByIDWithCache
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@param userID
//	@return *biz.User
//	@return error
func (r *userXRepo) FindByIDWithCache(ctx context.Context,
	userID int64) (*biz.User, error) {
	cacheKey := fmt.Sprintf("userinfo:%d", userID)
	var user = new(biz.User)
	rdbUserInfo, exist, err := r.data.cache.Get(ctx, cacheKey)
	if err == nil && exist {
		err = json.Unmarshal([]byte(rdbUserInfo), user)
		if err == nil {
			return user, nil
		}
	}
	user, err = r.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	userStr, _ := json.Marshal(user)
	if userStr != nil && len(userStr) != 0 {
		err = r.data.cache.Set(ctx, cacheKey, string(userStr), 10)
		if err != nil {
			log.InfoContextf(ctx, "set user cacheKey err %s", err.Error())
		}
	}
	return user, nil
}

// FindByID
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@param userID
//	@return *biz.User
//	@return error
func (r *userXRepo) FindByID(ctx context.Context, userID int64) (*biz.User, error) {
	var user biz.User
	err := r.data.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByName
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@param userName
//	@return *biz.User
//	@return error
func (r *userXRepo) FindByName(ctx context.Context, userName string) (*biz.User, error) {
	var user biz.User
	err := r.data.db.WithContext(ctx).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ListAll
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver r
//	@param ctx
//	@param data
//	@return []*biz.User
//	@return error
func (r *userXRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}
