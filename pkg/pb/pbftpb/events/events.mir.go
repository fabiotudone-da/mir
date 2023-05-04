package pbftpbevents

import (
	types4 "github.com/filecoin-project/mir/pkg/orderers/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types2 "github.com/filecoin-project/mir/pkg/pb/ordererpb/types"
	types3 "github.com/filecoin-project/mir/pkg/pb/pbftpb/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

func ProposeTimeout(destModule types.ModuleID, proposeTimeout uint64) *types1.Event {
	return &types1.Event{
		DestModule: destModule,
		Type: &types1.Event_Orderer{
			Orderer: &types2.Event{
				Type: &types2.Event_Pbft{
					Pbft: &types3.Event{
						Type: &types3.Event_ProposeTimeout{
							ProposeTimeout: &types3.ProposeTimeout{
								ProposeTimeout: proposeTimeout,
							},
						},
					},
				},
			},
		},
	}
}

func ViewChangeSNTimeout(destModule types.ModuleID, view types4.ViewNr, numCommitted uint64) *types1.Event {
	return &types1.Event{
		DestModule: destModule,
		Type: &types1.Event_Orderer{
			Orderer: &types2.Event{
				Type: &types2.Event_Pbft{
					Pbft: &types3.Event{
						Type: &types3.Event_ViewChangeSnTimeout{
							ViewChangeSnTimeout: &types3.ViewChangeSNTimeout{
								View:         view,
								NumCommitted: numCommitted,
							},
						},
					},
				},
			},
		},
	}
}

func ViewChangeSegTimeout(destModule types.ModuleID, viewChangeSegTimeout uint64) *types1.Event {
	return &types1.Event{
		DestModule: destModule,
		Type: &types1.Event_Orderer{
			Orderer: &types2.Event{
				Type: &types2.Event_Pbft{
					Pbft: &types3.Event{
						Type: &types3.Event_ViewChangeSegTimeout{
							ViewChangeSegTimeout: &types3.ViewChangeSegTimeout{
								ViewChangeSegTimeout: viewChangeSegTimeout,
							},
						},
					},
				},
			},
		},
	}
}
