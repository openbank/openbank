package accounts

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-cmp/cmp"

	accountspb "github.com/openbank/gunk/gunk/v1/accounts"
	typespb "github.com/openbank/gunk/gunk/v1/types"
	"github.com/openbank/openbank/storage/mock_storage"
)

func TestHandler_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_storage.NewMockAccountStore(ctrl)

	type args struct {
		req *accountspb.GetAccountRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *mock_storage.MockAccountStore)
		want      *accountspb.Account
		wantErr   bool
	}{
		{
			name: "get account ok",
			args: args{
				req: &accountspb.GetAccountRequest{
					AccountID: "1",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Return(&accountspb.Account{
						AccountID: "1",
					}, nil)
			},
			want: &accountspb.Account{
				AccountID: "1",
			},
		},
		{
			name: "get account not found",
			args: args{
				req: &accountspb.GetAccountRequest{
					AccountID: "1",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
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
			got, err := h.GetAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_GetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_storage.NewMockAccountStore(ctrl)

	type args struct {
		req *accountspb.GetAccountsRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *mock_storage.MockAccountStore)
		want      *accountspb.GetAccountsResponse
		wantErr   bool
	}{
		{
			name: "get accounts ok",
			args: args{
				req: &accountspb.GetAccountsRequest{
					NextStartingIndex: "1",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					GetAccounts(gomock.Any(), gomock.Any()).
					Return([]*accountspb.Account{
						&accountspb.Account{
							AccountID: "1",
						},
					}, true, nil)
			},
			want: &accountspb.GetAccountsResponse{
				Result: []*accountspb.Account{
					&accountspb.Account{
						AccountID: "1",
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
			got, err := h.GetAccounts(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_storage.NewMockAccountStore(ctrl)

	type args struct {
		req *accountspb.CreateAccountRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *mock_storage.MockAccountStore)
		want      *accountspb.CreateAccountResponse
		wantErr   bool
	}{
		{
			name: "create account ok",
			args: args{
				req: &accountspb.CreateAccountRequest{
					AccountID:    "1",
					Branch:       "branch",
					InterestRate: "rate",
					MajorType:    typespb.MajorType_Checking,
					MaturityDate: "date",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			want: &accountspb.CreateAccountResponse{
				AccountID: "1",
			},
		},
		{
			name: "create account failed due to duplicate account",
			args: args{
				req: &accountspb.CreateAccountRequest{
					AccountID:    "1",
					Branch:       "branch",
					InterestRate: "rate",
					MajorType:    typespb.MajorType_Checking,
					MaturityDate: "date",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("conflict"))
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
			got, err := h.CreateAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_UpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_storage.NewMockAccountStore(ctrl)

	type args struct {
		req *accountspb.UpdateAccountRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *mock_storage.MockAccountStore)
		want      *empty.Empty
		wantErr   bool
	}{
		{
			name: "update account ok",
			args: args{
				req: &accountspb.UpdateAccountRequest{
					AccountID:   "1",
					Description: "description",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					UpdateAccount(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			want: &empty.Empty{},
		},
		{
			name: "update account not found",
			args: args{
				req: &accountspb.UpdateAccountRequest{
					AccountID:   "1",
					Description: "description",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					UpdateAccount(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("not found"))
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
			got, err := h.UpdateAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.UpdateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_DeleteAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_storage.NewMockAccountStore(ctrl)

	type args struct {
		req *accountspb.DeleteAccountRequest
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(m *mock_storage.MockAccountStore)
		want      *empty.Empty
		wantErr   bool
	}{
		{
			name: "delete account ok",
			args: args{
				req: &accountspb.DeleteAccountRequest{
					AccountID: "1",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			want: &empty.Empty{},
		},
		{
			name: "delete account not found",
			args: args{
				req: &accountspb.DeleteAccountRequest{
					AccountID: "1",
				},
			},
			mockCalls: func(m *mock_storage.MockAccountStore) {
				m.
					EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("not found"))
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
			got, err := h.DeleteAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Handler.DeleteAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
