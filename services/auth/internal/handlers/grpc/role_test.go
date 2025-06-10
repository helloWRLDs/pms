package grpchandler

import (
	"context"
	"testing"

	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/utils"
)

func TestServer_ListRoles(t *testing.T) {
	res, err := client.ListRoles(context.Background(), &pb.ListRolesRequest{
		Filter: &dto.RoleFilter{
			// CompanyId:   "dee3b9c8-b6a4-4106-9304-525b3da7dc30",
			WithDefault: true,
			Page:        1,
			PerPage:     10,
		},
	})
	if err != nil {
		t.Fatalf("failed to list roles: %v", err)
	}
	t.Logf("roles: %v", utils.JSON(res))
}
