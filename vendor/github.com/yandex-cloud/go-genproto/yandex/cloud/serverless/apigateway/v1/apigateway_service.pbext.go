// Code generated by protoc-gen-goext. DO NOT EDIT.

package apigateway

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetApiGatewayRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *ListApiGatewayRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListApiGatewayRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListApiGatewayRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListApiGatewayRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListApiGatewayResponse) SetApiGateways(v []*ApiGateway) {
	m.ApiGateways = v
}

func (m *ListApiGatewayResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateApiGatewayRequest_Spec = isCreateApiGatewayRequest_Spec

func (m *CreateApiGatewayRequest) SetSpec(v CreateApiGatewayRequest_Spec) {
	m.Spec = v
}

func (m *CreateApiGatewayRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateApiGatewayRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateApiGatewayRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateApiGatewayRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateApiGatewayRequest) SetOpenapiSpec(v string) {
	m.Spec = &CreateApiGatewayRequest_OpenapiSpec{
		OpenapiSpec: v,
	}
}

func (m *CreateApiGatewayRequest) SetConnectivity(v *Connectivity) {
	m.Connectivity = v
}

type UpdateApiGatewayRequest_Spec = isUpdateApiGatewayRequest_Spec

func (m *UpdateApiGatewayRequest) SetSpec(v UpdateApiGatewayRequest_Spec) {
	m.Spec = v
}

func (m *UpdateApiGatewayRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *UpdateApiGatewayRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateApiGatewayRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateApiGatewayRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateApiGatewayRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateApiGatewayRequest) SetOpenapiSpec(v string) {
	m.Spec = &UpdateApiGatewayRequest_OpenapiSpec{
		OpenapiSpec: v,
	}
}

func (m *UpdateApiGatewayRequest) SetConnectivity(v *Connectivity) {
	m.Connectivity = v
}

func (m *DeleteApiGatewayRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *AddDomainRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *AddDomainRequest) SetDomainId(v string) {
	m.DomainId = v
}

func (m *RemoveDomainRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *RemoveDomainRequest) SetDomainId(v string) {
	m.DomainId = v
}

func (m *CreateApiGatewayMetadata) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *UpdateApiGatewayMetadata) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *DeleteApiGatewayMetadata) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *AddDomainMetadata) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *AddDomainMetadata) SetDomainId(v string) {
	m.DomainId = v
}

func (m *RemoveDomainMetadata) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *RemoveDomainMetadata) SetDomainId(v string) {
	m.DomainId = v
}

func (m *ListOperationsRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *ListOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetOpenapiSpecRequest) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *GetOpenapiSpecRequest) SetFormat(v GetOpenapiSpecRequest_Format) {
	m.Format = v
}

func (m *GetOpenapiSpecResponse) SetApiGatewayId(v string) {
	m.ApiGatewayId = v
}

func (m *GetOpenapiSpecResponse) SetOpenapiSpec(v string) {
	m.OpenapiSpec = v
}