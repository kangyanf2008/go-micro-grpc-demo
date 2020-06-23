package grpc_server

import (
	"context"
	"go-micro-grpc-demo/protobuffer_def"
)

type TestGrpcServiceImpl struct {

}
func (s TestGrpcServiceImpl) BaseInterface(context context.Context, request *protobuffer_def.BaseRequest, response*protobuffer_def.BaseResponse) error {
	//fmt.Println(request)
	response.C = request.GetC()
	response.RequestId = request.GetRequestId()
	response.Body = request.GetBody()
	return nil
}
