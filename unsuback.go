// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message
import "fmt"

// The UNSUBACK Packet is sent by the Server to the Client to confirm receipt of an
// UNSUBSCRIBE Packet.
type UnsubackMessage struct {
	// Shared message identifier.
	PacketId uint16
}

var _ Message = (*UnsubackMessage)(nil)

// NewUnsubackMessage creates a new UNSUBACK message.
func NewUnsubackMessage() *UnsubackMessage {
	msg := &UnsubackMessage{}
	return msg
}

func (this UnsubackMessage) Type() MessageType {
	return UNSUBACK
}

func (this *UnsubackMessage) Len() int {
	return identifiedMessageLen()
}

func (this *UnsubackMessage) Decode(src []byte) (int, error) {
	n, pid, err := identifiedMessageDecode(src, UNSUBACK)
	this.PacketId = pid
	return n, err
}

func (this *UnsubackMessage) Encode(dst []byte) (int, error) {
	return identifiedMessageEncode(dst, this.PacketId, UNSUBACK)
}

// String returns a string representation of the message.
func (this UnsubackMessage) String() string {
	return fmt.Sprintf("UNSUBACK: PacketId=%d", this.PacketId)
}
