# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: board_state.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import card_pb2 as card__pb2
import stage_pb2 as stage__pb2
import player_state_pb2 as player__state__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11\x62oard_state.proto\x12\tthegamblr\x1a\ncard.proto\x1a\x0bstage.proto\x1a\x12player_state.proto\"\xa7\x01\n\nBoardState\x12(\n\x0f\x63ommunity_cards\x18\x01 \x03(\x0b\x32\x0f.thegamblr.Card\x12\x0b\n\x03pot\x18\x02 \x01(\x04\x12\x1f\n\x05stage\x18\x03 \x01(\x0e\x32\x10.thegamblr.Stage\x12\x18\n\x10smallBlindButton\x18\x04 \x01(\r\x12\'\n\x07players\x18\x05 \x03(\x0b\x32\x16.thegamblr.PlayerStateB\x11Z\x0fthegamblr/protob\x06proto3')



_BOARDSTATE = DESCRIPTOR.message_types_by_name['BoardState']
BoardState = _reflection.GeneratedProtocolMessageType('BoardState', (_message.Message,), {
  'DESCRIPTOR' : _BOARDSTATE,
  '__module__' : 'board_state_pb2'
  # @@protoc_insertion_point(class_scope:thegamblr.BoardState)
  })
_sym_db.RegisterMessage(BoardState)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\017thegamblr/proto'
  _BOARDSTATE._serialized_start=78
  _BOARDSTATE._serialized_end=245
# @@protoc_insertion_point(module_scope)