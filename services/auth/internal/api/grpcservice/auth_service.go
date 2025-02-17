package grpcservice

import (
	"context"

	pb "github.com/nidea1/go-gavel/proto/auth"
	"github.com/nidea1/go-gavel/services/auth/internal/usecase/command"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	RegisterUC *command.RegisterUseCase
}

func NewAuthService(registerUC *command.RegisterUseCase) *AuthService {
	return &AuthService{RegisterUC: registerUC}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	email := req.GetEmail()
	password := req.GetPassword()

	status, message, valErrs, err := s.RegisterUC.Register(ctx, firstName, lastName, email, password)
	if err != nil {
		return &pb.RegisterResponse{
			Status:  status,
			Message: message,
			Errors:  nil, // Remove err.Error() here since Errors should be of type *pb.Errors
		}, err
	}

	var pbErrors *pb.Errors = nil
	if valErrs != nil {
		var pbValErrors []*pb.SingleValidationError
		for _, valErr := range valErrs.Errors {
			pbValErrors = append(pbValErrors, &pb.SingleValidationError{
				Field:   valErr.Field,
				Code:    valErr.Code,
				Message: valErr.Message,
			})
		}
		pbErrors = &pb.Errors{Errors: pbValErrors}
	}

	return &pb.RegisterResponse{
		Status:  status,
		Message: message,
		Errors:  pbErrors,
	}, nil
}
