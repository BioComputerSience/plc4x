/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetRecipientAddress is the data-structure of this message
type BACnetRecipientAddress struct {
	*BACnetRecipient
	AddressValue *BACnetAddressEnclosed
}

// IBACnetRecipientAddress is the corresponding interface of BACnetRecipientAddress
type IBACnetRecipientAddress interface {
	IBACnetRecipient
	// GetAddressValue returns AddressValue (property field)
	GetAddressValue() *BACnetAddressEnclosed
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetRecipientAddress) InitializeParent(parent *BACnetRecipient, peekedTagHeader *BACnetTagHeader) {
	m.BACnetRecipient.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetRecipientAddress) GetParent() *BACnetRecipient {
	return m.BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetRecipientAddress) GetAddressValue() *BACnetAddressEnclosed {
	return m.AddressValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientAddress factory function for BACnetRecipientAddress
func NewBACnetRecipientAddress(addressValue *BACnetAddressEnclosed, peekedTagHeader *BACnetTagHeader) *BACnetRecipientAddress {
	_result := &BACnetRecipientAddress{
		AddressValue:    addressValue,
		BACnetRecipient: NewBACnetRecipient(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetRecipientAddress(structType interface{}) *BACnetRecipientAddress {
	if casted, ok := structType.(BACnetRecipientAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetRecipientAddress); ok {
		return casted
	}
	if casted, ok := structType.(BACnetRecipient); ok {
		return CastBACnetRecipientAddress(casted.Child)
	}
	if casted, ok := structType.(*BACnetRecipient); ok {
		return CastBACnetRecipientAddress(casted.Child)
	}
	return nil
}

func (m *BACnetRecipientAddress) GetTypeName() string {
	return "BACnetRecipientAddress"
}

func (m *BACnetRecipientAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetRecipientAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (addressValue)
	lengthInBits += m.AddressValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetRecipientAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRecipientAddressParse(readBuffer utils.ReadBuffer) (*BACnetRecipientAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientAddress"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (addressValue)
	if pullErr := readBuffer.PullContext("addressValue"); pullErr != nil {
		return nil, pullErr
	}
	_addressValue, _addressValueErr := BACnetAddressEnclosedParse(readBuffer, uint8(uint8(1)))
	if _addressValueErr != nil {
		return nil, errors.Wrap(_addressValueErr, "Error parsing 'addressValue' field")
	}
	addressValue := CastBACnetAddressEnclosed(_addressValue)
	if closeErr := readBuffer.CloseContext("addressValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetRecipientAddress{
		AddressValue:    CastBACnetAddressEnclosed(addressValue),
		BACnetRecipient: &BACnetRecipient{},
	}
	_child.BACnetRecipient.Child = _child
	return _child, nil
}

func (m *BACnetRecipientAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientAddress"); pushErr != nil {
			return pushErr
		}

		// Simple Field (addressValue)
		if pushErr := writeBuffer.PushContext("addressValue"); pushErr != nil {
			return pushErr
		}
		_addressValueErr := m.AddressValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("addressValue"); popErr != nil {
			return popErr
		}
		if _addressValueErr != nil {
			return errors.Wrap(_addressValueErr, "Error serializing 'addressValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientAddress"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetRecipientAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
