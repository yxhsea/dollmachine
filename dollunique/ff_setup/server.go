package ff_setup

import (
	"github.com/micro/go-micro"
	proto "dollmachine/dollunique/proto"
	"context"
	"dollmachine/dollunique/ff_cache/ff_unique_id"
	"github.com/micro/go-micro/registry/consul"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type UniqueId struct {
}

func (u *UniqueId) GenerateUniqueId(context context.Context, req *proto.UniqueIdRequest, rsp *proto.UniqueIdResponse) error {
	UniqueIdDao := ff_unique_id.NewUniqueId()
	valInt64, err := UniqueIdDao.GetUniqueId(req.Key)
	if err != nil {
		logrus.Errorf("GenerateUniqueId failure for %v. Error : %v", req.Key, err)
		return err
	}
	rsp.Value = valInt64
	return nil
}

func SetupServer(host string) {
	service := micro.NewService(
		micro.Name("go.micro.srv.unique_id"),
		micro.Version("1.0"),
		micro.Registry(consul.NewRegistry(consul.Config(&api.Config{Address:host}))),
	)

	service.Init()
	proto.RegisterGenerateUniqueIdHandler(service.Server(), new(UniqueId))

	if err := service.Run(); err != nil {
		logrus.Fatalf("service init failure. Error : %v",err)
	}
}
