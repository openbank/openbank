package transactions

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	transactionspb "github.com/openbank/gunk/gunk/v1/transactions"
	storage "github.com/openbank/openbank/storage/mock_storage"
)

func TestHandler_GetTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := storage.NewMockTransactionStore(ctrl)

	type args struct {
		req *transactionspb.GetTransactionRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *storage.MockTransactionStore)
		want      *transactionspb.Transaction
		wantErr   bool
	}{
		{
			name: "Success",
			args: args{
				req: &transactionspb.GetTransactionRequest{
					TransactionID: "1",
				},
			},
			mockCalls: func(m *storage.MockTransactionStore) {
				m.
					EXPECT().
					GetTransaction(gomock.Any(), gomock.Any()).
					Return(&transactionspb.Transaction{
						TransactionID: "1",
					}, nil)
			},
			want: &transactionspb.Transaction{
				TransactionID: "1",
			},
		},
		{
			name: "Not_Found",
			args: args{
				req: &transactionspb.GetTransactionRequest{
					TransactionID: "1",
				},
			},
			mockCalls: func(m *storage.MockTransactionStore) {
				m.
					EXPECT().
					GetTransaction(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("not found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.mockCalls(m)
			h := NewHandler(m)

			// Act
			got, err := h.GetTransaction(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.GetTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_GetTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := storage.NewMockTransactionStore(ctrl)

	type args struct {
		req *transactionspb.GetTransactionsRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *storage.MockTransactionStore)
		want      *transactionspb.GetTransactionsResponse
		wantErr   bool
	}{
		{
			name: "Success",
			args: args{
				req: &transactionspb.GetTransactionsRequest{
					NextStartingIndex: "1",
				},
			},
			mockCalls: func(m *storage.MockTransactionStore) {
				m.
					EXPECT().
					GetTransactions(gomock.Any(), gomock.Any()).
					Return([]*transactionspb.Transaction{
						&transactionspb.Transaction{
							TransactionID: "1",
						},
					}, true, nil)
			},
			want: &transactionspb.GetTransactionsResponse{
				Result: []*transactionspb.Transaction{
					&transactionspb.Transaction{
						TransactionID: "1",
					},
				},
				HasMore: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.mockCalls(m)
			h := NewHandler(m)

			// Act
			got, err := h.GetTransactions(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.GetTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := storage.NewMockTransactionStore(ctrl)

	type args struct {
		req *transactionspb.CreateTransactionRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *storage.MockTransactionStore)
		wantErr   bool
	}{
		{
			name: "Success",
			args: args{
				req: &transactionspb.CreateTransactionRequest{
					SourceAccountID:      "1",
					DestinationAccountID: "2",
				},
			},
			mockCalls: func(m *storage.MockTransactionStore) {
				m.
					EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "Storage_Failed",
			args: args{
				req: &transactionspb.CreateTransactionRequest{
					SourceAccountID:      "1",
					DestinationAccountID: "2",
				},
			},
			mockCalls: func(m *storage.MockTransactionStore) {
				m.
					EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("failed"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.mockCalls(m)
			h := NewHandler(m)

			// Act
			got, err := h.CreateTransaction(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !tt.wantErr && got == nil {
				t.Errorf("Handler.CreateTransaction() got should not be empty")
			}
		})
	}
}
