package service

import (
	"context"
	"electricity/driver"
	"electricity/model"
	"electricity/proto"
	"electricity/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResponse := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}

	if user.Birthday != nil {
		userInfoResponse.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoResponse
}

// 获取用户列表

func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	// 获取用户列表
	var users []model.User
	result := driver.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	driver.DB.Scopes(utils.Paginate(int(req.Page), int(req.PageSize))).Find(&users)

	// 添加用户
	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

// 获取用户手机号

func (u *UserServer) GetUserPhone(ctx context.Context, req *proto.PhoneRequest) (*proto.UserInfoResponse, error) {
	// 通过手机号查询用户
	var user model.User

	result := driver.DB.Where(&model.User{Phone: req.Phone}).First(&user)

	if result.Error != nil {
		return nil,result.Error
	}

	if result.RowsAffected == 0 {
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp,nil
}