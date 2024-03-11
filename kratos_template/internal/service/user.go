package service

import (
	"context"
	pb "github.com/bitstormhub/bitstorm/userX/api/userX/v1"
	"github.com/bitstormhub/bitstorm/userX/internal/biz"
)

// CreateUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.CreateUserReply
//	@return error
func (s *UserXService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, err := s.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Message: "trytest"}, nil
}
func (s *UserXService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

// DeleteUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.DeleteUserReply
//	@return error
func (s *UserXService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}

// GetUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.GetUserReply
//	@return error
func (s *UserXService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	userInfo, err := s.uc.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// GetUserByName
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.GetUserByNameReply
//	@return error
func (s *UserXService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameReply, error) {
	userInfo, err := s.uc.GetUserInfoByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByNameReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			Id:       int64(userInfo.UserID),
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}

// ListUser
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@Receiver s
//	@param ctx
//	@param req
//	@return *pb.ListUserReply
//	@return error
func (s *UserXService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

// 执行定时任务
func (u *UserXService) Cronjob(ctx context.Context, req *pb.CreateUserRequest) {
	u.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
}
