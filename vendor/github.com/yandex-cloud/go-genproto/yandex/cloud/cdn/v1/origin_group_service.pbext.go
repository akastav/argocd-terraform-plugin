// Code generated by protoc-gen-goext. DO NOT EDIT.

package cdn

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *GetOriginGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *GetOriginGroupRequest) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *ListOriginGroupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListOriginGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListOriginGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListOriginGroupsResponse) SetOriginGroups(v []*OriginGroup) {
	m.OriginGroups = v
}

func (m *ListOriginGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateOriginGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateOriginGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateOriginGroupRequest) SetUseNext(v *wrapperspb.BoolValue) {
	m.UseNext = v
}

func (m *CreateOriginGroupRequest) SetOrigins(v []*OriginParams) {
	m.Origins = v
}

func (m *CreateOriginGroupMetadata) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *UpdateOriginGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *UpdateOriginGroupRequest) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *UpdateOriginGroupRequest) SetGroupName(v *wrapperspb.StringValue) {
	m.GroupName = v
}

func (m *UpdateOriginGroupRequest) SetUseNext(v *wrapperspb.BoolValue) {
	m.UseNext = v
}

func (m *UpdateOriginGroupRequest) SetOrigins(v []*OriginParams) {
	m.Origins = v
}

func (m *UpdateOriginGroupMetadata) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *DeleteOriginGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *DeleteOriginGroupRequest) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *DeleteOriginGroupMetadata) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}