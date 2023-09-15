// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: casino.proto

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

type ReceiveUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReceiveUpdatesRequest) Reset() {
	*x = ReceiveUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveUpdatesRequest) ProtoMessage() {}

func (x *ReceiveUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveUpdatesRequest.ProtoReflect.Descriptor instead.
func (*ReceiveUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveUpdatesRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ReceiveUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the board. Only gets included once per stage. So you will get this once PreFlop, and then it
	// will be null until the Flop.
	BoardState *BoardState `protobuf:"bytes,1,opt,name=board_state,json=boardState,proto3" json:"board_state,omitempty"`
	// These will tell you all actions player's took since the last response you received.
	ActionUpdates []*Action `protobuf:"bytes,2,rep,name=action_updates,json=actionUpdates,proto3" json:"action_updates,omitempty"`
	// If this bool is true, it is your turn to act. You need to make a call to Act.
	IsMyAction bool `protobuf:"varint,3,opt,name=is_my_action,json=isMyAction,proto3" json:"is_my_action,omitempty"`
	// You will always get this when you get the PreFlop board_state. Otherwise this will be nil.
	MyHand []*Card `protobuf:"bytes,4,rep,name=my_hand,json=myHand,proto3" json:"my_hand,omitempty"`
}

func (x *ReceiveUpdatesResponse) Reset() {
	*x = ReceiveUpdatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveUpdatesResponse) ProtoMessage() {}

func (x *ReceiveUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveUpdatesResponse.ProtoReflect.Descriptor instead.
func (*ReceiveUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveUpdatesResponse) GetBoardState() *BoardState {
	if x != nil {
		return x.BoardState
	}
	return nil
}

func (x *ReceiveUpdatesResponse) GetActionUpdates() []*Action {
	if x != nil {
		return x.ActionUpdates
	}
	return nil
}

func (x *ReceiveUpdatesResponse) GetIsMyAction() bool {
	if x != nil {
		return x.IsMyAction
	}
	return false
}

func (x *ReceiveUpdatesResponse) GetMyHand() []*Card {
	if x != nil {
		return x.MyHand
	}
	return nil
}

type ActRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the token you recieved in the JoinGameRequest. If you do not have this, your request will be ignored.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// What action are you taking? If this is invalid you will be folded.
	ActionType ActionType `protobuf:"varint,2,opt,name=action_type,json=actionType,proto3,enum=thegamblr.ActionType" json:"action_type,omitempty"`
	// You only need to include this if you choose Raise and you want to Raise anything other than the min raise.
	// If this is less than the min raise, it will automatically be increased to the min raise. It will be ignored in
	// all other cases.
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ActRequest) Reset() {
	*x = ActRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActRequest) ProtoMessage() {}

func (x *ActRequest) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActRequest.ProtoReflect.Descriptor instead.
func (*ActRequest) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{2}
}

func (x *ActRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ActRequest) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_FOLD
}

func (x *ActRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ActResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActResponse) Reset() {
	*x = ActResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActResponse) ProtoMessage() {}

func (x *ActResponse) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActResponse.ProtoReflect.Descriptor instead.
func (*ActResponse) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{3}
}

type JoinGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Give your player a name.
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	// What game are you joining?
	GameId string `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *JoinGameRequest) Reset() {
	*x = JoinGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameRequest) ProtoMessage() {}

func (x *JoinGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameRequest.ProtoReflect.Descriptor instead.
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{4}
}

func (x *JoinGameRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *JoinGameRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type JoinGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keep this token for future requests.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// This is the seat you have been placed it.
	SeatNumber uint32 `protobuf:"varint,2,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	// Your ID, should match what was in your request unless there was a conflict.
	PlayerId string `protobuf:"bytes,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *JoinGameResponse) Reset() {
	*x = JoinGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameResponse) ProtoMessage() {}

func (x *JoinGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameResponse.ProtoReflect.Descriptor instead.
func (*JoinGameResponse) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{5}
}

func (x *JoinGameResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JoinGameResponse) GetSeatNumber() uint32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

func (x *JoinGameResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// how much is the small blind? big blind is 2x. Defaults to 5 if not set (0)
	SmallBlind uint64 `protobuf:"varint,1,opt,name=small_blind,json=smallBlind,proto3" json:"small_blind,omitempty"`
	// NumRounds defines how many rounds will be played in this game. If this is -1, we will play until there is only
	// one player left with chips. Defaults to 200 if not set (0).
	NumRounds uint64 `protobuf:"varint,2,opt,name=num_rounds,json=numRounds,proto3" json:"num_rounds,omitempty"`
	// how many chips does each player start with? Defaults to 1000 if not set (0).
	StartingStack uint64 `protobuf:"varint,3,opt,name=starting_stack,json=startingStack,proto3" json:"starting_stack,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGameRequest) GetSmallBlind() uint64 {
	if x != nil {
		return x.SmallBlind
	}
	return 0
}

func (x *CreateGameRequest) GetNumRounds() uint64 {
	if x != nil {
		return x.NumRounds
	}
	return 0
}

func (x *CreateGameRequest) GetStartingStack() uint64 {
	if x != nil {
		return x.StartingStack
	}
	return 0
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// You and others will need this to join the game.
	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{7}
}

func (x *CreateGameResponse) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type StartGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// token must belong to the person who created the game
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StartGameRequest) Reset() {
	*x = StartGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameRequest) ProtoMessage() {}

func (x *StartGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameRequest.ProtoReflect.Descriptor instead.
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{8}
}

func (x *StartGameRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StartGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGameResponse) Reset() {
	*x = StartGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_casino_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameResponse) ProtoMessage() {}

func (x *StartGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_casino_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameResponse.ProtoReflect.Descriptor instead.
func (*StartGameResponse) Descriptor() ([]byte, []int) {
	return file_casino_proto_rawDescGZIP(), []int{9}
}

var File_casino_proto protoreflect.FileDescriptor

var file_casino_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x61, 0x73, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x1a, 0x11, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c,
	0x72, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6d, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4d, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x06, 0x6d, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x22, 0x72,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d,
	0x62, 0x6c, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x47, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x10, 0x4a, 0x6f,
	0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x7a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x5f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x75,
	0x6d, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x22, 0x2d,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a,
	0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x02, 0x0a,
	0x06, 0x43, 0x61, 0x73, 0x69, 0x6e, 0x6f, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x68, 0x65,
	0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x20, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x41, 0x63, 0x74, 0x12, 0x15, 0x2e,
	0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72,
	0x2e, 0x41, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f,
	0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_casino_proto_rawDescOnce sync.Once
	file_casino_proto_rawDescData = file_casino_proto_rawDesc
)

func file_casino_proto_rawDescGZIP() []byte {
	file_casino_proto_rawDescOnce.Do(func() {
		file_casino_proto_rawDescData = protoimpl.X.CompressGZIP(file_casino_proto_rawDescData)
	})
	return file_casino_proto_rawDescData
}

var file_casino_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_casino_proto_goTypes = []interface{}{
	(*ReceiveUpdatesRequest)(nil),  // 0: thegamblr.ReceiveUpdatesRequest
	(*ReceiveUpdatesResponse)(nil), // 1: thegamblr.ReceiveUpdatesResponse
	(*ActRequest)(nil),             // 2: thegamblr.ActRequest
	(*ActResponse)(nil),            // 3: thegamblr.ActResponse
	(*JoinGameRequest)(nil),        // 4: thegamblr.JoinGameRequest
	(*JoinGameResponse)(nil),       // 5: thegamblr.JoinGameResponse
	(*CreateGameRequest)(nil),      // 6: thegamblr.CreateGameRequest
	(*CreateGameResponse)(nil),     // 7: thegamblr.CreateGameResponse
	(*StartGameRequest)(nil),       // 8: thegamblr.StartGameRequest
	(*StartGameResponse)(nil),      // 9: thegamblr.StartGameResponse
	(*BoardState)(nil),             // 10: thegamblr.BoardState
	(*Action)(nil),                 // 11: thegamblr.Action
	(*Card)(nil),                   // 12: thegamblr.Card
	(ActionType)(0),                // 13: thegamblr.ActionType
}
var file_casino_proto_depIdxs = []int32{
	10, // 0: thegamblr.ReceiveUpdatesResponse.board_state:type_name -> thegamblr.BoardState
	11, // 1: thegamblr.ReceiveUpdatesResponse.action_updates:type_name -> thegamblr.Action
	12, // 2: thegamblr.ReceiveUpdatesResponse.my_hand:type_name -> thegamblr.Card
	13, // 3: thegamblr.ActRequest.action_type:type_name -> thegamblr.ActionType
	6,  // 4: thegamblr.Casino.CreateGame:input_type -> thegamblr.CreateGameRequest
	4,  // 5: thegamblr.Casino.JoinGame:input_type -> thegamblr.JoinGameRequest
	8,  // 6: thegamblr.Casino.StartGame:input_type -> thegamblr.StartGameRequest
	0,  // 7: thegamblr.Casino.ReceiveUpdates:input_type -> thegamblr.ReceiveUpdatesRequest
	2,  // 8: thegamblr.Casino.Act:input_type -> thegamblr.ActRequest
	7,  // 9: thegamblr.Casino.CreateGame:output_type -> thegamblr.CreateGameResponse
	5,  // 10: thegamblr.Casino.JoinGame:output_type -> thegamblr.JoinGameResponse
	9,  // 11: thegamblr.Casino.StartGame:output_type -> thegamblr.StartGameResponse
	1,  // 12: thegamblr.Casino.ReceiveUpdates:output_type -> thegamblr.ReceiveUpdatesResponse
	3,  // 13: thegamblr.Casino.Act:output_type -> thegamblr.ActResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_casino_proto_init() }
func file_casino_proto_init() {
	if File_casino_proto != nil {
		return
	}
	file_board_state_proto_init()
	file_action_proto_init()
	file_card_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_casino_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveUpdatesRequest); i {
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
		file_casino_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveUpdatesResponse); i {
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
		file_casino_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActRequest); i {
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
		file_casino_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActResponse); i {
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
		file_casino_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameRequest); i {
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
		file_casino_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameResponse); i {
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
		file_casino_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameRequest); i {
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
		file_casino_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameResponse); i {
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
		file_casino_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameRequest); i {
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
		file_casino_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameResponse); i {
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
			RawDescriptor: file_casino_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_casino_proto_goTypes,
		DependencyIndexes: file_casino_proto_depIdxs,
		MessageInfos:      file_casino_proto_msgTypes,
	}.Build()
	File_casino_proto = out.File
	file_casino_proto_rawDesc = nil
	file_casino_proto_goTypes = nil
	file_casino_proto_depIdxs = nil
}
