// Code generated by protoc-gen-go.
// source: feed_info.proto
// DO NOT EDIT!

/*
Package feed_info is a generated protocol buffer package.

It is generated from these files:
	feed_info.proto

It has these top-level messages:
	FeedInfo
*/
package feed_info

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FeedInfo struct {
	Title            *string `protobuf:"bytes,1,req,name=title" json:"title,omitempty"`
	SourceUrl        *string `protobuf:"bytes,2,req,name=source_url" json:"source_url,omitempty"`
	BlockSearch      *string `protobuf:"bytes,3,opt,name=block_search" json:"block_search,omitempty"`
	BlockExtract     *string `protobuf:"bytes,4,opt,name=block_extract" json:"block_extract,omitempty"`
	ItemSearch       *string `protobuf:"bytes,5,req,name=item_search" json:"item_search,omitempty"`
	ItemTitle        *string `protobuf:"bytes,6,req,name=item_title" json:"item_title,omitempty"`
	ItemUrl          *string `protobuf:"bytes,7,opt,name=item_url" json:"item_url,omitempty"`
	ItemText         *string `protobuf:"bytes,8,opt,name=item_text" json:"item_text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FeedInfo) Reset()         { *m = FeedInfo{} }
func (m *FeedInfo) String() string { return proto.CompactTextString(m) }
func (*FeedInfo) ProtoMessage()    {}

func (m *FeedInfo) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *FeedInfo) GetSourceUrl() string {
	if m != nil && m.SourceUrl != nil {
		return *m.SourceUrl
	}
	return ""
}

func (m *FeedInfo) GetBlockSearch() string {
	if m != nil && m.BlockSearch != nil {
		return *m.BlockSearch
	}
	return ""
}

func (m *FeedInfo) GetBlockExtract() string {
	if m != nil && m.BlockExtract != nil {
		return *m.BlockExtract
	}
	return ""
}

func (m *FeedInfo) GetItemSearch() string {
	if m != nil && m.ItemSearch != nil {
		return *m.ItemSearch
	}
	return ""
}

func (m *FeedInfo) GetItemTitle() string {
	if m != nil && m.ItemTitle != nil {
		return *m.ItemTitle
	}
	return ""
}

func (m *FeedInfo) GetItemUrl() string {
	if m != nil && m.ItemUrl != nil {
		return *m.ItemUrl
	}
	return ""
}

func (m *FeedInfo) GetItemText() string {
	if m != nil && m.ItemText != nil {
		return *m.ItemText
	}
	return ""
}

func init() {
}
