// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Metadata struct {
	_tab flatbuffers.Table
}

func GetRootAsMetadata(buf []byte, offset flatbuffers.UOffsetT) *Metadata {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Metadata{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMetadata(buf []byte, offset flatbuffers.UOffsetT) *Metadata {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Metadata{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Metadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Metadata) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Metadata) Inode() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateInode(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Metadata) Uid() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateUid(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Metadata) Gid() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateGid(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *Metadata) Sid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Metadata) Mode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateMode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *Metadata) Size() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateSize(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *Metadata) MtimeNs() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateMtimeNs(n int64) bool {
	return rcv._tab.MutateInt64Slot(16, n)
}

func (rcv *Metadata) CtimeNs() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Metadata) MutateCtimeNs(n int64) bool {
	return rcv._tab.MutateInt64Slot(18, n)
}

func (rcv *Metadata) Type() Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return Type(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *Metadata) MutateType(n Type) bool {
	return rcv._tab.MutateByteSlot(20, byte(n))
}

func (rcv *Metadata) Selinux() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Metadata) PosixAclAccess(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Metadata) PosixAclAccessLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Metadata) MutatePosixAclAccess(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func MetadataStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func MetadataAddInode(builder *flatbuffers.Builder, inode uint64) {
	builder.PrependUint64Slot(0, inode, 0)
}
func MetadataAddUid(builder *flatbuffers.Builder, uid uint32) {
	builder.PrependUint32Slot(1, uid, 0)
}
func MetadataAddGid(builder *flatbuffers.Builder, gid uint32) {
	builder.PrependUint32Slot(2, gid, 0)
}
func MetadataAddSid(builder *flatbuffers.Builder, sid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(sid), 0)
}
func MetadataAddMode(builder *flatbuffers.Builder, mode uint32) {
	builder.PrependUint32Slot(4, mode, 0)
}
func MetadataAddSize(builder *flatbuffers.Builder, size uint64) {
	builder.PrependUint64Slot(5, size, 0)
}
func MetadataAddMtimeNs(builder *flatbuffers.Builder, mtimeNs int64) {
	builder.PrependInt64Slot(6, mtimeNs, 0)
}
func MetadataAddCtimeNs(builder *flatbuffers.Builder, ctimeNs int64) {
	builder.PrependInt64Slot(7, ctimeNs, 0)
}
func MetadataAddType(builder *flatbuffers.Builder, type_ Type) {
	builder.PrependByteSlot(8, byte(type_), 1)
}
func MetadataAddSelinux(builder *flatbuffers.Builder, selinux flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(selinux), 0)
}
func MetadataAddPosixAclAccess(builder *flatbuffers.Builder, posixAclAccess flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(posixAclAccess), 0)
}
func MetadataStartPosixAclAccessVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
