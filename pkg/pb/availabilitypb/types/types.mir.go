package availabilitypbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types1 "github.com/filecoin-project/mir/codegen/model/types"
	availabilitypb "github.com/filecoin-project/mir/pkg/pb/availabilitypb"
	types5 "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb/types"
	types3 "github.com/filecoin-project/mir/pkg/pb/contextstorepb/types"
	types4 "github.com/filecoin-project/mir/pkg/pb/dslpb/types"
	trantorpb "github.com/filecoin-project/mir/pkg/pb/trantorpb"
	types "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	types2 "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() availabilitypb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb availabilitypb.Event_Type) Event_Type {
	switch pb := pb.(type) {
	case *availabilitypb.Event_RequestCert:
		return &Event_RequestCert{RequestCert: RequestCertFromPb(pb.RequestCert)}
	case *availabilitypb.Event_NewCert:
		return &Event_NewCert{NewCert: NewCertFromPb(pb.NewCert)}
	case *availabilitypb.Event_VerifyCert:
		return &Event_VerifyCert{VerifyCert: VerifyCertFromPb(pb.VerifyCert)}
	case *availabilitypb.Event_CertVerified:
		return &Event_CertVerified{CertVerified: CertVerifiedFromPb(pb.CertVerified)}
	case *availabilitypb.Event_RequestTransactions:
		return &Event_RequestTransactions{RequestTransactions: RequestTransactionsFromPb(pb.RequestTransactions)}
	case *availabilitypb.Event_ProvideTransactions:
		return &Event_ProvideTransactions{ProvideTransactions: ProvideTransactionsFromPb(pb.ProvideTransactions)}
	case *availabilitypb.Event_ComputeCert:
		return &Event_ComputeCert{ComputeCert: ComputeCertFromPb(pb.ComputeCert)}
	}
	return nil
}

type Event_RequestCert struct {
	RequestCert *RequestCert
}

func (*Event_RequestCert) isEvent_Type() {}

func (w *Event_RequestCert) Unwrap() *RequestCert {
	return w.RequestCert
}

func (w *Event_RequestCert) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_RequestCert{RequestCert: (w.RequestCert).Pb()}
}

func (*Event_RequestCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_RequestCert]()}
}

type Event_NewCert struct {
	NewCert *NewCert
}

func (*Event_NewCert) isEvent_Type() {}

func (w *Event_NewCert) Unwrap() *NewCert {
	return w.NewCert
}

func (w *Event_NewCert) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_NewCert{NewCert: (w.NewCert).Pb()}
}

func (*Event_NewCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_NewCert]()}
}

type Event_VerifyCert struct {
	VerifyCert *VerifyCert
}

func (*Event_VerifyCert) isEvent_Type() {}

func (w *Event_VerifyCert) Unwrap() *VerifyCert {
	return w.VerifyCert
}

func (w *Event_VerifyCert) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_VerifyCert{VerifyCert: (w.VerifyCert).Pb()}
}

func (*Event_VerifyCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_VerifyCert]()}
}

type Event_CertVerified struct {
	CertVerified *CertVerified
}

func (*Event_CertVerified) isEvent_Type() {}

func (w *Event_CertVerified) Unwrap() *CertVerified {
	return w.CertVerified
}

func (w *Event_CertVerified) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_CertVerified{CertVerified: (w.CertVerified).Pb()}
}

func (*Event_CertVerified) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_CertVerified]()}
}

type Event_RequestTransactions struct {
	RequestTransactions *RequestTransactions
}

func (*Event_RequestTransactions) isEvent_Type() {}

func (w *Event_RequestTransactions) Unwrap() *RequestTransactions {
	return w.RequestTransactions
}

func (w *Event_RequestTransactions) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_RequestTransactions{RequestTransactions: (w.RequestTransactions).Pb()}
}

func (*Event_RequestTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_RequestTransactions]()}
}

type Event_ProvideTransactions struct {
	ProvideTransactions *ProvideTransactions
}

func (*Event_ProvideTransactions) isEvent_Type() {}

func (w *Event_ProvideTransactions) Unwrap() *ProvideTransactions {
	return w.ProvideTransactions
}

func (w *Event_ProvideTransactions) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_ProvideTransactions{ProvideTransactions: (w.ProvideTransactions).Pb()}
}

func (*Event_ProvideTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_ProvideTransactions]()}
}

type Event_ComputeCert struct {
	ComputeCert *ComputeCert
}

func (*Event_ComputeCert) isEvent_Type() {}

func (w *Event_ComputeCert) Unwrap() *ComputeCert {
	return w.ComputeCert
}

func (w *Event_ComputeCert) Pb() availabilitypb.Event_Type {
	return &availabilitypb.Event_ComputeCert{ComputeCert: (w.ComputeCert).Pb()}
}

func (*Event_ComputeCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event_ComputeCert]()}
}

func EventFromPb(pb *availabilitypb.Event) *Event {
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *availabilitypb.Event {
	return &availabilitypb.Event{
		Type: (m.Type).Pb(),
	}
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Event]()}
}

type RequestCert struct {
	Origin *RequestCertOrigin
}

func RequestCertFromPb(pb *availabilitypb.RequestCert) *RequestCert {
	return &RequestCert{
		Origin: RequestCertOriginFromPb(pb.Origin),
	}
}

func (m *RequestCert) Pb() *availabilitypb.RequestCert {
	return &availabilitypb.RequestCert{
		Origin: (m.Origin).Pb(),
	}
}

func (*RequestCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestCert]()}
}

type NewCert struct {
	Cert   *Cert
	Origin *RequestCertOrigin
}

func NewCertFromPb(pb *availabilitypb.NewCert) *NewCert {
	return &NewCert{
		Cert:   CertFromPb(pb.Cert),
		Origin: RequestCertOriginFromPb(pb.Origin),
	}
}

func (m *NewCert) Pb() *availabilitypb.NewCert {
	return &availabilitypb.NewCert{
		Cert:   (m.Cert).Pb(),
		Origin: (m.Origin).Pb(),
	}
}

func (*NewCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.NewCert]()}
}

type VerifyCert struct {
	Cert   *Cert
	Origin *VerifyCertOrigin
}

func VerifyCertFromPb(pb *availabilitypb.VerifyCert) *VerifyCert {
	return &VerifyCert{
		Cert:   CertFromPb(pb.Cert),
		Origin: VerifyCertOriginFromPb(pb.Origin),
	}
}

func (m *VerifyCert) Pb() *availabilitypb.VerifyCert {
	return &availabilitypb.VerifyCert{
		Cert:   (m.Cert).Pb(),
		Origin: (m.Origin).Pb(),
	}
}

func (*VerifyCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.VerifyCert]()}
}

type CertVerified struct {
	Valid  bool
	Err    string
	Origin *VerifyCertOrigin
}

func CertVerifiedFromPb(pb *availabilitypb.CertVerified) *CertVerified {
	return &CertVerified{
		Valid:  pb.Valid,
		Err:    pb.Err,
		Origin: VerifyCertOriginFromPb(pb.Origin),
	}
}

func (m *CertVerified) Pb() *availabilitypb.CertVerified {
	return &availabilitypb.CertVerified{
		Valid:  m.Valid,
		Err:    m.Err,
		Origin: (m.Origin).Pb(),
	}
}

func (*CertVerified) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.CertVerified]()}
}

type RequestTransactions struct {
	Cert   *Cert
	Origin *RequestTransactionsOrigin
}

func RequestTransactionsFromPb(pb *availabilitypb.RequestTransactions) *RequestTransactions {
	return &RequestTransactions{
		Cert:   CertFromPb(pb.Cert),
		Origin: RequestTransactionsOriginFromPb(pb.Origin),
	}
}

func (m *RequestTransactions) Pb() *availabilitypb.RequestTransactions {
	return &availabilitypb.RequestTransactions{
		Cert:   (m.Cert).Pb(),
		Origin: (m.Origin).Pb(),
	}
}

func (*RequestTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestTransactions]()}
}

type ProvideTransactions struct {
	Txs    []*types.Transaction
	Origin *RequestTransactionsOrigin
}

func ProvideTransactionsFromPb(pb *availabilitypb.ProvideTransactions) *ProvideTransactions {
	return &ProvideTransactions{
		Txs: types1.ConvertSlice(pb.Txs, func(t *trantorpb.Transaction) *types.Transaction {
			return types.TransactionFromPb(t)
		}),
		Origin: RequestTransactionsOriginFromPb(pb.Origin),
	}
}

func (m *ProvideTransactions) Pb() *availabilitypb.ProvideTransactions {
	return &availabilitypb.ProvideTransactions{
		Txs: types1.ConvertSlice(m.Txs, func(t *types.Transaction) *trantorpb.Transaction {
			return (t).Pb()
		}),
		Origin: (m.Origin).Pb(),
	}
}

func (*ProvideTransactions) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.ProvideTransactions]()}
}

type RequestCertOrigin struct {
	Module types2.ModuleID
	Type   RequestCertOrigin_Type
}

type RequestCertOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestCertOrigin_Type()
	Pb() availabilitypb.RequestCertOrigin_Type
}

type RequestCertOrigin_TypeWrapper[T any] interface {
	RequestCertOrigin_Type
	Unwrap() *T
}

func RequestCertOrigin_TypeFromPb(pb availabilitypb.RequestCertOrigin_Type) RequestCertOrigin_Type {
	switch pb := pb.(type) {
	case *availabilitypb.RequestCertOrigin_ContextStore:
		return &RequestCertOrigin_ContextStore{ContextStore: types3.OriginFromPb(pb.ContextStore)}
	case *availabilitypb.RequestCertOrigin_Dsl:
		return &RequestCertOrigin_Dsl{Dsl: types4.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestCertOrigin_ContextStore struct {
	ContextStore *types3.Origin
}

func (*RequestCertOrigin_ContextStore) isRequestCertOrigin_Type() {}

func (w *RequestCertOrigin_ContextStore) Unwrap() *types3.Origin {
	return w.ContextStore
}

func (w *RequestCertOrigin_ContextStore) Pb() availabilitypb.RequestCertOrigin_Type {
	return &availabilitypb.RequestCertOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestCertOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestCertOrigin_ContextStore]()}
}

type RequestCertOrigin_Dsl struct {
	Dsl *types4.Origin
}

func (*RequestCertOrigin_Dsl) isRequestCertOrigin_Type() {}

func (w *RequestCertOrigin_Dsl) Unwrap() *types4.Origin {
	return w.Dsl
}

func (w *RequestCertOrigin_Dsl) Pb() availabilitypb.RequestCertOrigin_Type {
	return &availabilitypb.RequestCertOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestCertOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestCertOrigin_Dsl]()}
}

func RequestCertOriginFromPb(pb *availabilitypb.RequestCertOrigin) *RequestCertOrigin {
	return &RequestCertOrigin{
		Module: (types2.ModuleID)(pb.Module),
		Type:   RequestCertOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestCertOrigin) Pb() *availabilitypb.RequestCertOrigin {
	return &availabilitypb.RequestCertOrigin{
		Module: (string)(m.Module),
		Type:   (m.Type).Pb(),
	}
}

func (*RequestCertOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestCertOrigin]()}
}

type RequestTransactionsOrigin struct {
	Module types2.ModuleID
	Type   RequestTransactionsOrigin_Type
}

type RequestTransactionsOrigin_Type interface {
	mirreflect.GeneratedType
	isRequestTransactionsOrigin_Type()
	Pb() availabilitypb.RequestTransactionsOrigin_Type
}

type RequestTransactionsOrigin_TypeWrapper[T any] interface {
	RequestTransactionsOrigin_Type
	Unwrap() *T
}

func RequestTransactionsOrigin_TypeFromPb(pb availabilitypb.RequestTransactionsOrigin_Type) RequestTransactionsOrigin_Type {
	switch pb := pb.(type) {
	case *availabilitypb.RequestTransactionsOrigin_ContextStore:
		return &RequestTransactionsOrigin_ContextStore{ContextStore: types3.OriginFromPb(pb.ContextStore)}
	case *availabilitypb.RequestTransactionsOrigin_Dsl:
		return &RequestTransactionsOrigin_Dsl{Dsl: types4.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type RequestTransactionsOrigin_ContextStore struct {
	ContextStore *types3.Origin
}

func (*RequestTransactionsOrigin_ContextStore) isRequestTransactionsOrigin_Type() {}

func (w *RequestTransactionsOrigin_ContextStore) Unwrap() *types3.Origin {
	return w.ContextStore
}

func (w *RequestTransactionsOrigin_ContextStore) Pb() availabilitypb.RequestTransactionsOrigin_Type {
	return &availabilitypb.RequestTransactionsOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*RequestTransactionsOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestTransactionsOrigin_ContextStore]()}
}

type RequestTransactionsOrigin_Dsl struct {
	Dsl *types4.Origin
}

func (*RequestTransactionsOrigin_Dsl) isRequestTransactionsOrigin_Type() {}

func (w *RequestTransactionsOrigin_Dsl) Unwrap() *types4.Origin {
	return w.Dsl
}

func (w *RequestTransactionsOrigin_Dsl) Pb() availabilitypb.RequestTransactionsOrigin_Type {
	return &availabilitypb.RequestTransactionsOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*RequestTransactionsOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestTransactionsOrigin_Dsl]()}
}

func RequestTransactionsOriginFromPb(pb *availabilitypb.RequestTransactionsOrigin) *RequestTransactionsOrigin {
	return &RequestTransactionsOrigin{
		Module: (types2.ModuleID)(pb.Module),
		Type:   RequestTransactionsOrigin_TypeFromPb(pb.Type),
	}
}

func (m *RequestTransactionsOrigin) Pb() *availabilitypb.RequestTransactionsOrigin {
	return &availabilitypb.RequestTransactionsOrigin{
		Module: (string)(m.Module),
		Type:   (m.Type).Pb(),
	}
}

func (*RequestTransactionsOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.RequestTransactionsOrigin]()}
}

type VerifyCertOrigin struct {
	Module types2.ModuleID
	Type   VerifyCertOrigin_Type
}

type VerifyCertOrigin_Type interface {
	mirreflect.GeneratedType
	isVerifyCertOrigin_Type()
	Pb() availabilitypb.VerifyCertOrigin_Type
}

type VerifyCertOrigin_TypeWrapper[T any] interface {
	VerifyCertOrigin_Type
	Unwrap() *T
}

func VerifyCertOrigin_TypeFromPb(pb availabilitypb.VerifyCertOrigin_Type) VerifyCertOrigin_Type {
	switch pb := pb.(type) {
	case *availabilitypb.VerifyCertOrigin_ContextStore:
		return &VerifyCertOrigin_ContextStore{ContextStore: types3.OriginFromPb(pb.ContextStore)}
	case *availabilitypb.VerifyCertOrigin_Dsl:
		return &VerifyCertOrigin_Dsl{Dsl: types4.OriginFromPb(pb.Dsl)}
	}
	return nil
}

type VerifyCertOrigin_ContextStore struct {
	ContextStore *types3.Origin
}

func (*VerifyCertOrigin_ContextStore) isVerifyCertOrigin_Type() {}

func (w *VerifyCertOrigin_ContextStore) Unwrap() *types3.Origin {
	return w.ContextStore
}

func (w *VerifyCertOrigin_ContextStore) Pb() availabilitypb.VerifyCertOrigin_Type {
	return &availabilitypb.VerifyCertOrigin_ContextStore{ContextStore: (w.ContextStore).Pb()}
}

func (*VerifyCertOrigin_ContextStore) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.VerifyCertOrigin_ContextStore]()}
}

type VerifyCertOrigin_Dsl struct {
	Dsl *types4.Origin
}

func (*VerifyCertOrigin_Dsl) isVerifyCertOrigin_Type() {}

func (w *VerifyCertOrigin_Dsl) Unwrap() *types4.Origin {
	return w.Dsl
}

func (w *VerifyCertOrigin_Dsl) Pb() availabilitypb.VerifyCertOrigin_Type {
	return &availabilitypb.VerifyCertOrigin_Dsl{Dsl: (w.Dsl).Pb()}
}

func (*VerifyCertOrigin_Dsl) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.VerifyCertOrigin_Dsl]()}
}

func VerifyCertOriginFromPb(pb *availabilitypb.VerifyCertOrigin) *VerifyCertOrigin {
	return &VerifyCertOrigin{
		Module: (types2.ModuleID)(pb.Module),
		Type:   VerifyCertOrigin_TypeFromPb(pb.Type),
	}
}

func (m *VerifyCertOrigin) Pb() *availabilitypb.VerifyCertOrigin {
	return &availabilitypb.VerifyCertOrigin{
		Module: (string)(m.Module),
		Type:   (m.Type).Pb(),
	}
}

func (*VerifyCertOrigin) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.VerifyCertOrigin]()}
}

type Cert struct {
	Type Cert_Type
}

type Cert_Type interface {
	mirreflect.GeneratedType
	isCert_Type()
	Pb() availabilitypb.Cert_Type
}

type Cert_TypeWrapper[T any] interface {
	Cert_Type
	Unwrap() *T
}

func Cert_TypeFromPb(pb availabilitypb.Cert_Type) Cert_Type {
	switch pb := pb.(type) {
	case *availabilitypb.Cert_Mscs:
		return &Cert_Mscs{Mscs: types5.CertsFromPb(pb.Mscs)}
	}
	return nil
}

type Cert_Mscs struct {
	Mscs *types5.Certs
}

func (*Cert_Mscs) isCert_Type() {}

func (w *Cert_Mscs) Unwrap() *types5.Certs {
	return w.Mscs
}

func (w *Cert_Mscs) Pb() availabilitypb.Cert_Type {
	return &availabilitypb.Cert_Mscs{Mscs: (w.Mscs).Pb()}
}

func (*Cert_Mscs) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Cert_Mscs]()}
}

func CertFromPb(pb *availabilitypb.Cert) *Cert {
	return &Cert{
		Type: Cert_TypeFromPb(pb.Type),
	}
}

func (m *Cert) Pb() *availabilitypb.Cert {
	return &availabilitypb.Cert{
		Type: (m.Type).Pb(),
	}
}

func (*Cert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.Cert]()}
}

type ComputeCert struct{}

func ComputeCertFromPb(pb *availabilitypb.ComputeCert) *ComputeCert {
	return &ComputeCert{}
}

func (m *ComputeCert) Pb() *availabilitypb.ComputeCert {
	return &availabilitypb.ComputeCert{}
}

func (*ComputeCert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*availabilitypb.ComputeCert]()}
}
