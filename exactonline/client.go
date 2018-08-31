// go:generate go run gen-services.go -v

package exactonline

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

const (
	defaultBaseURL = "https://start.exactonline.nl/"
	userAgent      = "go-exactonline"
)

// A Client manages communication with the Exact Online API.
type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	// BaseURL for API requests. Defaults to the Dutch API. See more available base urls in the API documentation. @TODO
	BaseURL *url.URL

	// UserAgent used when communicating with the Exact Online API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Exact Online API
	AccountancyAccountInvolvedAccounts                    *AccountancyAccountInvolvedAccountsService
	AccountancyAccountOwners                              *AccountancyAccountOwnersService
	AccountancyInvolvedUserRoles                          *AccountancyInvolvedUserRolesService
	AccountancyInvolvedUsers                              *AccountancyInvolvedUsersService
	AccountancySolutionLinks                              *AccountancySolutionLinksService
	AccountancyTaskTypes                                  *AccountancyTaskTypesService
	ActivitiesCommunicationNotes                          *ActivitiesCommunicationNotesService
	ActivitiesComplaints                                  *ActivitiesComplaintsService
	ActivitiesEvents                                      *ActivitiesEventsService
	ActivitiesServiceRequests                             *ActivitiesServiceRequestsService
	ActivitiesTasks                                       *ActivitiesTasksService
	AssetsAssetGroups                                     *AssetsAssetGroupsService
	AssetsAssets                                          *AssetsAssetsService
	AssetsDepreciationMethods                             *AssetsDepreciationMethodsService
	BudgetBudgets                                         *BudgetBudgetsService
	BudgetBudgetScenarios                                 *BudgetBudgetScenariosService
	BulkDocumentsDocumentAttachments                      *BulkDocumentsDocumentAttachmentsService
	BulkDocumentsDocuments                                *BulkDocumentsDocumentsService
	BulkFinancialTransactionLines                         *BulkFinancialTransactionLinesService
	BulkSalesOrderSalesOrderLines                         *BulkSalesOrderSalesOrderLinesService
	CashflowAllocationRule                                *CashflowAllocationRuleService
	CashflowBanks                                         *CashflowBanksService
	CashflowDirectDebitMandates                           *CashflowDirectDebitMandatesService
	CashflowImportNotificationDetails                     *CashflowImportNotificationDetailsService
	CashflowImportNotifications                           *CashflowImportNotificationsService
	CashflowPaymentConditions                             *CashflowPaymentConditionsService
	CashflowPayments                                      *CashflowPaymentsService
	CashflowReceivables                                   *CashflowReceivablesService
	ContinuousMonitoringIndicatorBalances                 *ContinuousMonitoringIndicatorBalancesService
	ContinuousMonitoringIndicatorDeviatingAmountEntereds  *ContinuousMonitoringIndicatorDeviatingAmountEnteredsService
	ContinuousMonitoringIndicatorDifferenceByPeriods      *ContinuousMonitoringIndicatorDifferenceByPeriodsService
	ContinuousMonitoringIndicatorDifferentVatCodes        *ContinuousMonitoringIndicatorDifferentVatCodesService
	ContinuousMonitoringIndicatorGLAccounts               *ContinuousMonitoringIndicatorGLAccountsService
	ContinuousMonitoringIndicatorLiquidities              *ContinuousMonitoringIndicatorLiquiditiesService
	ContinuousMonitoringIndicatorSignals                  *ContinuousMonitoringIndicatorSignalsService
	ContinuousMonitoringIndicatorStates                   *ContinuousMonitoringIndicatorStatesService
	ContinuousMonitoringIndicatorUsageOfJournals          *ContinuousMonitoringIndicatorUsageOfJournalsService
	CRMAccountClasses                                     *CRMAccountClassesService
	CRMAccountClassificationNames                         *CRMAccountClassificationNamesService
	CRMAccountClassifications                             *CRMAccountClassificationsService
	CRMAccounts                                           *CRMAccountsService
	CRMAddresses                                          *CRMAddressesService
	CRMAddressStates                                      *CRMAddressStatesService
	CRMBankAccounts                                       *CRMBankAccountsService
	CRMContacts                                           *CRMContactsService
	CRMDocuments                                          *CRMDocumentsService
	CRMDocumentsAttachments                               *CRMDocumentsAttachmentsService
	CRMHostingOpportunities                               *CRMHostingOpportunitiesService
	CRMOpportunities                                      *CRMOpportunitiesService
	CRMOpportunityContacts                                *CRMOpportunityContactsService
	CRMQuotationLines                                     *CRMQuotationLinesService
	CRMQuotations                                         *CRMQuotationsService
	CRMReasonCodes                                        *CRMReasonCodesService
	DocumentsDocumentAttachments                          *DocumentsDocumentAttachmentsService
	DocumentsDocumentCategories                           *DocumentsDocumentCategoriesService
	DocumentsDocumentFolders                              *DocumentsDocumentFoldersService
	DocumentsDocuments                                    *DocumentsDocumentsService
	DocumentsDocumentTypeCategories                       *DocumentsDocumentTypeCategoriesService
	DocumentsDocumentTypeFolders                          *DocumentsDocumentTypeFoldersService
	DocumentsDocumentTypes                                *DocumentsDocumentTypesService
	FinancialAgingOverview                                *FinancialAgingOverviewService
	FinancialAgingPayablesList                            *FinancialAgingPayablesListService
	FinancialAgingReceivablesList                         *FinancialAgingReceivablesListService
	FinancialExchangeRates                                *FinancialExchangeRatesService
	FinancialFinancialPeriods                             *FinancialFinancialPeriodsService
	FinancialGLAccountClassificationMappings              *FinancialGLAccountClassificationMappingsService
	FinancialGLAccounts                                   *FinancialGLAccountsService
	FinancialGLClassifications                            *FinancialGLClassificationsService
	FinancialGLSchemes                                    *FinancialGLSchemesService
	FinancialGLTransactionTypes                           *FinancialGLTransactionTypesService
	FinancialJournals                                     *FinancialJournalsService
	FinancialJournalStatusList                            *FinancialJournalStatusListService
	FinancialOutstandingInvoicesOverview                  *FinancialOutstandingInvoicesOverviewService
	FinancialPayablesList                                 *FinancialPayablesListService
	FinancialProfitLossOverview                           *FinancialProfitLossOverviewService
	FinancialReceivablesList                              *FinancialReceivablesListService
	FinancialReportingBalance                             *FinancialReportingBalanceService
	FinancialReturns                                      *FinancialReturnsService
	FinancialRevenueList                                  *FinancialRevenueListService
	FinancialTransactionBankEntries                       *FinancialTransactionBankEntriesService
	FinancialTransactionBankEntryLines                    *FinancialTransactionBankEntryLinesService
	FinancialTransactionCashEntries                       *FinancialTransactionCashEntriesService
	FinancialTransactionCashEntryLines                    *FinancialTransactionCashEntryLinesService
	FinancialTransactionTransactionLines                  *FinancialTransactionTransactionLinesService
	FinancialTransactionTransactions                      *FinancialTransactionTransactionsService
	GeneralCurrencies                                     *GeneralCurrenciesService
	GeneralJournalEntryGeneralJournalEntries              *GeneralJournalEntryGeneralJournalEntriesService
	GeneralJournalEntryGeneralJournalEntryLines           *GeneralJournalEntryGeneralJournalEntryLinesService
	HRMAbsenceRegistrations                               *HRMAbsenceRegistrationsService
	HRMAbsenceRegistrationTransactions                    *HRMAbsenceRegistrationTransactionsService
	HRMCostcenters                                        *HRMCostcentersService
	HRMCostunits                                          *HRMCostunitsService
	HRMDepartments                                        *HRMDepartmentsService
	HRMDivisionClasses                                    *HRMDivisionClassesService
	HRMDivisionClassNames                                 *HRMDivisionClassNamesService
	HRMDivisionClassValues                                *HRMDivisionClassValuesService
	HRMDivisions                                          *HRMDivisionsService
	HRMJobGroups                                          *HRMJobGroupsService
	HRMJobTitles                                          *HRMJobTitlesService
	HRMLeaveBuildUpRegistrations                          *HRMLeaveBuildUpRegistrationsService
	HRMLeaveRegistrations                                 *HRMLeaveRegistrationsService
	HRMSchedules                                          *HRMSchedulesService
	InventoryAssemblyOrders                               *InventoryAssemblyOrdersService
	InventoryBatchNumbers                                 *InventoryBatchNumbersService
	InventoryItemWarehousePlanningDetails                 *InventoryItemWarehousePlanningDetailsService
	InventoryItemWarehouses                               *InventoryItemWarehousesService
	InventoryItemWarehouseStorageLocations                *InventoryItemWarehouseStorageLocationsService
	InventorySerialNumbers                                *InventorySerialNumbersService
	InventoryStockBatchNumbers                            *InventoryStockBatchNumbersService
	InventoryStockCountLines                              *InventoryStockCountLinesService
	InventoryStockCounts                                  *InventoryStockCountsService
	InventoryStockSerialNumbers                           *InventoryStockSerialNumbersService
	InventoryStorageLocations                             *InventoryStorageLocationsService
	InventoryWarehouses                                   *InventoryWarehousesService
	LogisticsItemGroups                                   *LogisticsItemGroupsService
	LogisticsItems                                        *LogisticsItemsService
	LogisticsItemVersions                                 *LogisticsItemVersionsService
	LogisticsSalesItemPrices                              *LogisticsSalesItemPricesService
	LogisticsSupplierItem                                 *LogisticsSupplierItemService
	LogisticsUnits                                        *LogisticsUnitsService
	MailboxDefaultMailbox                                 *MailboxDefaultMailboxService
	MailboxMailboxes                                      *MailboxMailboxesService
	MailboxMailMessageAttachments                         *MailboxMailMessageAttachmentsService
	MailboxMailMessagesReceived                           *MailboxMailMessagesReceivedService
	MailboxMailMessagesSent                               *MailboxMailMessagesSentService
	MailboxPreferredMailbox                               *MailboxPreferredMailboxService
	ManufacturingBillOfMaterialMaterials                  *ManufacturingBillOfMaterialMaterialsService
	ManufacturingBillOfMaterialVersions                   *ManufacturingBillOfMaterialVersionsService
	ManufacturingByProductReceipts                        *ManufacturingByProductReceiptsService
	ManufacturingByProductReversals                       *ManufacturingByProductReversalsService
	ManufacturingMaterialIssues                           *ManufacturingMaterialIssuesService
	ManufacturingMaterialReversals                        *ManufacturingMaterialReversalsService
	ManufacturingOperationResources                       *ManufacturingOperationResourcesService
	ManufacturingOperations                               *ManufacturingOperationsService
	ManufacturingProductionAreas                          *ManufacturingProductionAreasService
	ManufacturingShopOrderMaterialPlans                   *ManufacturingShopOrderMaterialPlansService
	ManufacturingShopOrderReceipts                        *ManufacturingShopOrderReceiptsService
	ManufacturingShopOrderReversals                       *ManufacturingShopOrderReversalsService
	ManufacturingShopOrderRoutingStepPlans                *ManufacturingShopOrderRoutingStepPlansService
	ManufacturingShopOrders                               *ManufacturingShopOrdersService
	ManufacturingStageForDeliveryReceipts                 *ManufacturingStageForDeliveryReceiptsService
	ManufacturingStageForDeliveryReversals                *ManufacturingStageForDeliveryReversalsService
	ManufacturingSubOrderReceipts                         *ManufacturingSubOrderReceiptsService
	ManufacturingSubOrderReversals                        *ManufacturingSubOrderReversalsService
	ManufacturingTimeTransactions                         *ManufacturingTimeTransactionsService
	ManufacturingWorkcenters                              *ManufacturingWorkcentersService
	OpeningBalanceCurrentYearAfterEntry                   *OpeningBalanceCurrentYearAfterEntryService
	OpeningBalanceCurrentYearProcessed                    *OpeningBalanceCurrentYearProcessedService
	OpeningBalancePreviousYearAfterEntry                  *OpeningBalancePreviousYearAfterEntryService
	OpeningBalancePreviousYearProcessed                   *OpeningBalancePreviousYearProcessedService
	PayrollActiveEmployments                              *PayrollActiveEmploymentsService
	PayrollEmployees                                      *PayrollEmployeesService
	PayrollEmploymentContractFlexPhases                   *PayrollEmploymentContractFlexPhasesService
	PayrollEmploymentContracts                            *PayrollEmploymentContractsService
	PayrollEmploymentEndReasons                           *PayrollEmploymentEndReasonsService
	PayrollEmploymentOrganizations                        *PayrollEmploymentOrganizationsService
	PayrollEmployments                                    *PayrollEmploymentsService
	PayrollEmploymentSalaries                             *PayrollEmploymentSalariesService
	PayrollTaxEmploymentEndFlexCodes                      *PayrollTaxEmploymentEndFlexCodesService
	ProjectCostTransactions                               *ProjectCostTransactionsService
	ProjectHourCostTypes                                  *ProjectHourCostTypesService
	ProjectInvoiceTerms                                   *ProjectInvoiceTermsService
	ProjectProjectBudgetTypes                             *ProjectProjectBudgetTypesService
	ProjectProjectHourBudgets                             *ProjectProjectHourBudgetsService
	ProjectProjectPlanning                                *ProjectProjectPlanningService
	ProjectProjectPlanningRecurring                       *ProjectProjectPlanningRecurringService
	ProjectProjectRestrictionEmployees                    *ProjectProjectRestrictionEmployeesService
	ProjectProjectRestrictionItems                        *ProjectProjectRestrictionItemsService
	ProjectProjectRestrictionRebillings                   *ProjectProjectRestrictionRebillingsService
	ProjectProjects                                       *ProjectProjectsService
	ProjectRecentCosts                                    *ProjectRecentCostsService
	ProjectRecentHours                                    *ProjectRecentHoursService
	ProjectTimeAndBillingAccountDetails                   *ProjectTimeAndBillingAccountDetailsService
	ProjectTimeAndBillingActivitiesAndExpenses            *ProjectTimeAndBillingActivitiesAndExpensesService
	ProjectTimeAndBillingEntryAccounts                    *ProjectTimeAndBillingEntryAccountsService
	ProjectTimeAndBillingEntryProjects                    *ProjectTimeAndBillingEntryProjectsService
	ProjectTimeAndBillingEntryRecentAccounts              *ProjectTimeAndBillingEntryRecentAccountsService
	ProjectTimeAndBillingEntryRecentActivitiesAndExpenses *ProjectTimeAndBillingEntryRecentActivitiesAndExpensesService
	ProjectTimeAndBillingEntryRecentHourCostTypes         *ProjectTimeAndBillingEntryRecentHourCostTypesService
	ProjectTimeAndBillingEntryRecentProjects              *ProjectTimeAndBillingEntryRecentProjectsService
	ProjectTimeAndBillingItemDetails                      *ProjectTimeAndBillingItemDetailsService
	ProjectTimeAndBillingProjectDetails                   *ProjectTimeAndBillingProjectDetailsService
	ProjectTimeCorrections                                *ProjectTimeCorrectionsService
	ProjectTimeTransactions                               *ProjectTimeTransactionsService
	PurchasePurchaseInvoiceLines                          *PurchasePurchaseInvoiceLinesService
	PurchasePurchaseInvoices                              *PurchasePurchaseInvoicesService
	PurchaseEntryPurchaseEntries                          *PurchaseEntryPurchaseEntriesService
	PurchaseEntryPurchaseEntryLines                       *PurchaseEntryPurchaseEntryLinesService
	PurchaseOrderGoodsReceiptLines                        *PurchaseOrderGoodsReceiptLinesService
	PurchaseOrderGoodsReceipts                            *PurchaseOrderGoodsReceiptsService
	PurchaseOrderPurchaseOrderLines                       *PurchaseOrderPurchaseOrderLinesService
	PurchaseOrderPurchaseOrders                           *PurchaseOrderPurchaseOrdersService
	SalesPriceLists                                       *SalesPriceListsService
	SalesSalesPriceListDetails                            *SalesSalesPriceListDetailsService
	SalesShippingMethods                                  *SalesShippingMethodsService
	SalesEntrySalesEntries                                *SalesEntrySalesEntriesService
	SalesEntrySalesEntryLines                             *SalesEntrySalesEntryLinesService
	SalesInvoiceLayouts                                   *SalesInvoiceLayoutsService
	SalesInvoiceSalesInvoiceLines                         *SalesInvoiceSalesInvoiceLinesService
	SalesInvoiceSalesInvoices                             *SalesInvoiceSalesInvoicesService
	SalesOrderGoodsDeliveries                             *SalesOrderGoodsDeliveriesService
	SalesOrderGoodsDeliveryLines                          *SalesOrderGoodsDeliveryLinesService
	SalesOrderSalesOrderLines                             *SalesOrderSalesOrderLinesService
	SalesOrderSalesOrders                                 *SalesOrderSalesOrdersService
	SubscriptionSubscriptionLines                         *SubscriptionSubscriptionLinesService
	SubscriptionSubscriptionLineTypes                     *SubscriptionSubscriptionLineTypesService
	SubscriptionSubscriptionReasonCodes                   *SubscriptionSubscriptionReasonCodesService
	SubscriptionSubscriptionRestrictionEmployees          *SubscriptionSubscriptionRestrictionEmployeesService
	SubscriptionSubscriptionRestrictionItems              *SubscriptionSubscriptionRestrictionItemsService
	SubscriptionSubscriptions                             *SubscriptionSubscriptionsService
	SubscriptionSubscriptionTypes                         *SubscriptionSubscriptionTypesService
	SystemAccountantInfo                                  *SystemAccountantInfoService
	SystemAvailableFeatures                               *SystemAvailableFeaturesService
	SystemDivisions                                       *SystemDivisionsService
	SystemGetMostRecentlyUsedDivisions                    *SystemGetMostRecentlyUsedDivisionsService
	SystemMe                                              *SystemMeService
	UsersUserRoles                                        *UsersUserRolesService
	UsersUserRolesPerDivision                             *UsersUserRolesPerDivisionService
	UsersUsers                                            *UsersUsersService
	VATVATCodes                                           *VATVATCodesService
	VATVatPercentages                                     *VATVatPercentagesService
	WebhooksWebhookSubscriptions                          *WebhooksWebhookSubscriptionsService
	WorkflowRequestAttachments                            *WorkflowRequestAttachmentsService
}

type service struct {
	client *Client
}

// NewClient returns a new Exact Online API client. Provide a http.Client that
// will perform the authentication for you (such as that provided by the
// golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL) // #nosec

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.AccountancyAccountInvolvedAccounts = (*AccountancyAccountInvolvedAccountsService)(&c.common)
	c.AccountancyAccountOwners = (*AccountancyAccountOwnersService)(&c.common)
	c.AccountancyInvolvedUserRoles = (*AccountancyInvolvedUserRolesService)(&c.common)
	c.AccountancyInvolvedUsers = (*AccountancyInvolvedUsersService)(&c.common)
	c.AccountancySolutionLinks = (*AccountancySolutionLinksService)(&c.common)
	c.AccountancyTaskTypes = (*AccountancyTaskTypesService)(&c.common)
	c.ActivitiesCommunicationNotes = (*ActivitiesCommunicationNotesService)(&c.common)
	c.ActivitiesComplaints = (*ActivitiesComplaintsService)(&c.common)
	c.ActivitiesEvents = (*ActivitiesEventsService)(&c.common)
	c.ActivitiesServiceRequests = (*ActivitiesServiceRequestsService)(&c.common)
	c.ActivitiesTasks = (*ActivitiesTasksService)(&c.common)
	c.AssetsAssetGroups = (*AssetsAssetGroupsService)(&c.common)
	c.AssetsAssets = (*AssetsAssetsService)(&c.common)
	c.AssetsDepreciationMethods = (*AssetsDepreciationMethodsService)(&c.common)
	c.BudgetBudgets = (*BudgetBudgetsService)(&c.common)
	c.BudgetBudgetScenarios = (*BudgetBudgetScenariosService)(&c.common)
	c.BulkDocumentsDocumentAttachments = (*BulkDocumentsDocumentAttachmentsService)(&c.common)
	c.BulkDocumentsDocuments = (*BulkDocumentsDocumentsService)(&c.common)
	c.BulkFinancialTransactionLines = (*BulkFinancialTransactionLinesService)(&c.common)
	c.BulkSalesOrderSalesOrderLines = (*BulkSalesOrderSalesOrderLinesService)(&c.common)
	c.CashflowAllocationRule = (*CashflowAllocationRuleService)(&c.common)
	c.CashflowBanks = (*CashflowBanksService)(&c.common)
	c.CashflowDirectDebitMandates = (*CashflowDirectDebitMandatesService)(&c.common)
	c.CashflowImportNotificationDetails = (*CashflowImportNotificationDetailsService)(&c.common)
	c.CashflowImportNotifications = (*CashflowImportNotificationsService)(&c.common)
	c.CashflowPaymentConditions = (*CashflowPaymentConditionsService)(&c.common)
	c.CashflowPayments = (*CashflowPaymentsService)(&c.common)
	c.CashflowReceivables = (*CashflowReceivablesService)(&c.common)
	c.ContinuousMonitoringIndicatorBalances = (*ContinuousMonitoringIndicatorBalancesService)(&c.common)
	c.ContinuousMonitoringIndicatorDeviatingAmountEntereds = (*ContinuousMonitoringIndicatorDeviatingAmountEnteredsService)(&c.common)
	c.ContinuousMonitoringIndicatorDifferenceByPeriods = (*ContinuousMonitoringIndicatorDifferenceByPeriodsService)(&c.common)
	c.ContinuousMonitoringIndicatorDifferentVatCodes = (*ContinuousMonitoringIndicatorDifferentVatCodesService)(&c.common)
	c.ContinuousMonitoringIndicatorGLAccounts = (*ContinuousMonitoringIndicatorGLAccountsService)(&c.common)
	c.ContinuousMonitoringIndicatorLiquidities = (*ContinuousMonitoringIndicatorLiquiditiesService)(&c.common)
	c.ContinuousMonitoringIndicatorSignals = (*ContinuousMonitoringIndicatorSignalsService)(&c.common)
	c.ContinuousMonitoringIndicatorStates = (*ContinuousMonitoringIndicatorStatesService)(&c.common)
	c.ContinuousMonitoringIndicatorUsageOfJournals = (*ContinuousMonitoringIndicatorUsageOfJournalsService)(&c.common)
	c.CRMAccountClasses = (*CRMAccountClassesService)(&c.common)
	c.CRMAccountClassificationNames = (*CRMAccountClassificationNamesService)(&c.common)
	c.CRMAccountClassifications = (*CRMAccountClassificationsService)(&c.common)
	c.CRMAccounts = (*CRMAccountsService)(&c.common)
	c.CRMAddresses = (*CRMAddressesService)(&c.common)
	c.CRMAddressStates = (*CRMAddressStatesService)(&c.common)
	c.CRMBankAccounts = (*CRMBankAccountsService)(&c.common)
	c.CRMContacts = (*CRMContactsService)(&c.common)
	c.CRMDocuments = (*CRMDocumentsService)(&c.common)
	c.CRMDocumentsAttachments = (*CRMDocumentsAttachmentsService)(&c.common)
	c.CRMHostingOpportunities = (*CRMHostingOpportunitiesService)(&c.common)
	c.CRMOpportunities = (*CRMOpportunitiesService)(&c.common)
	c.CRMOpportunityContacts = (*CRMOpportunityContactsService)(&c.common)
	c.CRMQuotationLines = (*CRMQuotationLinesService)(&c.common)
	c.CRMQuotations = (*CRMQuotationsService)(&c.common)
	c.CRMReasonCodes = (*CRMReasonCodesService)(&c.common)
	c.DocumentsDocumentAttachments = (*DocumentsDocumentAttachmentsService)(&c.common)
	c.DocumentsDocumentCategories = (*DocumentsDocumentCategoriesService)(&c.common)
	c.DocumentsDocumentFolders = (*DocumentsDocumentFoldersService)(&c.common)
	c.DocumentsDocuments = (*DocumentsDocumentsService)(&c.common)
	c.DocumentsDocumentTypeCategories = (*DocumentsDocumentTypeCategoriesService)(&c.common)
	c.DocumentsDocumentTypeFolders = (*DocumentsDocumentTypeFoldersService)(&c.common)
	c.DocumentsDocumentTypes = (*DocumentsDocumentTypesService)(&c.common)
	c.FinancialAgingOverview = (*FinancialAgingOverviewService)(&c.common)
	c.FinancialAgingPayablesList = (*FinancialAgingPayablesListService)(&c.common)
	c.FinancialAgingReceivablesList = (*FinancialAgingReceivablesListService)(&c.common)
	c.FinancialExchangeRates = (*FinancialExchangeRatesService)(&c.common)
	c.FinancialFinancialPeriods = (*FinancialFinancialPeriodsService)(&c.common)
	c.FinancialGLAccountClassificationMappings = (*FinancialGLAccountClassificationMappingsService)(&c.common)
	c.FinancialGLAccounts = (*FinancialGLAccountsService)(&c.common)
	c.FinancialGLClassifications = (*FinancialGLClassificationsService)(&c.common)
	c.FinancialGLSchemes = (*FinancialGLSchemesService)(&c.common)
	c.FinancialGLTransactionTypes = (*FinancialGLTransactionTypesService)(&c.common)
	c.FinancialJournals = (*FinancialJournalsService)(&c.common)
	c.FinancialJournalStatusList = (*FinancialJournalStatusListService)(&c.common)
	c.FinancialOutstandingInvoicesOverview = (*FinancialOutstandingInvoicesOverviewService)(&c.common)
	c.FinancialPayablesList = (*FinancialPayablesListService)(&c.common)
	c.FinancialProfitLossOverview = (*FinancialProfitLossOverviewService)(&c.common)
	c.FinancialReceivablesList = (*FinancialReceivablesListService)(&c.common)
	c.FinancialReportingBalance = (*FinancialReportingBalanceService)(&c.common)
	c.FinancialReturns = (*FinancialReturnsService)(&c.common)
	c.FinancialRevenueList = (*FinancialRevenueListService)(&c.common)
	c.FinancialTransactionBankEntries = (*FinancialTransactionBankEntriesService)(&c.common)
	c.FinancialTransactionBankEntryLines = (*FinancialTransactionBankEntryLinesService)(&c.common)
	c.FinancialTransactionCashEntries = (*FinancialTransactionCashEntriesService)(&c.common)
	c.FinancialTransactionCashEntryLines = (*FinancialTransactionCashEntryLinesService)(&c.common)
	c.FinancialTransactionTransactionLines = (*FinancialTransactionTransactionLinesService)(&c.common)
	c.FinancialTransactionTransactions = (*FinancialTransactionTransactionsService)(&c.common)
	c.GeneralCurrencies = (*GeneralCurrenciesService)(&c.common)
	c.GeneralJournalEntryGeneralJournalEntries = (*GeneralJournalEntryGeneralJournalEntriesService)(&c.common)
	c.GeneralJournalEntryGeneralJournalEntryLines = (*GeneralJournalEntryGeneralJournalEntryLinesService)(&c.common)
	c.HRMAbsenceRegistrations = (*HRMAbsenceRegistrationsService)(&c.common)
	c.HRMAbsenceRegistrationTransactions = (*HRMAbsenceRegistrationTransactionsService)(&c.common)
	c.HRMCostcenters = (*HRMCostcentersService)(&c.common)
	c.HRMCostunits = (*HRMCostunitsService)(&c.common)
	c.HRMDepartments = (*HRMDepartmentsService)(&c.common)
	c.HRMDivisionClasses = (*HRMDivisionClassesService)(&c.common)
	c.HRMDivisionClassNames = (*HRMDivisionClassNamesService)(&c.common)
	c.HRMDivisionClassValues = (*HRMDivisionClassValuesService)(&c.common)
	c.HRMDivisions = (*HRMDivisionsService)(&c.common)
	c.HRMJobGroups = (*HRMJobGroupsService)(&c.common)
	c.HRMJobTitles = (*HRMJobTitlesService)(&c.common)
	c.HRMLeaveBuildUpRegistrations = (*HRMLeaveBuildUpRegistrationsService)(&c.common)
	c.HRMLeaveRegistrations = (*HRMLeaveRegistrationsService)(&c.common)
	c.HRMSchedules = (*HRMSchedulesService)(&c.common)
	c.InventoryAssemblyOrders = (*InventoryAssemblyOrdersService)(&c.common)
	c.InventoryBatchNumbers = (*InventoryBatchNumbersService)(&c.common)
	c.InventoryItemWarehousePlanningDetails = (*InventoryItemWarehousePlanningDetailsService)(&c.common)
	c.InventoryItemWarehouses = (*InventoryItemWarehousesService)(&c.common)
	c.InventoryItemWarehouseStorageLocations = (*InventoryItemWarehouseStorageLocationsService)(&c.common)
	c.InventorySerialNumbers = (*InventorySerialNumbersService)(&c.common)
	c.InventoryStockBatchNumbers = (*InventoryStockBatchNumbersService)(&c.common)
	c.InventoryStockCountLines = (*InventoryStockCountLinesService)(&c.common)
	c.InventoryStockCounts = (*InventoryStockCountsService)(&c.common)
	c.InventoryStockSerialNumbers = (*InventoryStockSerialNumbersService)(&c.common)
	c.InventoryStorageLocations = (*InventoryStorageLocationsService)(&c.common)
	c.InventoryWarehouses = (*InventoryWarehousesService)(&c.common)
	c.LogisticsItemGroups = (*LogisticsItemGroupsService)(&c.common)
	c.LogisticsItems = (*LogisticsItemsService)(&c.common)
	c.LogisticsItemVersions = (*LogisticsItemVersionsService)(&c.common)
	c.LogisticsSalesItemPrices = (*LogisticsSalesItemPricesService)(&c.common)
	c.LogisticsSupplierItem = (*LogisticsSupplierItemService)(&c.common)
	c.LogisticsUnits = (*LogisticsUnitsService)(&c.common)
	c.MailboxDefaultMailbox = (*MailboxDefaultMailboxService)(&c.common)
	c.MailboxMailboxes = (*MailboxMailboxesService)(&c.common)
	c.MailboxMailMessageAttachments = (*MailboxMailMessageAttachmentsService)(&c.common)
	c.MailboxMailMessagesReceived = (*MailboxMailMessagesReceivedService)(&c.common)
	c.MailboxMailMessagesSent = (*MailboxMailMessagesSentService)(&c.common)
	c.MailboxPreferredMailbox = (*MailboxPreferredMailboxService)(&c.common)
	c.ManufacturingBillOfMaterialMaterials = (*ManufacturingBillOfMaterialMaterialsService)(&c.common)
	c.ManufacturingBillOfMaterialVersions = (*ManufacturingBillOfMaterialVersionsService)(&c.common)
	c.ManufacturingByProductReceipts = (*ManufacturingByProductReceiptsService)(&c.common)
	c.ManufacturingByProductReversals = (*ManufacturingByProductReversalsService)(&c.common)
	c.ManufacturingMaterialIssues = (*ManufacturingMaterialIssuesService)(&c.common)
	c.ManufacturingMaterialReversals = (*ManufacturingMaterialReversalsService)(&c.common)
	c.ManufacturingOperationResources = (*ManufacturingOperationResourcesService)(&c.common)
	c.ManufacturingOperations = (*ManufacturingOperationsService)(&c.common)
	c.ManufacturingProductionAreas = (*ManufacturingProductionAreasService)(&c.common)
	c.ManufacturingShopOrderMaterialPlans = (*ManufacturingShopOrderMaterialPlansService)(&c.common)
	c.ManufacturingShopOrderReceipts = (*ManufacturingShopOrderReceiptsService)(&c.common)
	c.ManufacturingShopOrderReversals = (*ManufacturingShopOrderReversalsService)(&c.common)
	c.ManufacturingShopOrderRoutingStepPlans = (*ManufacturingShopOrderRoutingStepPlansService)(&c.common)
	c.ManufacturingShopOrders = (*ManufacturingShopOrdersService)(&c.common)
	c.ManufacturingStageForDeliveryReceipts = (*ManufacturingStageForDeliveryReceiptsService)(&c.common)
	c.ManufacturingStageForDeliveryReversals = (*ManufacturingStageForDeliveryReversalsService)(&c.common)
	c.ManufacturingSubOrderReceipts = (*ManufacturingSubOrderReceiptsService)(&c.common)
	c.ManufacturingSubOrderReversals = (*ManufacturingSubOrderReversalsService)(&c.common)
	c.ManufacturingTimeTransactions = (*ManufacturingTimeTransactionsService)(&c.common)
	c.ManufacturingWorkcenters = (*ManufacturingWorkcentersService)(&c.common)
	c.OpeningBalanceCurrentYearAfterEntry = (*OpeningBalanceCurrentYearAfterEntryService)(&c.common)
	c.OpeningBalanceCurrentYearProcessed = (*OpeningBalanceCurrentYearProcessedService)(&c.common)
	c.OpeningBalancePreviousYearAfterEntry = (*OpeningBalancePreviousYearAfterEntryService)(&c.common)
	c.OpeningBalancePreviousYearProcessed = (*OpeningBalancePreviousYearProcessedService)(&c.common)
	c.PayrollActiveEmployments = (*PayrollActiveEmploymentsService)(&c.common)
	c.PayrollEmployees = (*PayrollEmployeesService)(&c.common)
	c.PayrollEmploymentContractFlexPhases = (*PayrollEmploymentContractFlexPhasesService)(&c.common)
	c.PayrollEmploymentContracts = (*PayrollEmploymentContractsService)(&c.common)
	c.PayrollEmploymentEndReasons = (*PayrollEmploymentEndReasonsService)(&c.common)
	c.PayrollEmploymentOrganizations = (*PayrollEmploymentOrganizationsService)(&c.common)
	c.PayrollEmployments = (*PayrollEmploymentsService)(&c.common)
	c.PayrollEmploymentSalaries = (*PayrollEmploymentSalariesService)(&c.common)
	c.PayrollTaxEmploymentEndFlexCodes = (*PayrollTaxEmploymentEndFlexCodesService)(&c.common)
	c.ProjectCostTransactions = (*ProjectCostTransactionsService)(&c.common)
	c.ProjectHourCostTypes = (*ProjectHourCostTypesService)(&c.common)
	c.ProjectInvoiceTerms = (*ProjectInvoiceTermsService)(&c.common)
	c.ProjectProjectBudgetTypes = (*ProjectProjectBudgetTypesService)(&c.common)
	c.ProjectProjectHourBudgets = (*ProjectProjectHourBudgetsService)(&c.common)
	c.ProjectProjectPlanning = (*ProjectProjectPlanningService)(&c.common)
	c.ProjectProjectPlanningRecurring = (*ProjectProjectPlanningRecurringService)(&c.common)
	c.ProjectProjectRestrictionEmployees = (*ProjectProjectRestrictionEmployeesService)(&c.common)
	c.ProjectProjectRestrictionItems = (*ProjectProjectRestrictionItemsService)(&c.common)
	c.ProjectProjectRestrictionRebillings = (*ProjectProjectRestrictionRebillingsService)(&c.common)
	c.ProjectProjects = (*ProjectProjectsService)(&c.common)
	c.ProjectRecentCosts = (*ProjectRecentCostsService)(&c.common)
	c.ProjectRecentHours = (*ProjectRecentHoursService)(&c.common)
	c.ProjectTimeAndBillingAccountDetails = (*ProjectTimeAndBillingAccountDetailsService)(&c.common)
	c.ProjectTimeAndBillingActivitiesAndExpenses = (*ProjectTimeAndBillingActivitiesAndExpensesService)(&c.common)
	c.ProjectTimeAndBillingEntryAccounts = (*ProjectTimeAndBillingEntryAccountsService)(&c.common)
	c.ProjectTimeAndBillingEntryProjects = (*ProjectTimeAndBillingEntryProjectsService)(&c.common)
	c.ProjectTimeAndBillingEntryRecentAccounts = (*ProjectTimeAndBillingEntryRecentAccountsService)(&c.common)
	c.ProjectTimeAndBillingEntryRecentActivitiesAndExpenses = (*ProjectTimeAndBillingEntryRecentActivitiesAndExpensesService)(&c.common)
	c.ProjectTimeAndBillingEntryRecentHourCostTypes = (*ProjectTimeAndBillingEntryRecentHourCostTypesService)(&c.common)
	c.ProjectTimeAndBillingEntryRecentProjects = (*ProjectTimeAndBillingEntryRecentProjectsService)(&c.common)
	c.ProjectTimeAndBillingItemDetails = (*ProjectTimeAndBillingItemDetailsService)(&c.common)
	c.ProjectTimeAndBillingProjectDetails = (*ProjectTimeAndBillingProjectDetailsService)(&c.common)
	c.ProjectTimeCorrections = (*ProjectTimeCorrectionsService)(&c.common)
	c.ProjectTimeTransactions = (*ProjectTimeTransactionsService)(&c.common)
	c.PurchasePurchaseInvoiceLines = (*PurchasePurchaseInvoiceLinesService)(&c.common)
	c.PurchasePurchaseInvoices = (*PurchasePurchaseInvoicesService)(&c.common)
	c.PurchaseEntryPurchaseEntries = (*PurchaseEntryPurchaseEntriesService)(&c.common)
	c.PurchaseEntryPurchaseEntryLines = (*PurchaseEntryPurchaseEntryLinesService)(&c.common)
	c.PurchaseOrderGoodsReceiptLines = (*PurchaseOrderGoodsReceiptLinesService)(&c.common)
	c.PurchaseOrderGoodsReceipts = (*PurchaseOrderGoodsReceiptsService)(&c.common)
	c.PurchaseOrderPurchaseOrderLines = (*PurchaseOrderPurchaseOrderLinesService)(&c.common)
	c.PurchaseOrderPurchaseOrders = (*PurchaseOrderPurchaseOrdersService)(&c.common)
	c.SalesPriceLists = (*SalesPriceListsService)(&c.common)
	c.SalesSalesPriceListDetails = (*SalesSalesPriceListDetailsService)(&c.common)
	c.SalesShippingMethods = (*SalesShippingMethodsService)(&c.common)
	c.SalesEntrySalesEntries = (*SalesEntrySalesEntriesService)(&c.common)
	c.SalesEntrySalesEntryLines = (*SalesEntrySalesEntryLinesService)(&c.common)
	c.SalesInvoiceLayouts = (*SalesInvoiceLayoutsService)(&c.common)
	c.SalesInvoiceSalesInvoiceLines = (*SalesInvoiceSalesInvoiceLinesService)(&c.common)
	c.SalesInvoiceSalesInvoices = (*SalesInvoiceSalesInvoicesService)(&c.common)
	c.SalesOrderGoodsDeliveries = (*SalesOrderGoodsDeliveriesService)(&c.common)
	c.SalesOrderGoodsDeliveryLines = (*SalesOrderGoodsDeliveryLinesService)(&c.common)
	c.SalesOrderSalesOrderLines = (*SalesOrderSalesOrderLinesService)(&c.common)
	c.SalesOrderSalesOrders = (*SalesOrderSalesOrdersService)(&c.common)
	c.SubscriptionSubscriptionLines = (*SubscriptionSubscriptionLinesService)(&c.common)
	c.SubscriptionSubscriptionLineTypes = (*SubscriptionSubscriptionLineTypesService)(&c.common)
	c.SubscriptionSubscriptionReasonCodes = (*SubscriptionSubscriptionReasonCodesService)(&c.common)
	c.SubscriptionSubscriptionRestrictionEmployees = (*SubscriptionSubscriptionRestrictionEmployeesService)(&c.common)
	c.SubscriptionSubscriptionRestrictionItems = (*SubscriptionSubscriptionRestrictionItemsService)(&c.common)
	c.SubscriptionSubscriptions = (*SubscriptionSubscriptionsService)(&c.common)
	c.SubscriptionSubscriptionTypes = (*SubscriptionSubscriptionTypesService)(&c.common)
	c.SystemAccountantInfo = (*SystemAccountantInfoService)(&c.common)
	c.SystemAvailableFeatures = (*SystemAvailableFeaturesService)(&c.common)
	c.SystemDivisions = (*SystemDivisionsService)(&c.common)
	c.SystemGetMostRecentlyUsedDivisions = (*SystemGetMostRecentlyUsedDivisionsService)(&c.common)
	c.SystemMe = (*SystemMeService)(&c.common)
	c.UsersUserRoles = (*UsersUserRolesService)(&c.common)
	c.UsersUserRolesPerDivision = (*UsersUserRolesPerDivisionService)(&c.common)
	c.UsersUsers = (*UsersUsersService)(&c.common)
	c.VATVATCodes = (*VATVATCodesService)(&c.common)
	c.VATVatPercentages = (*VATVatPercentagesService)(&c.common)
	c.WebhooksWebhookSubscriptions = (*WebhooksWebhookSubscriptionsService)(&c.common)
	c.WorkflowRequestAttachments = (*WorkflowRequestAttachmentsService)(&c.common)
	return c
}

// NewClientFromTokenSource is a wrapper around NewClient if you have a valid
// token source. If no context is available you can use context.Background()
func NewClientFromTokenSource(ctx context.Context, tokenSource oauth2.TokenSource) *Client {
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return NewClient(httpClient)
}

// ResolveURL will either return either a resolved path or a valid absolute URI
func (c *Client) ResolveURL(urlStr string) (*url.URL, error) {
	if abs, err := url.Parse(urlStr); err == nil && abs.IsAbs() {
		return abs, nil
	}

	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	return c.BaseURL.Parse(urlStr)
}

// ResolvePathWithDivision will resolve the base url for paths that need a division prefix
func (c *Client) ResolvePathWithDivision(path string, division int) (*url.URL, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	return c.BaseURL.Parse(strings.Replace(path, "{division}", strconv.Itoa(division), 1))
}

// NewRequest creates an API request. An absolute URL must be provided in url.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, parseErr := c.ResolveURL(urlStr)
	if parseErr != nil {
		return nil, parseErr
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return req, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	if err := handleResponseError(resp, req.URL.String()); err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, copyErr := io.Copy(w, resp.Body)
			if copyErr != nil {
				err = copyErr
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// NewRequestAndDo combines NewRequest and Do methods
func (c *Client) NewRequestAndDo(ctx context.Context, method, urlStr string, body, v interface{}) (*http.Request, *http.Response, error) {
	req, e := c.NewRequest(method, urlStr, body)
	if e != nil {
		return req, nil, e
	}
	res, err := c.Do(ctx, req, v)
	return req, res, err
}

// ListRequestAndDo combines NewRequestAndDo and unmarshalls in general ListResponse
func (c *Client) ListRequestAndDo(ctx context.Context, urlStr string, v interface{}) (*ListResponse, *http.Request, *http.Response, error) {
	var listResponse ListResponse
	req, res, err := c.NewRequestAndDo(ctx, "GET", urlStr, nil, &listResponse)
	if err != nil {
		return nil, nil, nil, err
	}

	if v != nil {
		err = json.Unmarshal(listResponse.Data.Results, v)
	}

	return &listResponse, req, res, err
}

// ListRequestAndDoAll requests all paginated pages ListRequestAndDo
func (c *Client) ListRequestAndDoAll(ctx context.Context, urlStr string, v interface{}) error {
	var s []json.RawMessage
	f, _, _, err := c.ListRequestAndDo(ctx, urlStr, &s)
	if err != nil {
		return err
	}

	var next = f.Data.Next
	for next != "" {
		var i []json.RawMessage
		l, _, _, rErr := c.ListRequestAndDo(ctx, next, &i)
		if rErr != nil {
			return rErr
		}
		s = append(s, i...)
		next = l.Data.Next
	}

	err = unmarshalRawMessages(s, v)
	return err
}

func unmarshalRawMessages(m []json.RawMessage, v interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	return err
}

func handleResponseError(r *http.Response, u string) error {
	if r.StatusCode == 500 {
		var e InternalServerErrorResponse
		f := json.NewDecoder(r.Body).Decode(e)
		if f != nil {
			return fmt.Errorf("%s: ListRequestAndDo for %s, also encountered an error "+
				"Unmarshalling the error response", r.Status, u)
		}
		return fmt.Errorf("%s: ListRequestAndDo for %s, with message %s", r.Status,
			u, e.Error.Message.Value)
	}

	if r.StatusCode == 400 || r.StatusCode == 401 || r.StatusCode == 403 ||
		r.StatusCode == 404 {
		return fmt.Errorf("%s: ListRequestAndDo for %s", r.Status, u)
	}

	return nil
}
