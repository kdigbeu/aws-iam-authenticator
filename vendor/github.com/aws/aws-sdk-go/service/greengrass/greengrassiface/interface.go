// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package greengrassiface provides an interface to enable mocking the AWS Greengrass service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package greengrassiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/greengrass"
)

// GreengrassAPI provides an interface to enable mocking the
// greengrass.Greengrass service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Greengrass.
//    func myFunc(svc greengrassiface.GreengrassAPI) bool {
//        // Make svc.AssociateRoleToGroup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := greengrass.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGreengrassClient struct {
//        greengrassiface.GreengrassAPI
//    }
//    func (m *mockGreengrassClient) AssociateRoleToGroup(input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGreengrassClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type GreengrassAPI interface {
	AssociateRoleToGroup(*greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error)
	AssociateRoleToGroupWithContext(aws.Context, *greengrass.AssociateRoleToGroupInput, ...request.Option) (*greengrass.AssociateRoleToGroupOutput, error)
	AssociateRoleToGroupRequest(*greengrass.AssociateRoleToGroupInput) (*request.Request, *greengrass.AssociateRoleToGroupOutput)

	AssociateServiceRoleToAccount(*greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error)
	AssociateServiceRoleToAccountWithContext(aws.Context, *greengrass.AssociateServiceRoleToAccountInput, ...request.Option) (*greengrass.AssociateServiceRoleToAccountOutput, error)
	AssociateServiceRoleToAccountRequest(*greengrass.AssociateServiceRoleToAccountInput) (*request.Request, *greengrass.AssociateServiceRoleToAccountOutput)

	CreateConnectorDefinition(*greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error)
	CreateConnectorDefinitionWithContext(aws.Context, *greengrass.CreateConnectorDefinitionInput, ...request.Option) (*greengrass.CreateConnectorDefinitionOutput, error)
	CreateConnectorDefinitionRequest(*greengrass.CreateConnectorDefinitionInput) (*request.Request, *greengrass.CreateConnectorDefinitionOutput)

	CreateConnectorDefinitionVersion(*greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error)
	CreateConnectorDefinitionVersionWithContext(aws.Context, *greengrass.CreateConnectorDefinitionVersionInput, ...request.Option) (*greengrass.CreateConnectorDefinitionVersionOutput, error)
	CreateConnectorDefinitionVersionRequest(*greengrass.CreateConnectorDefinitionVersionInput) (*request.Request, *greengrass.CreateConnectorDefinitionVersionOutput)

	CreateCoreDefinition(*greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error)
	CreateCoreDefinitionWithContext(aws.Context, *greengrass.CreateCoreDefinitionInput, ...request.Option) (*greengrass.CreateCoreDefinitionOutput, error)
	CreateCoreDefinitionRequest(*greengrass.CreateCoreDefinitionInput) (*request.Request, *greengrass.CreateCoreDefinitionOutput)

	CreateCoreDefinitionVersion(*greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error)
	CreateCoreDefinitionVersionWithContext(aws.Context, *greengrass.CreateCoreDefinitionVersionInput, ...request.Option) (*greengrass.CreateCoreDefinitionVersionOutput, error)
	CreateCoreDefinitionVersionRequest(*greengrass.CreateCoreDefinitionVersionInput) (*request.Request, *greengrass.CreateCoreDefinitionVersionOutput)

	CreateDeployment(*greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error)
	CreateDeploymentWithContext(aws.Context, *greengrass.CreateDeploymentInput, ...request.Option) (*greengrass.CreateDeploymentOutput, error)
	CreateDeploymentRequest(*greengrass.CreateDeploymentInput) (*request.Request, *greengrass.CreateDeploymentOutput)

	CreateDeviceDefinition(*greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error)
	CreateDeviceDefinitionWithContext(aws.Context, *greengrass.CreateDeviceDefinitionInput, ...request.Option) (*greengrass.CreateDeviceDefinitionOutput, error)
	CreateDeviceDefinitionRequest(*greengrass.CreateDeviceDefinitionInput) (*request.Request, *greengrass.CreateDeviceDefinitionOutput)

	CreateDeviceDefinitionVersion(*greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error)
	CreateDeviceDefinitionVersionWithContext(aws.Context, *greengrass.CreateDeviceDefinitionVersionInput, ...request.Option) (*greengrass.CreateDeviceDefinitionVersionOutput, error)
	CreateDeviceDefinitionVersionRequest(*greengrass.CreateDeviceDefinitionVersionInput) (*request.Request, *greengrass.CreateDeviceDefinitionVersionOutput)

	CreateFunctionDefinition(*greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error)
	CreateFunctionDefinitionWithContext(aws.Context, *greengrass.CreateFunctionDefinitionInput, ...request.Option) (*greengrass.CreateFunctionDefinitionOutput, error)
	CreateFunctionDefinitionRequest(*greengrass.CreateFunctionDefinitionInput) (*request.Request, *greengrass.CreateFunctionDefinitionOutput)

	CreateFunctionDefinitionVersion(*greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error)
	CreateFunctionDefinitionVersionWithContext(aws.Context, *greengrass.CreateFunctionDefinitionVersionInput, ...request.Option) (*greengrass.CreateFunctionDefinitionVersionOutput, error)
	CreateFunctionDefinitionVersionRequest(*greengrass.CreateFunctionDefinitionVersionInput) (*request.Request, *greengrass.CreateFunctionDefinitionVersionOutput)

	CreateGroup(*greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *greengrass.CreateGroupInput, ...request.Option) (*greengrass.CreateGroupOutput, error)
	CreateGroupRequest(*greengrass.CreateGroupInput) (*request.Request, *greengrass.CreateGroupOutput)

	CreateGroupCertificateAuthority(*greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error)
	CreateGroupCertificateAuthorityWithContext(aws.Context, *greengrass.CreateGroupCertificateAuthorityInput, ...request.Option) (*greengrass.CreateGroupCertificateAuthorityOutput, error)
	CreateGroupCertificateAuthorityRequest(*greengrass.CreateGroupCertificateAuthorityInput) (*request.Request, *greengrass.CreateGroupCertificateAuthorityOutput)

	CreateGroupVersion(*greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error)
	CreateGroupVersionWithContext(aws.Context, *greengrass.CreateGroupVersionInput, ...request.Option) (*greengrass.CreateGroupVersionOutput, error)
	CreateGroupVersionRequest(*greengrass.CreateGroupVersionInput) (*request.Request, *greengrass.CreateGroupVersionOutput)

	CreateLoggerDefinition(*greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error)
	CreateLoggerDefinitionWithContext(aws.Context, *greengrass.CreateLoggerDefinitionInput, ...request.Option) (*greengrass.CreateLoggerDefinitionOutput, error)
	CreateLoggerDefinitionRequest(*greengrass.CreateLoggerDefinitionInput) (*request.Request, *greengrass.CreateLoggerDefinitionOutput)

	CreateLoggerDefinitionVersion(*greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error)
	CreateLoggerDefinitionVersionWithContext(aws.Context, *greengrass.CreateLoggerDefinitionVersionInput, ...request.Option) (*greengrass.CreateLoggerDefinitionVersionOutput, error)
	CreateLoggerDefinitionVersionRequest(*greengrass.CreateLoggerDefinitionVersionInput) (*request.Request, *greengrass.CreateLoggerDefinitionVersionOutput)

	CreateResourceDefinition(*greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error)
	CreateResourceDefinitionWithContext(aws.Context, *greengrass.CreateResourceDefinitionInput, ...request.Option) (*greengrass.CreateResourceDefinitionOutput, error)
	CreateResourceDefinitionRequest(*greengrass.CreateResourceDefinitionInput) (*request.Request, *greengrass.CreateResourceDefinitionOutput)

	CreateResourceDefinitionVersion(*greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error)
	CreateResourceDefinitionVersionWithContext(aws.Context, *greengrass.CreateResourceDefinitionVersionInput, ...request.Option) (*greengrass.CreateResourceDefinitionVersionOutput, error)
	CreateResourceDefinitionVersionRequest(*greengrass.CreateResourceDefinitionVersionInput) (*request.Request, *greengrass.CreateResourceDefinitionVersionOutput)

	CreateSoftwareUpdateJob(*greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error)
	CreateSoftwareUpdateJobWithContext(aws.Context, *greengrass.CreateSoftwareUpdateJobInput, ...request.Option) (*greengrass.CreateSoftwareUpdateJobOutput, error)
	CreateSoftwareUpdateJobRequest(*greengrass.CreateSoftwareUpdateJobInput) (*request.Request, *greengrass.CreateSoftwareUpdateJobOutput)

	CreateSubscriptionDefinition(*greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error)
	CreateSubscriptionDefinitionWithContext(aws.Context, *greengrass.CreateSubscriptionDefinitionInput, ...request.Option) (*greengrass.CreateSubscriptionDefinitionOutput, error)
	CreateSubscriptionDefinitionRequest(*greengrass.CreateSubscriptionDefinitionInput) (*request.Request, *greengrass.CreateSubscriptionDefinitionOutput)

	CreateSubscriptionDefinitionVersion(*greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error)
	CreateSubscriptionDefinitionVersionWithContext(aws.Context, *greengrass.CreateSubscriptionDefinitionVersionInput, ...request.Option) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error)
	CreateSubscriptionDefinitionVersionRequest(*greengrass.CreateSubscriptionDefinitionVersionInput) (*request.Request, *greengrass.CreateSubscriptionDefinitionVersionOutput)

	DeleteConnectorDefinition(*greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error)
	DeleteConnectorDefinitionWithContext(aws.Context, *greengrass.DeleteConnectorDefinitionInput, ...request.Option) (*greengrass.DeleteConnectorDefinitionOutput, error)
	DeleteConnectorDefinitionRequest(*greengrass.DeleteConnectorDefinitionInput) (*request.Request, *greengrass.DeleteConnectorDefinitionOutput)

	DeleteCoreDefinition(*greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error)
	DeleteCoreDefinitionWithContext(aws.Context, *greengrass.DeleteCoreDefinitionInput, ...request.Option) (*greengrass.DeleteCoreDefinitionOutput, error)
	DeleteCoreDefinitionRequest(*greengrass.DeleteCoreDefinitionInput) (*request.Request, *greengrass.DeleteCoreDefinitionOutput)

	DeleteDeviceDefinition(*greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error)
	DeleteDeviceDefinitionWithContext(aws.Context, *greengrass.DeleteDeviceDefinitionInput, ...request.Option) (*greengrass.DeleteDeviceDefinitionOutput, error)
	DeleteDeviceDefinitionRequest(*greengrass.DeleteDeviceDefinitionInput) (*request.Request, *greengrass.DeleteDeviceDefinitionOutput)

	DeleteFunctionDefinition(*greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error)
	DeleteFunctionDefinitionWithContext(aws.Context, *greengrass.DeleteFunctionDefinitionInput, ...request.Option) (*greengrass.DeleteFunctionDefinitionOutput, error)
	DeleteFunctionDefinitionRequest(*greengrass.DeleteFunctionDefinitionInput) (*request.Request, *greengrass.DeleteFunctionDefinitionOutput)

	DeleteGroup(*greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *greengrass.DeleteGroupInput, ...request.Option) (*greengrass.DeleteGroupOutput, error)
	DeleteGroupRequest(*greengrass.DeleteGroupInput) (*request.Request, *greengrass.DeleteGroupOutput)

	DeleteLoggerDefinition(*greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error)
	DeleteLoggerDefinitionWithContext(aws.Context, *greengrass.DeleteLoggerDefinitionInput, ...request.Option) (*greengrass.DeleteLoggerDefinitionOutput, error)
	DeleteLoggerDefinitionRequest(*greengrass.DeleteLoggerDefinitionInput) (*request.Request, *greengrass.DeleteLoggerDefinitionOutput)

	DeleteResourceDefinition(*greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error)
	DeleteResourceDefinitionWithContext(aws.Context, *greengrass.DeleteResourceDefinitionInput, ...request.Option) (*greengrass.DeleteResourceDefinitionOutput, error)
	DeleteResourceDefinitionRequest(*greengrass.DeleteResourceDefinitionInput) (*request.Request, *greengrass.DeleteResourceDefinitionOutput)

	DeleteSubscriptionDefinition(*greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error)
	DeleteSubscriptionDefinitionWithContext(aws.Context, *greengrass.DeleteSubscriptionDefinitionInput, ...request.Option) (*greengrass.DeleteSubscriptionDefinitionOutput, error)
	DeleteSubscriptionDefinitionRequest(*greengrass.DeleteSubscriptionDefinitionInput) (*request.Request, *greengrass.DeleteSubscriptionDefinitionOutput)

	DisassociateRoleFromGroup(*greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error)
	DisassociateRoleFromGroupWithContext(aws.Context, *greengrass.DisassociateRoleFromGroupInput, ...request.Option) (*greengrass.DisassociateRoleFromGroupOutput, error)
	DisassociateRoleFromGroupRequest(*greengrass.DisassociateRoleFromGroupInput) (*request.Request, *greengrass.DisassociateRoleFromGroupOutput)

	DisassociateServiceRoleFromAccount(*greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error)
	DisassociateServiceRoleFromAccountWithContext(aws.Context, *greengrass.DisassociateServiceRoleFromAccountInput, ...request.Option) (*greengrass.DisassociateServiceRoleFromAccountOutput, error)
	DisassociateServiceRoleFromAccountRequest(*greengrass.DisassociateServiceRoleFromAccountInput) (*request.Request, *greengrass.DisassociateServiceRoleFromAccountOutput)

	GetAssociatedRole(*greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error)
	GetAssociatedRoleWithContext(aws.Context, *greengrass.GetAssociatedRoleInput, ...request.Option) (*greengrass.GetAssociatedRoleOutput, error)
	GetAssociatedRoleRequest(*greengrass.GetAssociatedRoleInput) (*request.Request, *greengrass.GetAssociatedRoleOutput)

	GetBulkDeploymentStatus(*greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error)
	GetBulkDeploymentStatusWithContext(aws.Context, *greengrass.GetBulkDeploymentStatusInput, ...request.Option) (*greengrass.GetBulkDeploymentStatusOutput, error)
	GetBulkDeploymentStatusRequest(*greengrass.GetBulkDeploymentStatusInput) (*request.Request, *greengrass.GetBulkDeploymentStatusOutput)

	GetConnectivityInfo(*greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error)
	GetConnectivityInfoWithContext(aws.Context, *greengrass.GetConnectivityInfoInput, ...request.Option) (*greengrass.GetConnectivityInfoOutput, error)
	GetConnectivityInfoRequest(*greengrass.GetConnectivityInfoInput) (*request.Request, *greengrass.GetConnectivityInfoOutput)

	GetConnectorDefinition(*greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error)
	GetConnectorDefinitionWithContext(aws.Context, *greengrass.GetConnectorDefinitionInput, ...request.Option) (*greengrass.GetConnectorDefinitionOutput, error)
	GetConnectorDefinitionRequest(*greengrass.GetConnectorDefinitionInput) (*request.Request, *greengrass.GetConnectorDefinitionOutput)

	GetConnectorDefinitionVersion(*greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error)
	GetConnectorDefinitionVersionWithContext(aws.Context, *greengrass.GetConnectorDefinitionVersionInput, ...request.Option) (*greengrass.GetConnectorDefinitionVersionOutput, error)
	GetConnectorDefinitionVersionRequest(*greengrass.GetConnectorDefinitionVersionInput) (*request.Request, *greengrass.GetConnectorDefinitionVersionOutput)

	GetCoreDefinition(*greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error)
	GetCoreDefinitionWithContext(aws.Context, *greengrass.GetCoreDefinitionInput, ...request.Option) (*greengrass.GetCoreDefinitionOutput, error)
	GetCoreDefinitionRequest(*greengrass.GetCoreDefinitionInput) (*request.Request, *greengrass.GetCoreDefinitionOutput)

	GetCoreDefinitionVersion(*greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error)
	GetCoreDefinitionVersionWithContext(aws.Context, *greengrass.GetCoreDefinitionVersionInput, ...request.Option) (*greengrass.GetCoreDefinitionVersionOutput, error)
	GetCoreDefinitionVersionRequest(*greengrass.GetCoreDefinitionVersionInput) (*request.Request, *greengrass.GetCoreDefinitionVersionOutput)

	GetDeploymentStatus(*greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error)
	GetDeploymentStatusWithContext(aws.Context, *greengrass.GetDeploymentStatusInput, ...request.Option) (*greengrass.GetDeploymentStatusOutput, error)
	GetDeploymentStatusRequest(*greengrass.GetDeploymentStatusInput) (*request.Request, *greengrass.GetDeploymentStatusOutput)

	GetDeviceDefinition(*greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error)
	GetDeviceDefinitionWithContext(aws.Context, *greengrass.GetDeviceDefinitionInput, ...request.Option) (*greengrass.GetDeviceDefinitionOutput, error)
	GetDeviceDefinitionRequest(*greengrass.GetDeviceDefinitionInput) (*request.Request, *greengrass.GetDeviceDefinitionOutput)

	GetDeviceDefinitionVersion(*greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error)
	GetDeviceDefinitionVersionWithContext(aws.Context, *greengrass.GetDeviceDefinitionVersionInput, ...request.Option) (*greengrass.GetDeviceDefinitionVersionOutput, error)
	GetDeviceDefinitionVersionRequest(*greengrass.GetDeviceDefinitionVersionInput) (*request.Request, *greengrass.GetDeviceDefinitionVersionOutput)

	GetFunctionDefinition(*greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error)
	GetFunctionDefinitionWithContext(aws.Context, *greengrass.GetFunctionDefinitionInput, ...request.Option) (*greengrass.GetFunctionDefinitionOutput, error)
	GetFunctionDefinitionRequest(*greengrass.GetFunctionDefinitionInput) (*request.Request, *greengrass.GetFunctionDefinitionOutput)

	GetFunctionDefinitionVersion(*greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error)
	GetFunctionDefinitionVersionWithContext(aws.Context, *greengrass.GetFunctionDefinitionVersionInput, ...request.Option) (*greengrass.GetFunctionDefinitionVersionOutput, error)
	GetFunctionDefinitionVersionRequest(*greengrass.GetFunctionDefinitionVersionInput) (*request.Request, *greengrass.GetFunctionDefinitionVersionOutput)

	GetGroup(*greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error)
	GetGroupWithContext(aws.Context, *greengrass.GetGroupInput, ...request.Option) (*greengrass.GetGroupOutput, error)
	GetGroupRequest(*greengrass.GetGroupInput) (*request.Request, *greengrass.GetGroupOutput)

	GetGroupCertificateAuthority(*greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error)
	GetGroupCertificateAuthorityWithContext(aws.Context, *greengrass.GetGroupCertificateAuthorityInput, ...request.Option) (*greengrass.GetGroupCertificateAuthorityOutput, error)
	GetGroupCertificateAuthorityRequest(*greengrass.GetGroupCertificateAuthorityInput) (*request.Request, *greengrass.GetGroupCertificateAuthorityOutput)

	GetGroupCertificateConfiguration(*greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error)
	GetGroupCertificateConfigurationWithContext(aws.Context, *greengrass.GetGroupCertificateConfigurationInput, ...request.Option) (*greengrass.GetGroupCertificateConfigurationOutput, error)
	GetGroupCertificateConfigurationRequest(*greengrass.GetGroupCertificateConfigurationInput) (*request.Request, *greengrass.GetGroupCertificateConfigurationOutput)

	GetGroupVersion(*greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error)
	GetGroupVersionWithContext(aws.Context, *greengrass.GetGroupVersionInput, ...request.Option) (*greengrass.GetGroupVersionOutput, error)
	GetGroupVersionRequest(*greengrass.GetGroupVersionInput) (*request.Request, *greengrass.GetGroupVersionOutput)

	GetLoggerDefinition(*greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error)
	GetLoggerDefinitionWithContext(aws.Context, *greengrass.GetLoggerDefinitionInput, ...request.Option) (*greengrass.GetLoggerDefinitionOutput, error)
	GetLoggerDefinitionRequest(*greengrass.GetLoggerDefinitionInput) (*request.Request, *greengrass.GetLoggerDefinitionOutput)

	GetLoggerDefinitionVersion(*greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error)
	GetLoggerDefinitionVersionWithContext(aws.Context, *greengrass.GetLoggerDefinitionVersionInput, ...request.Option) (*greengrass.GetLoggerDefinitionVersionOutput, error)
	GetLoggerDefinitionVersionRequest(*greengrass.GetLoggerDefinitionVersionInput) (*request.Request, *greengrass.GetLoggerDefinitionVersionOutput)

	GetResourceDefinition(*greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error)
	GetResourceDefinitionWithContext(aws.Context, *greengrass.GetResourceDefinitionInput, ...request.Option) (*greengrass.GetResourceDefinitionOutput, error)
	GetResourceDefinitionRequest(*greengrass.GetResourceDefinitionInput) (*request.Request, *greengrass.GetResourceDefinitionOutput)

	GetResourceDefinitionVersion(*greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error)
	GetResourceDefinitionVersionWithContext(aws.Context, *greengrass.GetResourceDefinitionVersionInput, ...request.Option) (*greengrass.GetResourceDefinitionVersionOutput, error)
	GetResourceDefinitionVersionRequest(*greengrass.GetResourceDefinitionVersionInput) (*request.Request, *greengrass.GetResourceDefinitionVersionOutput)

	GetServiceRoleForAccount(*greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error)
	GetServiceRoleForAccountWithContext(aws.Context, *greengrass.GetServiceRoleForAccountInput, ...request.Option) (*greengrass.GetServiceRoleForAccountOutput, error)
	GetServiceRoleForAccountRequest(*greengrass.GetServiceRoleForAccountInput) (*request.Request, *greengrass.GetServiceRoleForAccountOutput)

	GetSubscriptionDefinition(*greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error)
	GetSubscriptionDefinitionWithContext(aws.Context, *greengrass.GetSubscriptionDefinitionInput, ...request.Option) (*greengrass.GetSubscriptionDefinitionOutput, error)
	GetSubscriptionDefinitionRequest(*greengrass.GetSubscriptionDefinitionInput) (*request.Request, *greengrass.GetSubscriptionDefinitionOutput)

	GetSubscriptionDefinitionVersion(*greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error)
	GetSubscriptionDefinitionVersionWithContext(aws.Context, *greengrass.GetSubscriptionDefinitionVersionInput, ...request.Option) (*greengrass.GetSubscriptionDefinitionVersionOutput, error)
	GetSubscriptionDefinitionVersionRequest(*greengrass.GetSubscriptionDefinitionVersionInput) (*request.Request, *greengrass.GetSubscriptionDefinitionVersionOutput)

	ListBulkDeploymentDetailedReports(*greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error)
	ListBulkDeploymentDetailedReportsWithContext(aws.Context, *greengrass.ListBulkDeploymentDetailedReportsInput, ...request.Option) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error)
	ListBulkDeploymentDetailedReportsRequest(*greengrass.ListBulkDeploymentDetailedReportsInput) (*request.Request, *greengrass.ListBulkDeploymentDetailedReportsOutput)

	ListBulkDeployments(*greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error)
	ListBulkDeploymentsWithContext(aws.Context, *greengrass.ListBulkDeploymentsInput, ...request.Option) (*greengrass.ListBulkDeploymentsOutput, error)
	ListBulkDeploymentsRequest(*greengrass.ListBulkDeploymentsInput) (*request.Request, *greengrass.ListBulkDeploymentsOutput)

	ListConnectorDefinitionVersions(*greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error)
	ListConnectorDefinitionVersionsWithContext(aws.Context, *greengrass.ListConnectorDefinitionVersionsInput, ...request.Option) (*greengrass.ListConnectorDefinitionVersionsOutput, error)
	ListConnectorDefinitionVersionsRequest(*greengrass.ListConnectorDefinitionVersionsInput) (*request.Request, *greengrass.ListConnectorDefinitionVersionsOutput)

	ListConnectorDefinitions(*greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error)
	ListConnectorDefinitionsWithContext(aws.Context, *greengrass.ListConnectorDefinitionsInput, ...request.Option) (*greengrass.ListConnectorDefinitionsOutput, error)
	ListConnectorDefinitionsRequest(*greengrass.ListConnectorDefinitionsInput) (*request.Request, *greengrass.ListConnectorDefinitionsOutput)

	ListCoreDefinitionVersions(*greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error)
	ListCoreDefinitionVersionsWithContext(aws.Context, *greengrass.ListCoreDefinitionVersionsInput, ...request.Option) (*greengrass.ListCoreDefinitionVersionsOutput, error)
	ListCoreDefinitionVersionsRequest(*greengrass.ListCoreDefinitionVersionsInput) (*request.Request, *greengrass.ListCoreDefinitionVersionsOutput)

	ListCoreDefinitions(*greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error)
	ListCoreDefinitionsWithContext(aws.Context, *greengrass.ListCoreDefinitionsInput, ...request.Option) (*greengrass.ListCoreDefinitionsOutput, error)
	ListCoreDefinitionsRequest(*greengrass.ListCoreDefinitionsInput) (*request.Request, *greengrass.ListCoreDefinitionsOutput)

	ListDeployments(*greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error)
	ListDeploymentsWithContext(aws.Context, *greengrass.ListDeploymentsInput, ...request.Option) (*greengrass.ListDeploymentsOutput, error)
	ListDeploymentsRequest(*greengrass.ListDeploymentsInput) (*request.Request, *greengrass.ListDeploymentsOutput)

	ListDeviceDefinitionVersions(*greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error)
	ListDeviceDefinitionVersionsWithContext(aws.Context, *greengrass.ListDeviceDefinitionVersionsInput, ...request.Option) (*greengrass.ListDeviceDefinitionVersionsOutput, error)
	ListDeviceDefinitionVersionsRequest(*greengrass.ListDeviceDefinitionVersionsInput) (*request.Request, *greengrass.ListDeviceDefinitionVersionsOutput)

	ListDeviceDefinitions(*greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error)
	ListDeviceDefinitionsWithContext(aws.Context, *greengrass.ListDeviceDefinitionsInput, ...request.Option) (*greengrass.ListDeviceDefinitionsOutput, error)
	ListDeviceDefinitionsRequest(*greengrass.ListDeviceDefinitionsInput) (*request.Request, *greengrass.ListDeviceDefinitionsOutput)

	ListFunctionDefinitionVersions(*greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error)
	ListFunctionDefinitionVersionsWithContext(aws.Context, *greengrass.ListFunctionDefinitionVersionsInput, ...request.Option) (*greengrass.ListFunctionDefinitionVersionsOutput, error)
	ListFunctionDefinitionVersionsRequest(*greengrass.ListFunctionDefinitionVersionsInput) (*request.Request, *greengrass.ListFunctionDefinitionVersionsOutput)

	ListFunctionDefinitions(*greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error)
	ListFunctionDefinitionsWithContext(aws.Context, *greengrass.ListFunctionDefinitionsInput, ...request.Option) (*greengrass.ListFunctionDefinitionsOutput, error)
	ListFunctionDefinitionsRequest(*greengrass.ListFunctionDefinitionsInput) (*request.Request, *greengrass.ListFunctionDefinitionsOutput)

	ListGroupCertificateAuthorities(*greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error)
	ListGroupCertificateAuthoritiesWithContext(aws.Context, *greengrass.ListGroupCertificateAuthoritiesInput, ...request.Option) (*greengrass.ListGroupCertificateAuthoritiesOutput, error)
	ListGroupCertificateAuthoritiesRequest(*greengrass.ListGroupCertificateAuthoritiesInput) (*request.Request, *greengrass.ListGroupCertificateAuthoritiesOutput)

	ListGroupVersions(*greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error)
	ListGroupVersionsWithContext(aws.Context, *greengrass.ListGroupVersionsInput, ...request.Option) (*greengrass.ListGroupVersionsOutput, error)
	ListGroupVersionsRequest(*greengrass.ListGroupVersionsInput) (*request.Request, *greengrass.ListGroupVersionsOutput)

	ListGroups(*greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *greengrass.ListGroupsInput, ...request.Option) (*greengrass.ListGroupsOutput, error)
	ListGroupsRequest(*greengrass.ListGroupsInput) (*request.Request, *greengrass.ListGroupsOutput)

	ListLoggerDefinitionVersions(*greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error)
	ListLoggerDefinitionVersionsWithContext(aws.Context, *greengrass.ListLoggerDefinitionVersionsInput, ...request.Option) (*greengrass.ListLoggerDefinitionVersionsOutput, error)
	ListLoggerDefinitionVersionsRequest(*greengrass.ListLoggerDefinitionVersionsInput) (*request.Request, *greengrass.ListLoggerDefinitionVersionsOutput)

	ListLoggerDefinitions(*greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error)
	ListLoggerDefinitionsWithContext(aws.Context, *greengrass.ListLoggerDefinitionsInput, ...request.Option) (*greengrass.ListLoggerDefinitionsOutput, error)
	ListLoggerDefinitionsRequest(*greengrass.ListLoggerDefinitionsInput) (*request.Request, *greengrass.ListLoggerDefinitionsOutput)

	ListResourceDefinitionVersions(*greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error)
	ListResourceDefinitionVersionsWithContext(aws.Context, *greengrass.ListResourceDefinitionVersionsInput, ...request.Option) (*greengrass.ListResourceDefinitionVersionsOutput, error)
	ListResourceDefinitionVersionsRequest(*greengrass.ListResourceDefinitionVersionsInput) (*request.Request, *greengrass.ListResourceDefinitionVersionsOutput)

	ListResourceDefinitions(*greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error)
	ListResourceDefinitionsWithContext(aws.Context, *greengrass.ListResourceDefinitionsInput, ...request.Option) (*greengrass.ListResourceDefinitionsOutput, error)
	ListResourceDefinitionsRequest(*greengrass.ListResourceDefinitionsInput) (*request.Request, *greengrass.ListResourceDefinitionsOutput)

	ListSubscriptionDefinitionVersions(*greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error)
	ListSubscriptionDefinitionVersionsWithContext(aws.Context, *greengrass.ListSubscriptionDefinitionVersionsInput, ...request.Option) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error)
	ListSubscriptionDefinitionVersionsRequest(*greengrass.ListSubscriptionDefinitionVersionsInput) (*request.Request, *greengrass.ListSubscriptionDefinitionVersionsOutput)

	ListSubscriptionDefinitions(*greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error)
	ListSubscriptionDefinitionsWithContext(aws.Context, *greengrass.ListSubscriptionDefinitionsInput, ...request.Option) (*greengrass.ListSubscriptionDefinitionsOutput, error)
	ListSubscriptionDefinitionsRequest(*greengrass.ListSubscriptionDefinitionsInput) (*request.Request, *greengrass.ListSubscriptionDefinitionsOutput)

	ResetDeployments(*greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error)
	ResetDeploymentsWithContext(aws.Context, *greengrass.ResetDeploymentsInput, ...request.Option) (*greengrass.ResetDeploymentsOutput, error)
	ResetDeploymentsRequest(*greengrass.ResetDeploymentsInput) (*request.Request, *greengrass.ResetDeploymentsOutput)

	StartBulkDeployment(*greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error)
	StartBulkDeploymentWithContext(aws.Context, *greengrass.StartBulkDeploymentInput, ...request.Option) (*greengrass.StartBulkDeploymentOutput, error)
	StartBulkDeploymentRequest(*greengrass.StartBulkDeploymentInput) (*request.Request, *greengrass.StartBulkDeploymentOutput)

	StopBulkDeployment(*greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error)
	StopBulkDeploymentWithContext(aws.Context, *greengrass.StopBulkDeploymentInput, ...request.Option) (*greengrass.StopBulkDeploymentOutput, error)
	StopBulkDeploymentRequest(*greengrass.StopBulkDeploymentInput) (*request.Request, *greengrass.StopBulkDeploymentOutput)

	UpdateConnectivityInfo(*greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error)
	UpdateConnectivityInfoWithContext(aws.Context, *greengrass.UpdateConnectivityInfoInput, ...request.Option) (*greengrass.UpdateConnectivityInfoOutput, error)
	UpdateConnectivityInfoRequest(*greengrass.UpdateConnectivityInfoInput) (*request.Request, *greengrass.UpdateConnectivityInfoOutput)

	UpdateConnectorDefinition(*greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error)
	UpdateConnectorDefinitionWithContext(aws.Context, *greengrass.UpdateConnectorDefinitionInput, ...request.Option) (*greengrass.UpdateConnectorDefinitionOutput, error)
	UpdateConnectorDefinitionRequest(*greengrass.UpdateConnectorDefinitionInput) (*request.Request, *greengrass.UpdateConnectorDefinitionOutput)

	UpdateCoreDefinition(*greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error)
	UpdateCoreDefinitionWithContext(aws.Context, *greengrass.UpdateCoreDefinitionInput, ...request.Option) (*greengrass.UpdateCoreDefinitionOutput, error)
	UpdateCoreDefinitionRequest(*greengrass.UpdateCoreDefinitionInput) (*request.Request, *greengrass.UpdateCoreDefinitionOutput)

	UpdateDeviceDefinition(*greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error)
	UpdateDeviceDefinitionWithContext(aws.Context, *greengrass.UpdateDeviceDefinitionInput, ...request.Option) (*greengrass.UpdateDeviceDefinitionOutput, error)
	UpdateDeviceDefinitionRequest(*greengrass.UpdateDeviceDefinitionInput) (*request.Request, *greengrass.UpdateDeviceDefinitionOutput)

	UpdateFunctionDefinition(*greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error)
	UpdateFunctionDefinitionWithContext(aws.Context, *greengrass.UpdateFunctionDefinitionInput, ...request.Option) (*greengrass.UpdateFunctionDefinitionOutput, error)
	UpdateFunctionDefinitionRequest(*greengrass.UpdateFunctionDefinitionInput) (*request.Request, *greengrass.UpdateFunctionDefinitionOutput)

	UpdateGroup(*greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *greengrass.UpdateGroupInput, ...request.Option) (*greengrass.UpdateGroupOutput, error)
	UpdateGroupRequest(*greengrass.UpdateGroupInput) (*request.Request, *greengrass.UpdateGroupOutput)

	UpdateGroupCertificateConfiguration(*greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error)
	UpdateGroupCertificateConfigurationWithContext(aws.Context, *greengrass.UpdateGroupCertificateConfigurationInput, ...request.Option) (*greengrass.UpdateGroupCertificateConfigurationOutput, error)
	UpdateGroupCertificateConfigurationRequest(*greengrass.UpdateGroupCertificateConfigurationInput) (*request.Request, *greengrass.UpdateGroupCertificateConfigurationOutput)

	UpdateLoggerDefinition(*greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error)
	UpdateLoggerDefinitionWithContext(aws.Context, *greengrass.UpdateLoggerDefinitionInput, ...request.Option) (*greengrass.UpdateLoggerDefinitionOutput, error)
	UpdateLoggerDefinitionRequest(*greengrass.UpdateLoggerDefinitionInput) (*request.Request, *greengrass.UpdateLoggerDefinitionOutput)

	UpdateResourceDefinition(*greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error)
	UpdateResourceDefinitionWithContext(aws.Context, *greengrass.UpdateResourceDefinitionInput, ...request.Option) (*greengrass.UpdateResourceDefinitionOutput, error)
	UpdateResourceDefinitionRequest(*greengrass.UpdateResourceDefinitionInput) (*request.Request, *greengrass.UpdateResourceDefinitionOutput)

	UpdateSubscriptionDefinition(*greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error)
	UpdateSubscriptionDefinitionWithContext(aws.Context, *greengrass.UpdateSubscriptionDefinitionInput, ...request.Option) (*greengrass.UpdateSubscriptionDefinitionOutput, error)
	UpdateSubscriptionDefinitionRequest(*greengrass.UpdateSubscriptionDefinitionInput) (*request.Request, *greengrass.UpdateSubscriptionDefinitionOutput)
}

var _ GreengrassAPI = (*greengrass.Greengrass)(nil)
