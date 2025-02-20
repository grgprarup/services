// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: proto/wordle.proto

package wordle

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

type Guess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the full guess word
	Word string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	// the highlighted word e.g n[o]is{e}
	// where [ ] is correct, { } is in word
	Highlight string `protobuf:"bytes,2,opt,name=highlight,proto3" json:"highlight,omitempty"`
	// individual characters
	Chars []*Char `protobuf:"bytes,3,rep,name=chars,proto3" json:"chars,omitempty"`
}

func (x *Guess) Reset() {
	*x = Guess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guess) ProtoMessage() {}

func (x *Guess) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guess.ProtoReflect.Descriptor instead.
func (*Guess) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{0}
}

func (x *Guess) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Guess) GetHighlight() string {
	if x != nil {
		return x.Highlight
	}
	return ""
}

func (x *Guess) GetChars() []*Char {
	if x != nil {
		return x.Chars
	}
	return nil
}

type Char struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the character itself
	Letter string `protobuf:"bytes,1,opt,name=letter,proto3" json:"letter,omitempty"`
	// position in the string
	Position int32 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	// whether it was correct
	Correct bool `protobuf:"varint,3,opt,name=correct,proto3" json:"correct,omitempty"`
	// whether it's in the word
	InWord bool `protobuf:"varint,4,opt,name=in_word,json=inWord,proto3" json:"in_word,omitempty"`
}

func (x *Char) Reset() {
	*x = Char{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Char) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Char) ProtoMessage() {}

func (x *Char) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Char.ProtoReflect.Descriptor instead.
func (*Char) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{1}
}

func (x *Char) GetLetter() string {
	if x != nil {
		return x.Letter
	}
	return ""
}

func (x *Char) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Char) GetCorrect() bool {
	if x != nil {
		return x.Correct
	}
	return false
}

func (x *Char) GetInWord() bool {
	if x != nil {
		return x.InWord
	}
	return false
}

// Make a guess
type GuessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player
	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	// guess word
	Word string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *GuessRequest) Reset() {
	*x = GuessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessRequest) ProtoMessage() {}

func (x *GuessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessRequest.ProtoReflect.Descriptor instead.
func (*GuessRequest) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{2}
}

func (x *GuessRequest) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *GuessRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type GuessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// whether it was correct
	Correct bool `protobuf:"varint,1,opt,name=correct,proto3" json:"correct,omitempty"`
	// number of tries left
	TriesLeft int32 `protobuf:"varint,2,opt,name=tries_left,json=triesLeft,proto3" json:"tries_left,omitempty"`
	// the guess words tried
	Guesses []*Guess `protobuf:"bytes,3,rep,name=guesses,proto3" json:"guesses,omitempty"`
	// the actual word if failed
	Answer string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
	// informational message
	Status string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GuessResponse) Reset() {
	*x = GuessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessResponse) ProtoMessage() {}

func (x *GuessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessResponse.ProtoReflect.Descriptor instead.
func (*GuessResponse) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{3}
}

func (x *GuessResponse) GetCorrect() bool {
	if x != nil {
		return x.Correct
	}
	return false
}

func (x *GuessResponse) GetTriesLeft() int32 {
	if x != nil {
		return x.TriesLeft
	}
	return 0
}

func (x *GuessResponse) GetGuesses() []*Guess {
	if x != nil {
		return x.Guesses
	}
	return nil
}

func (x *GuessResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *GuessResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// When does the next game start
type NextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextRequest) Reset() {
	*x = NextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRequest) ProtoMessage() {}

func (x *NextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRequest.ProtoReflect.Descriptor instead.
func (*NextRequest) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{4}
}

type NextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of seconds
	Seconds int32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// in hh:mm:ss
	Duration string `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *NextResponse) Reset() {
	*x = NextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wordle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextResponse) ProtoMessage() {}

func (x *NextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wordle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextResponse.ProtoReflect.Descriptor instead.
func (*NextResponse) Descriptor() ([]byte, []int) {
	return file_proto_wordle_proto_rawDescGZIP(), []int{5}
}

func (x *NextResponse) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *NextResponse) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

var File_proto_wordle_proto protoreflect.FileDescriptor

var file_proto_wordle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x22, 0x5d, 0x0a, 0x05,
	0x47, 0x75, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x67,
	0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x69,
	0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x22, 0x6d, 0x0a, 0x04, 0x43,
	0x68, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x6e, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x0a, 0x0c, 0x47, 0x75,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x47, 0x75, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x66, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x65, 0x66,
	0x74, 0x12, 0x27, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e, 0x47, 0x75, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0c, 0x4e, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x75, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x47, 0x75, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e, 0x47, 0x75, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x6c,
	0x65, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x64,
	0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_wordle_proto_rawDescOnce sync.Once
	file_proto_wordle_proto_rawDescData = file_proto_wordle_proto_rawDesc
)

func file_proto_wordle_proto_rawDescGZIP() []byte {
	file_proto_wordle_proto_rawDescOnce.Do(func() {
		file_proto_wordle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_wordle_proto_rawDescData)
	})
	return file_proto_wordle_proto_rawDescData
}

var file_proto_wordle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_wordle_proto_goTypes = []interface{}{
	(*Guess)(nil),         // 0: wordle.Guess
	(*Char)(nil),          // 1: wordle.Char
	(*GuessRequest)(nil),  // 2: wordle.GuessRequest
	(*GuessResponse)(nil), // 3: wordle.GuessResponse
	(*NextRequest)(nil),   // 4: wordle.NextRequest
	(*NextResponse)(nil),  // 5: wordle.NextResponse
}
var file_proto_wordle_proto_depIdxs = []int32{
	1, // 0: wordle.Guess.chars:type_name -> wordle.Char
	0, // 1: wordle.GuessResponse.guesses:type_name -> wordle.Guess
	2, // 2: wordle.Wordle.Guess:input_type -> wordle.GuessRequest
	4, // 3: wordle.Wordle.Next:input_type -> wordle.NextRequest
	3, // 4: wordle.Wordle.Guess:output_type -> wordle.GuessResponse
	5, // 5: wordle.Wordle.Next:output_type -> wordle.NextResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_wordle_proto_init() }
func file_proto_wordle_proto_init() {
	if File_proto_wordle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_wordle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guess); i {
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
		file_proto_wordle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Char); i {
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
		file_proto_wordle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessRequest); i {
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
		file_proto_wordle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessResponse); i {
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
		file_proto_wordle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRequest); i {
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
		file_proto_wordle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextResponse); i {
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
			RawDescriptor: file_proto_wordle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_wordle_proto_goTypes,
		DependencyIndexes: file_proto_wordle_proto_depIdxs,
		MessageInfos:      file_proto_wordle_proto_msgTypes,
	}.Build()
	File_proto_wordle_proto = out.File
	file_proto_wordle_proto_rawDesc = nil
	file_proto_wordle_proto_goTypes = nil
	file_proto_wordle_proto_depIdxs = nil
}
