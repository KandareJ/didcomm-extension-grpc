# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: basic-message.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='basic-message.proto',
  package='didcomm.messaging',
  syntax='proto3',
  serialized_options=b'\252\002\021DIDComm.Messaging',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x13\x62\x61sic-message.proto\x12\x11\x64idcomm.messaging\"\x1c\n\x0c\x42\x61sicMessage\x12\x0c\n\x04text\x18\x01 \x01(\tB\x14\xaa\x02\x11\x44IDComm.Messagingb\x06proto3'
)




_BASICMESSAGE = _descriptor.Descriptor(
  name='BasicMessage',
  full_name='didcomm.messaging.BasicMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='didcomm.messaging.BasicMessage.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=70,
)

DESCRIPTOR.message_types_by_name['BasicMessage'] = _BASICMESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

BasicMessage = _reflection.GeneratedProtocolMessageType('BasicMessage', (_message.Message,), {
  'DESCRIPTOR' : _BASICMESSAGE,
  '__module__' : 'basic_message_pb2'
  # @@protoc_insertion_point(class_scope:didcomm.messaging.BasicMessage)
  })
_sym_db.RegisterMessage(BasicMessage)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
