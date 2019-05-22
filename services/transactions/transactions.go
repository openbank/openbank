package transactions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	transactionspb "github.com/openbank/gunk/gunk/v1/transactions"
	"github.com/openbank/openbank/storage"
)

// Handler implements the transactions service.
type Handler struct {
	store storage.TransactionStore
}

// NewHandler returns a new handler for the transactions service.
func NewHandler(store storage.TransactionStore) *Handler {
	if store == nil {
		panic("store cannot be nil")
	}
	return &Handler{store}
}

// RegisterService registers the transactions service to the gRPC server.
func RegisterService(store storage.TransactionStore) func(srv *grpc.Server) error {
	return func(srv *grpc.Server) error {
		h := NewHandler(store)
		transactionspb.RegisterTransactionServiceServer(srv, h)
		return nil
	}
}

// RegisterGateway registers the transactions service to the gRPC gateway.
func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, addr string, opts []grpc.DialOption) error {
	return transactionspb.RegisterTransactionServiceHandlerFromEndpoint(ctx, mux, addr, opts)
}

// GetTransaction retrieves the detail of a transaction, selected by its id.
func (h *Handler) GetTransaction(ctx context.Context, req *transactionspb.GetTransactionRequest) (*transactionspb.Transaction, error) {
	return h.store.GetTransaction(ctx, req.GetTransactionID())
}

// GetTransactions returns a list containing up to 20 transactions.
func (h *Handler) GetTransactions(ctx context.Context, req *transactionspb.GetTransactionsRequest) (*transactionspb.GetTransactionsResponse, error) {
	res, hasMore, err := h.store.GetTransactions(ctx, req.GetNextStartingIndex())
	if err != nil {
		return nil, err
	}
	return &transactionspb.GetTransactionsResponse{
		Result:  res,
		HasMore: hasMore,
	}, nil
}

// CreateTransaction creates a new transaction and returns its id.
func (h *Handler) CreateTransaction(ctx context.Context, req *transactionspb.CreateTransactionRequest) (*transactionspb.CreateTransactionResponse, error) {
	id := uuid.New().String()
	tr := &transactionspb.Transaction{
		TransactionID: id,
		SourceAccount: &transactionspb.BankAccountInfo{
			AccountID: req.GetSourceAccountID(),
		},
		DestinationAccount: &transactionspb.BankAccountInfo{
			AccountID: req.GetDestinationAccountID(),
		},
		Date: time.Now().UTC().String(),
		// TODO: Type
		Status: transactionspb.Status_Pending,
		Amount: req.GetAmount(),
		// TODO: Description
		// TODO: UserID
		Remarks: req.GetRemarks(),
	}
	if err := h.store.CreateTransaction(ctx, tr); err != nil {
		return nil, err
	}
	return &transactionspb.CreateTransactionResponse{
		TransactionID: id,
	}, nil
}
