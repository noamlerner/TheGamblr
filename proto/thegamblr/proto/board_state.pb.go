// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: board_state.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BoardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommunityCards represent the card available to everyone on the board. They will always be ordered -
	// Cards 0, 1 and 2 will be the flop. Card 3 will be the Turn and card 4 will be the River.
	CommunityCards []*Card `protobuf:"bytes,1,rep,name=community_cards,json=communityCards,proto3" json:"community_cards,omitempty"`
	// Provides info: How much money is in the pot.
	Pot uint64 `protobuf:"varint,2,opt,name=pot,proto3" json:"pot,omitempty"`
	// Enum representing what stage the current round is in.
	Stage Stage `protobuf:"varint,3,opt,name=stage,proto3,enum=thegamblr.Stage" json:"stage,omitempty"`
	// An index representing where the small blind button is. The small blind button is the index of the player that pays
	// the small blind and the one that is first to act on the Flop, Turn and River.
	SmallBlindButton uint32 `protobuf:"varint,4,opt,name=smallBlindButton,proto3" json:"smallBlindButton,omitempty"`
	// An ordered list of up to 8 players. Some entries may be nil if a seat is empty.
	Players []*PlayerState `protobuf:"bytes,5,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *BoardState) Reset() {
	*x = BoardState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardState) ProtoMessage() {}

func (x *BoardState) ProtoReflect() protoreflect.Message {
	mi := &file_board_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardState.ProtoReflect.Descriptor instead.
func (*BoardState) Descriptor() ([]byte, []int) {
	return file_board_state_proto_rawDescGZIP(), []int{0}
}

func (x *BoardState) GetCommunityCards() []*Card {
	if x != nil {
		return x.CommunityCards
	}
	return nil
}

func (x *BoardState) GetPot() uint64 {
	if x != nil {
		return x.Pot
	}
	return 0
}

func (x *BoardState) GetStage() Stage {
	if x != nil {
		return x.Stage
	}
	return Stage_PRE_FLOP
}

func (x *BoardState) GetSmallBlindButton() uint32 {
	if x != nil {
		return x.SmallBlindButton
	}
	return 0
}

func (x *BoardState) GetPlayers() []*PlayerState {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_board_state_proto protoreflect.FileDescriptor

var file_board_state_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x1a, 0x0a,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x0a,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x70, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x42,
	0x6c, 0x69, 0x6e, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68,
	0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_state_proto_rawDescOnce sync.Once
	file_board_state_proto_rawDescData = file_board_state_proto_rawDesc
)

func file_board_state_proto_rawDescGZIP() []byte {
	file_board_state_proto_rawDescOnce.Do(func() {
		file_board_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_state_proto_rawDescData)
	})
	return file_board_state_proto_rawDescData
}

var file_board_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_board_state_proto_goTypes = []interface{}{
	(*BoardState)(nil),  // 0: thegamblr.BoardState
	(*Card)(nil),        // 1: thegamblr.Card
	(Stage)(0),          // 2: thegamblr.Stage
	(*PlayerState)(nil), // 3: thegamblr.PlayerState
}
var file_board_state_proto_depIdxs = []int32{
	1, // 0: thegamblr.BoardState.community_cards:type_name -> thegamblr.Card
	2, // 1: thegamblr.BoardState.stage:type_name -> thegamblr.Stage
	3, // 2: thegamblr.BoardState.players:type_name -> thegamblr.PlayerState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_board_state_proto_init() }
func file_board_state_proto_init() {
	if File_board_state_proto != nil {
		return
	}
	file_card_proto_init()
	file_stage_proto_init()
	file_player_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_board_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_board_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_board_state_proto_goTypes,
		DependencyIndexes: file_board_state_proto_depIdxs,
		MessageInfos:      file_board_state_proto_msgTypes,
	}.Build()
	File_board_state_proto = out.File
	file_board_state_proto_rawDesc = nil
	file_board_state_proto_goTypes = nil
	file_board_state_proto_depIdxs = nil
}
