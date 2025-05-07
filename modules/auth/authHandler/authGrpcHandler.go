package authhandler

import (
	"context"

	authPb "github.com/naphat-sirisubkulchai/shop/modules/auth/authPb"
	authusecase "github.com/naphat-sirisubkulchai/shop/modules/auth/authUsecase"
)
type(
	authGrpcHandler struct{
		authUsecase authusecase.AuthUsecaseService
		authPb.UnimplementedAuthGrpcServiceServer
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{
		authUsecase:authUsecase}
}

func (g *authGrpcHandler) CredeatialSearch(ctx context.Context,req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error){
	return g.authUsecase.AccessTokenSearch(ctx, req.AccessToken)
}
func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return g.authUsecase.RolesCount(ctx)
}