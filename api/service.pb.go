// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: service.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Coreference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Confidence for the coreference
	Score float64 `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
	// The target sentence ID
	SentenceId int32 `protobuf:"varint,2,opt,name=sentenceId,proto3" json:"sentenceId,omitempty"`
	// The target token ID
	TokenId int32 `protobuf:"varint,3,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
}

func (x *Coreference) Reset() {
	*x = Coreference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coreference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coreference) ProtoMessage() {}

func (x *Coreference) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coreference.ProtoReflect.Descriptor instead.
func (*Coreference) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Coreference) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Coreference) GetSentenceId() int32 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *Coreference) GetTokenId() int32 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Morphology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Grammatical case
	Case string `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	// Grammatical gender in the sentence context
	CtxGender string `protobuf:"bytes,2,opt,name=ctx_gender,json=ctxGender,proto3" json:"ctx_gender,omitempty"`
	// Verbal mood in the sentence context
	CtxMood string `protobuf:"bytes,3,opt,name=ctx_mood,json=ctxMood,proto3" json:"ctx_mood,omitempty"`
	// Number in the sentence context
	CtxNumber string `protobuf:"bytes,4,opt,name=ctx_number,json=ctxNumber,proto3" json:"ctx_number,omitempty"`
	// Person in the sentence context
	CtxPerson string `protobuf:"bytes,5,opt,name=ctx_person,json=ctxPerson,proto3" json:"ctx_person,omitempty"`
	// Verbal voice in the sentence context
	CtxVoice string `protobuf:"bytes,6,opt,name=ctx_voice,json=ctxVoice,proto3" json:"ctx_voice,omitempty"`
	// Degree
	Degree string `protobuf:"bytes,7,opt,name=degree,proto3" json:"degree,omitempty"`
	// Grammatical gender
	Gender string `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	// The lemma
	Lemma string `protobuf:"bytes,9,opt,name=lemma,proto3" json:"lemma,omitempty"`
	// Verbal mood
	Mood string `protobuf:"bytes,10,opt,name=mood,proto3" json:"mood,omitempty"`
	// Number
	Number string `protobuf:"bytes,11,opt,name=number,proto3" json:"number,omitempty"`
	// Person
	Person string `protobuf:"bytes,12,opt,name=person,proto3" json:"person,omitempty"`
	// Part-of-speech tag
	Pos string `protobuf:"bytes,13,opt,name=pos,proto3" json:"pos,omitempty"`
	// Verbal tense
	Tense string `protobuf:"bytes,14,opt,name=tense,proto3" json:"tense,omitempty"`
}

func (x *Morphology) Reset() {
	*x = Morphology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Morphology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Morphology) ProtoMessage() {}

func (x *Morphology) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Morphology.ProtoReflect.Descriptor instead.
func (*Morphology) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Morphology) GetCase() string {
	if x != nil {
		return x.Case
	}
	return ""
}

func (x *Morphology) GetCtxGender() string {
	if x != nil {
		return x.CtxGender
	}
	return ""
}

func (x *Morphology) GetCtxMood() string {
	if x != nil {
		return x.CtxMood
	}
	return ""
}

func (x *Morphology) GetCtxNumber() string {
	if x != nil {
		return x.CtxNumber
	}
	return ""
}

func (x *Morphology) GetCtxPerson() string {
	if x != nil {
		return x.CtxPerson
	}
	return ""
}

func (x *Morphology) GetCtxVoice() string {
	if x != nil {
		return x.CtxVoice
	}
	return ""
}

func (x *Morphology) GetDegree() string {
	if x != nil {
		return x.Degree
	}
	return ""
}

func (x *Morphology) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Morphology) GetLemma() string {
	if x != nil {
		return x.Lemma
	}
	return ""
}

func (x *Morphology) GetMood() string {
	if x != nil {
		return x.Mood
	}
	return ""
}

func (x *Morphology) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Morphology) GetPerson() string {
	if x != nil {
		return x.Person
	}
	return ""
}

func (x *Morphology) GetPos() string {
	if x != nil {
		return x.Pos
	}
	return ""
}

func (x *Morphology) GetTense() string {
	if x != nil {
		return x.Tense
	}
	return ""
}

type PostForParsingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request to the parsing service
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *PostForParsingRequest) Reset() {
	*x = PostForParsingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostForParsingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostForParsingRequest) ProtoMessage() {}

func (x *PostForParsingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostForParsingRequest.ProtoReflect.Descriptor instead.
func (*PostForParsingRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostForParsingRequest) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text to parse
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of sentences
	Sentences []*Sentence `protobuf:"bytes,1,rep,name=sentences,proto3" json:"sentences,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetSentences() []*Sentence {
	if x != nil {
		return x.Sentences
	}
	return nil
}

type Sentence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the sentence
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Confidence of overall sentence attachments
	Score float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	// List of tokens
	Tokens []*Token `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *Sentence) Reset() {
	*x = Sentence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sentence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sentence) ProtoMessage() {}

func (x *Sentence) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sentence.ProtoReflect.Descriptor instead.
func (*Sentence) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *Sentence) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sentence) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Sentence) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of tokens
	Coref []*Coreference `protobuf:"bytes,1,rep,name=coref,proto3" json:"coref,omitempty"`
	// Dependency relation
	Deprel string `protobuf:"bytes,2,opt,name=deprel,proto3" json:"deprel,omitempty"`
	// Index of the last character of the token in the original text
	End int32 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	// Form of the token as in the original text
	Form string `protobuf:"bytes,4,opt,name=form,proto3" json:"form,omitempty"`
	// Syntactic head (-1 for the ROOT)
	Head int32 `protobuf:"varint,5,opt,name=head,proto3" json:"head,omitempty"`
	// ID of the token
	Id int32 `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	// List of possible morphological interpretations
	Morpho []*Morphology `protobuf:"bytes,7,rep,name=morpho,proto3" json:"morpho,omitempty"`
	// Coarse-grained part-of-speech tag
	Pos string `protobuf:"bytes,8,opt,name=pos,proto3" json:"pos,omitempty"`
	// Index of the token (not the ID)
	Position string `protobuf:"bytes,9,opt,name=position,proto3" json:"position,omitempty"`
	// Confidence for the syntactic attachment
	Score float64 `protobuf:"fixed64,10,opt,name=score,proto3" json:"score,omitempty"`
	// List of semantic labels
	Semantic []string `protobuf:"bytes,11,rep,name=semantic,proto3" json:"semantic,omitempty"`
	// Index of the first character of the token in the original text
	Start int32 `protobuf:"varint,12,opt,name=start,proto3" json:"start,omitempty"`
	// Token type (one of WORD, WORD_TRACE, TRACE)
	Type string `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *Token) GetCoref() []*Coreference {
	if x != nil {
		return x.Coref
	}
	return nil
}

func (x *Token) GetDeprel() string {
	if x != nil {
		return x.Deprel
	}
	return ""
}

func (x *Token) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Token) GetForm() string {
	if x != nil {
		return x.Form
	}
	return ""
}

func (x *Token) GetHead() int32 {
	if x != nil {
		return x.Head
	}
	return 0
}

func (x *Token) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Token) GetMorpho() []*Morphology {
	if x != nil {
		return x.Morpho
	}
	return nil
}

func (x *Token) GetPos() string {
	if x != nil {
		return x.Pos
	}
	return ""
}

func (x *Token) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Token) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Token) GetSemantic() []string {
	if x != nil {
		return x.Semantic
	}
	return nil
}

func (x *Token) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Token) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d,
	0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x21, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xe7, 0x02, 0x0a, 0x0a, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x74, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x74, 0x78, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x74, 0x78, 0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x74, 0x78, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x74, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x74, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x74, 0x78, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x74, 0x78, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x74, 0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x74, 0x78, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x6d, 0x6d,
	0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x6d, 0x6d, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f,
	0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x15, 0x50, 0x6f,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xdc, 0x02,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a,
	0x06, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x06, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x84, 0x01, 0x0a,
	0x14, 0x50, 0x61, 0x72, 0x73, 0x69, 0x74, 0x41, 0x70, 0x69, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72,
	0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f,
	0x72, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x22, 0x06, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_proto_goTypes = []interface{}{
	(*Coreference)(nil),           // 0: parsit_api_spec.Coreference
	(*Error)(nil),                 // 1: parsit_api_spec.Error
	(*Morphology)(nil),            // 2: parsit_api_spec.Morphology
	(*PostForParsingRequest)(nil), // 3: parsit_api_spec.PostForParsingRequest
	(*Request)(nil),               // 4: parsit_api_spec.Request
	(*Response)(nil),              // 5: parsit_api_spec.Response
	(*Sentence)(nil),              // 6: parsit_api_spec.Sentence
	(*Token)(nil),                 // 7: parsit_api_spec.Token
}
var file_service_proto_depIdxs = []int32{
	4, // 0: parsit_api_spec.PostForParsingRequest.request:type_name -> parsit_api_spec.Request
	6, // 1: parsit_api_spec.Response.sentences:type_name -> parsit_api_spec.Sentence
	7, // 2: parsit_api_spec.Sentence.tokens:type_name -> parsit_api_spec.Token
	0, // 3: parsit_api_spec.Token.coref:type_name -> parsit_api_spec.Coreference
	2, // 4: parsit_api_spec.Token.morpho:type_name -> parsit_api_spec.Morphology
	3, // 5: parsit_api_spec.ParsitApiSpecService.PostForParsing:input_type -> parsit_api_spec.PostForParsingRequest
	5, // 6: parsit_api_spec.ParsitApiSpecService.PostForParsing:output_type -> parsit_api_spec.Response
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coreference); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Morphology); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostForParsingRequest); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sentence); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
