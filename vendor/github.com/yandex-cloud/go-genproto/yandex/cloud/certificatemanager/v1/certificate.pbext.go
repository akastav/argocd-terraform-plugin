// Code generated by protoc-gen-goext. DO NOT EDIT.

package certificatemanager

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Certificate) SetId(v string) {
	m.Id = v
}

func (m *Certificate) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Certificate) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Certificate) SetName(v string) {
	m.Name = v
}

func (m *Certificate) SetDescription(v string) {
	m.Description = v
}

func (m *Certificate) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Certificate) SetType(v CertificateType) {
	m.Type = v
}

func (m *Certificate) SetDomains(v []string) {
	m.Domains = v
}

func (m *Certificate) SetStatus(v Certificate_Status) {
	m.Status = v
}

func (m *Certificate) SetIssuer(v string) {
	m.Issuer = v
}

func (m *Certificate) SetSubject(v string) {
	m.Subject = v
}

func (m *Certificate) SetSerial(v string) {
	m.Serial = v
}

func (m *Certificate) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Certificate) SetIssuedAt(v *timestamppb.Timestamp) {
	m.IssuedAt = v
}

func (m *Certificate) SetNotAfter(v *timestamppb.Timestamp) {
	m.NotAfter = v
}

func (m *Certificate) SetNotBefore(v *timestamppb.Timestamp) {
	m.NotBefore = v
}

func (m *Certificate) SetChallenges(v []*Challenge) {
	m.Challenges = v
}

type Challenge_Challenge = isChallenge_Challenge

func (m *Challenge) SetChallenge(v Challenge_Challenge) {
	m.Challenge = v
}

func (m *Challenge) SetDomain(v string) {
	m.Domain = v
}

func (m *Challenge) SetType(v ChallengeType) {
	m.Type = v
}

func (m *Challenge) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Challenge) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Challenge) SetStatus(v Challenge_Status) {
	m.Status = v
}

func (m *Challenge) SetMessage(v string) {
	m.Message = v
}

func (m *Challenge) SetError(v string) {
	m.Error = v
}

func (m *Challenge) SetDnsChallenge(v *Challenge_DnsRecord) {
	m.Challenge = &Challenge_DnsChallenge{
		DnsChallenge: v,
	}
}

func (m *Challenge) SetHttpChallenge(v *Challenge_HttpFile) {
	m.Challenge = &Challenge_HttpChallenge{
		HttpChallenge: v,
	}
}

func (m *Challenge_DnsRecord) SetName(v string) {
	m.Name = v
}

func (m *Challenge_DnsRecord) SetType(v string) {
	m.Type = v
}

func (m *Challenge_DnsRecord) SetValue(v string) {
	m.Value = v
}

func (m *Challenge_HttpFile) SetUrl(v string) {
	m.Url = v
}

func (m *Challenge_HttpFile) SetContent(v string) {
	m.Content = v
}

func (m *Version) SetId(v string) {
	m.Id = v
}

func (m *Version) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *Version) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}
