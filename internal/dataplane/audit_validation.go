package dataplane

import (
	"errors"
	"strings"
)

var ErrAuditInvalid = errors.New("invalid mutation audit event")

func ValidateAuditEvent(e AuditEvent) error {
	if strings.TrimSpace(e.Action)=="" || strings.TrimSpace(e.Resource)=="" || strings.TrimSpace(e.RequestID)=="" || strings.TrimSpace(e.ApprovalID)=="" { return ErrAuditInvalid }
	for _, v := range []string{e.Action,e.Resource,e.Actor,e.RequestID,e.ApprovalID,e.Result} { if strings.ContainsAny(v,"\r\n\x00") { return ErrAuditInvalid } }
	return nil
}
