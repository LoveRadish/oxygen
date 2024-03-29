// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AdvisoryTxLock(ctx context.Context, pgAdvisoryXactLock int64) error
	CalculateCustomerPayments(ctx context.Context, arg CalculateCustomerPaymentsParams) (int64, error)
	CancelTransaction(ctx context.Context, arg CancelTransactionParams) error
	CheckSystemWalletExistsByAddress(ctx context.Context, address string) (Wallet, error)
	CreateAPIToken(ctx context.Context, arg CreateAPITokenParams) (ApiToken, error)
	CreateBalance(ctx context.Context, arg CreateBalanceParams) (Balance, error)
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateJobLog(ctx context.Context, arg CreateJobLogParams) error
	CreateMerchant(ctx context.Context, arg CreateMerchantParams) (Merchant, error)
	CreateMerchantAddress(ctx context.Context, arg CreateMerchantAddressParams) (MerchantAddress, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	CreatePaymentLink(ctx context.Context, arg CreatePaymentLinkParams) (PaymentLink, error)
	CreateRegistryItem(ctx context.Context, arg CreateRegistryItemParams) (Registry, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error)
	CreateWalletLock(ctx context.Context, arg CreateWalletLockParams) (WalletLock, error)
	DeleteAPITokenByID(ctx context.Context, id int64) error
	DeleteAPITokenByToken(ctx context.Context, token string) error
	DeleteMerchantAddress(ctx context.Context, arg DeleteMerchantAddressParams) error
	DeletePaymentLinkByPublicID(ctx context.Context, arg DeletePaymentLinkByPublicIDParams) error
	DeleteUser(ctx context.Context, id int64) error
	EagerLoadTransactionsByPaymentID(ctx context.Context, arg EagerLoadTransactionsByPaymentIDParams) ([]Transaction, error)
	GetAPIToken(ctx context.Context, arg GetAPITokenParams) (ApiToken, error)
	GetAPITokenByUUID(ctx context.Context, uuid uuid.UUID) (ApiToken, error)
	GetAvailableWallet(ctx context.Context, arg GetAvailableWalletParams) (Wallet, error)
	GetBalanceByFilter(ctx context.Context, arg GetBalanceByFilterParams) (Balance, error)
	GetBalanceByFilterWithLock(ctx context.Context, arg GetBalanceByFilterWithLockParams) (Balance, error)
	GetBalanceByID(ctx context.Context, arg GetBalanceByIDParams) (Balance, error)
	GetBalanceByIDWithLock(ctx context.Context, id int64) (Balance, error)
	GetBalanceByUUID(ctx context.Context, arg GetBalanceByUUIDParams) (Balance, error)
	GetBatchCustomers(ctx context.Context, arg GetBatchCustomersParams) ([]Customer, error)
	GetBatchExpiredPayments(ctx context.Context, arg GetBatchExpiredPaymentsParams) ([]Payment, error)
	GetCustomerByEmail(ctx context.Context, arg GetCustomerByEmailParams) (Customer, error)
	GetCustomerByID(ctx context.Context, arg GetCustomerByIDParams) (Customer, error)
	GetCustomerByUUID(ctx context.Context, arg GetCustomerByUUIDParams) (Customer, error)
	GetLatestTransactionByPaymentID(ctx context.Context, entityID sql.NullInt64) (Transaction, error)
	GetMerchantAddressByAddress(ctx context.Context, arg GetMerchantAddressByAddressParams) (MerchantAddress, error)
	GetMerchantAddressByID(ctx context.Context, arg GetMerchantAddressByIDParams) (MerchantAddress, error)
	GetMerchantAddressByUUID(ctx context.Context, arg GetMerchantAddressByUUIDParams) (MerchantAddress, error)
	GetMerchantByID(ctx context.Context, arg GetMerchantByIDParams) (Merchant, error)
	GetMerchantByUUID(ctx context.Context, arg GetMerchantByUUIDParams) (Merchant, error)
	GetMerchantByUUIDAndCreatorID(ctx context.Context, arg GetMerchantByUUIDAndCreatorIDParams) (Merchant, error)
	GetPaymentByID(ctx context.Context, arg GetPaymentByIDParams) (Payment, error)
	GetPaymentByMerchantIDAndOrderUUID(ctx context.Context, arg GetPaymentByMerchantIDAndOrderUUIDParams) (Payment, error)
	GetPaymentByMerchantIDs(ctx context.Context, arg GetPaymentByMerchantIDsParams) (Payment, error)
	GetPaymentByPublicID(ctx context.Context, publicID uuid.UUID) (Payment, error)
	GetPaymentLinkByID(ctx context.Context, arg GetPaymentLinkByIDParams) (PaymentLink, error)
	GetPaymentLinkByPublicID(ctx context.Context, arg GetPaymentLinkByPublicIDParams) (PaymentLink, error)
	GetPaymentLinkBySlug(ctx context.Context, slug string) (PaymentLink, error)
	GetPaymentsByType(ctx context.Context, arg GetPaymentsByTypeParams) ([]Payment, error)
	GetRecentCustomerPayments(ctx context.Context, arg GetRecentCustomerPaymentsParams) ([]Payment, error)
	GetRegistryItemByKey(ctx context.Context, arg GetRegistryItemByKeyParams) (Registry, error)
	GetTransactionByHashAndNetworkID(ctx context.Context, arg GetTransactionByHashAndNetworkIDParams) (Transaction, error)
	GetTransactionByID(ctx context.Context, arg GetTransactionByIDParams) (Transaction, error)
	GetTransactionsByFilter(ctx context.Context, arg GetTransactionsByFilterParams) ([]Transaction, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByGoogleID(ctx context.Context, googleID sql.NullString) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetWalletByID(ctx context.Context, id int64) (Wallet, error)
	GetWalletByUUID(ctx context.Context, uuid uuid.UUID) (Wallet, error)
	GetWalletForUpdateByID(ctx context.Context, id int64) (Wallet, error)
	GetWalletLock(ctx context.Context, arg GetWalletLockParams) (WalletLock, error)
	InsertBalanceAuditLog(ctx context.Context, arg InsertBalanceAuditLogParams) error
	ListAPITokensByEntity(ctx context.Context, arg ListAPITokensByEntityParams) ([]ApiToken, error)
	ListAllBalancesByType(ctx context.Context, arg ListAllBalancesByTypeParams) ([]Balance, error)
	ListBalances(ctx context.Context, arg ListBalancesParams) ([]Balance, error)
	ListJobLogsByID(ctx context.Context, arg ListJobLogsByIDParams) ([]JobLog, error)
	ListMerchantAddresses(ctx context.Context, merchantID int64) ([]MerchantAddress, error)
	ListMerchantsByCreatorID(ctx context.Context, arg ListMerchantsByCreatorIDParams) ([]Merchant, error)
	ListPaymentLinks(ctx context.Context, arg ListPaymentLinksParams) ([]PaymentLink, error)
	ListUsers(ctx context.Context) ([]User, error)
	PaginateCustomersAsc(ctx context.Context, arg PaginateCustomersAscParams) ([]Customer, error)
	PaginateCustomersDesc(ctx context.Context, arg PaginateCustomersDescParams) ([]Customer, error)
	PaginatePaymentsAsc(ctx context.Context, arg PaginatePaymentsAscParams) ([]Payment, error)
	PaginatePaymentsDesc(ctx context.Context, arg PaginatePaymentsDescParams) ([]Payment, error)
	PaginateWalletsByID(ctx context.Context, arg PaginateWalletsByIDParams) ([]Wallet, error)
	ReleaseWalletLock(ctx context.Context, id int64) error
	SetTransactionHash(ctx context.Context, arg SetTransactionHashParams) error
	SoftDeleteMerchantByUUID(ctx context.Context, uuid uuid.UUID) error
	UpdateBalanceByID(ctx context.Context, arg UpdateBalanceByIDParams) (Balance, error)
	UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) (Merchant, error)
	UpdateMerchantAddress(ctx context.Context, arg UpdateMerchantAddressParams) (MerchantAddress, error)
	UpdateMerchantSettings(ctx context.Context, arg UpdateMerchantSettingsParams) error
	UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error)
	UpdatePaymentCustomerID(ctx context.Context, arg UpdatePaymentCustomerIDParams) error
	UpdatePaymentWebhookInfo(ctx context.Context, arg UpdatePaymentWebhookInfoParams) error
	UpdateRegistryItem(ctx context.Context, arg UpdateRegistryItemParams) (Registry, error)
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateWalletMainnetTransactionCounters(ctx context.Context, arg UpdateWalletMainnetTransactionCountersParams) error
	UpdateWalletTatumFields(ctx context.Context, arg UpdateWalletTatumFieldsParams) (Wallet, error)
	UpdateWalletTestnetTransactionCounters(ctx context.Context, arg UpdateWalletTestnetTransactionCountersParams) error
}

var _ Querier = (*Queries)(nil)
