// Code generated by protoc-gen-goext. DO NOT EDIT.

package resourcemanager

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Cloud) SetId(v string) {
	m.Id = v
}

func (m *Cloud) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Cloud) SetName(v string) {
	m.Name = v
}

func (m *Cloud) SetDescription(v string) {
	m.Description = v
}

func (m *Cloud) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *Cloud) SetLabels(v map[string]string) {
	m.Labels = v
}