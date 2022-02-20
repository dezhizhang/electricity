package service

import (
	"context"
	"electricity/driver"
	"electricity/model"
	"electricity/proto"
	"electricity/utils"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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


func (u *UserServer) GetUserId(ctx context.Context,req *proto.IdRequest) (*proto.UserInfoResponse,error)  {
	// 通过id获取用户信息
	var user model.User
	result := driver.DB.First(&user,req.Id)
	if result.RowsAffected == 0 {
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}

	if result.Error != nil {
		return nil,result.Error
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp,nil
}

func (u *UserServer) CreateUser(ctx context.Context,req *proto.CreateUserInfo) (*proto.UserInfoResponse,error){
	//新建用户
	var user model.User
	result := driver.DB.Where(&model.User{Phone: req.Phone}).First(&user)
	if result.RowsAffected == 1 {
		return nil,status.Errorf(codes.AlreadyExists,"用户已存在")
	}

	user.Phone = req.Phone
	user.NickName = req.NickName
	user.Password = utils.Md5String(req.Password)

	// 插入数据
	result = driver.DB.Create(&user)
	if result.Error != nil {
		return nil,status.Errorf(codes.Internal,result.Error.Error())
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp,nil
}

func (u *UserServer) UpdateUser(ctx context.Context,req *proto.UpdateUserInfo) (*empty.Empty,error)  {
	// 更新用户
	var user model.User
	result := driver.DB.First(&user,req.Id)
	if result.RowsAffected == 0 {
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}

	birthDay := time.Unix(int64(req.Birthday),0)
	user.NickName = req.NickName
	user.Gender = req.Gender
	user.Birthday = &birthDay

	result = driver.DB.Save(user)

	if result.Error != nil {
		return nil,status.Errorf(codes.Internal,result.Error.Error())
	}
	return &empty.Empty{},nil
}


