// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of entitled applications.
func (c *Client) ListEntitledApplications(ctx context.Context, params *ListEntitledApplicationsInput, optFns ...func(*Options)) (*ListEntitledApplicationsOutput, error) {
	if params == nil {
		params = &ListEntitledApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitledApplications", params, optFns, c.addOperationListEntitledApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitledApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitledApplicationsInput struct {

	// The name of the entitlement.
	//
	// This member is required.
	EntitlementName *string

	// The name of the stack with which the entitlement is associated.
	//
	// This member is required.
	StackName *string

	// The maximum size of each page of results.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntitledApplicationsOutput struct {

	// The entitled applications.
	EntitledApplications []types.EntitledApplication

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntitledApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntitledApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntitledApplications{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListEntitledApplicationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitledApplications(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListEntitledApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "ListEntitledApplications",
	}
}
