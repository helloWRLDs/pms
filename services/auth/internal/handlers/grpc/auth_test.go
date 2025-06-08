package grpchandler

import (
	"context"
	"testing"

	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/utils"
)

func TestServer_LoginUser(t *testing.T) {
	res, err := client.LoginUser(context.Background(), &pb.LoginUserRequest{
		Credentials: &dto.UserCredentials{
			Email:    "IvanIvanov@gmail.com",
			Password: "IvanIvanov1",
		},
	})

	if err != nil {
		t.Fatalf("failed to login user: %v", err)
	}

	t.Logf("response: %v", utils.JSON(res))
	t.Logf("permissions: %v", utils.JSON(res.Payload.User.Permissions))
}
