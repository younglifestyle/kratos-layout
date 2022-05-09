package service

import (
	"context"
	"fmt"
	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// A function that maps field mask field names to the names used in Go structs.
// It has to be implemented according to your needs.
func naming(s string) string {
	return s
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (res *v1.HelloReply, err error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	fmt.Println("mask123 :", in.FieldMask.String())

	v := &v1.HelloReply{Message: "Hello " + in.GetName(), Me1: "123"}

	if in.Ones != nil {
		mask, _ := fieldmask_utils.MaskFromProtoFieldMask(in.FieldMask, naming)
		userDst := make(map[string]interface{})
		err := fieldmask_utils.StructToMap(mask, v, userDst)
		fmt.Printf("dadf: %v %+v \n", err, userDst)
	}
	//in.FieldMask.IsValid()
	//in.FieldMask.Normalize()

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}

	return v, nil
}
