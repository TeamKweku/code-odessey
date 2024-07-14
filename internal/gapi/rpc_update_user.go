package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/pb"
	"github.com/teamkweku/code-odessey/pkg/utils"
	"github.com/teamkweku/code-odessey/pkg/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authUser, err := server.store.GetUserByUsername(ctx, authPayload.Username)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.UpdateUserParams{
		ID: authUser.ID,
	}

	if req.Username != nil {
		arg.Username = pgtype.Text{
			String: req.GetUsername(),
			Valid:  true,
		}
	}
	if req.FullName != nil {
		arg.FullName = pgtype.Text{
			String: req.GetFullName(),
			Valid:  true,
		}
	}
	if req.Email != nil {
		arg.Email = pgtype.Text{
			String: req.GetEmail(),
			Valid:  true,
		}
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashedPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Username != nil {
		if err := val.ValidateUsername(req.GetUsername()); err != nil {
			violations = append(violations, fieldViolation("username", err))
		}
	}

	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}

	}

	if req.FullName != nil {
		if err := val.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}
