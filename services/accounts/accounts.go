package accounts

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	accountspb "github.com/openbank/gunk/gunk/v1/accounts"
	apierrorpb "github.com/openbank/gunk/gunk/v1/apierror"
	"github.com/openbank/openbank/pkg/apierror"
	"github.com/openbank/openbank/storage"
)

// Handler implements the accounts service.
type Handler struct {
	store storage.AccountStore
}

// NewHandler returns a new handler for the accounts service.
func NewHandler(store storage.AccountStore) *Handler {
	if store == nil {
		panic("store cannot be nil")
	}
	return &Handler{store}
}

// RegisterService registers the accounts service to the gRPC server.
func RegisterService(store storage.AccountStore) func(srv *grpc.Server) error {
	return func(srv *grpc.Server) error {
		h := NewHandler(store)
		accountspb.RegisterAccountServiceServer(srv, h)
		return nil
	}
}

// RegisterGateway registers the accounts service to the gRPC gateway.
func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, addr string, opts []grpc.DialOption) error {
	return accountspb.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, addr, opts)
}

// GetAccount retrieves the detail of an account, selected by its id.
func (h *Handler) GetAccount(ctx context.Context, req *accountspb.GetAccountRequest) (*accountspb.Account, error) {
	a, err := h.store.GetAccount(ctx, req.AccountID)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, apierror.NotFound("account %s not found", req.AccountID)
		}
		return nil, apierror.Internal(err)
	}
	return a, nil
}

// GetAccounts returns a list containing up to 20 accounts.
func (h *Handler) GetAccounts(ctx context.Context, req *accountspb.GetAccountsRequest) (*accountspb.GetAccountsResponse, error) {
	res, hasMore, err := h.store.GetAccounts(ctx, req.NextStartingIndex)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return &accountspb.GetAccountsResponse{
		Result:  res,
		HasMore: hasMore,
	}, nil
}

// CreateAccount creates a new account and return its id.
func (h *Handler) CreateAccount(ctx context.Context, req *accountspb.CreateAccountRequest) (*accountspb.CreateAccountResponse, error) {
	id := req.AccountID
	if id == "" {
		id = uuid.New().String()
	}
	// TODO: what to do with CreateAccountRequest fields
	// that can not be mapped to Account?
	err := h.store.CreateAccount(ctx, &accountspb.Account{
		AccountID: id,
		// Description:      req.Description,
		// AccountRoles:     req.AccountRoles,
		Branch: req.Branch,
		// Customer:         req.Customer,
		// DebitTransaction: req.DebitTransaction,
		InterestRate: req.InterestRate,
		MajorType:    req.MajorType,
		MaturityDate: req.MaturityDate,
		// Minor:            req.Minor,
	})
	if err != nil {
		if err == storage.ErrConflict {
			return nil, apierror.Conflict(apierrorpb.ErrorType_AccountAlreadyExists, "account %s already exists", req.AccountID)
		}
		return nil, apierror.Internal(err)
	}
	return &accountspb.CreateAccountResponse{
		AccountID: id,
	}, nil
}

// UpdateAccount updates an account.
func (h *Handler) UpdateAccount(ctx context.Context, req *accountspb.UpdateAccountRequest) (*empty.Empty, error) {
	// TODO: what is to update ?
	err := h.store.UpdateAccount(ctx, &accountspb.Account{
		AccountID: req.AccountID,
		// Description: req.Description,
	})
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, apierror.NotFound("account %s not found", req.AccountID)
		}
		return nil, apierror.Internal(err)
	}
	return &empty.Empty{}, nil
}

// DeleteAccount deletes an account.
func (h *Handler) DeleteAccount(ctx context.Context, req *accountspb.DeleteAccountRequest) (*empty.Empty, error) {
	err := h.store.DeleteAccount(ctx, req.AccountID)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, apierror.NotFound("account %s not found", req.AccountID)
		}
		return nil, apierror.Internal(err)
	}
	return &empty.Empty{}, nil
}
