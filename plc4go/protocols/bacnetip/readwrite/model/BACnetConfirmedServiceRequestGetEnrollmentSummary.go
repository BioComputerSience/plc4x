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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestGetEnrollmentSummary is the data-structure of this message
type BACnetConfirmedServiceRequestGetEnrollmentSummary struct {
	*BACnetConfirmedServiceRequest
	AcknowledgmentFilter    *BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	EnrollmentFilter        *BACnetRecipientProcessEnclosed
	EventStateFilter        *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	EventTypeFilter         *BACnetEventTypeTagged
	PriorityFilter          *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	NotificationClassFilter *BACnetContextTagUnsignedInteger

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestGetEnrollmentSummary is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummary
type IBACnetConfirmedServiceRequestGetEnrollmentSummary interface {
	IBACnetConfirmedServiceRequest
	// GetAcknowledgmentFilter returns AcknowledgmentFilter (property field)
	GetAcknowledgmentFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	// GetEnrollmentFilter returns EnrollmentFilter (property field)
	GetEnrollmentFilter() *BACnetRecipientProcessEnclosed
	// GetEventStateFilter returns EventStateFilter (property field)
	GetEventStateFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	// GetEventTypeFilter returns EventTypeFilter (property field)
	GetEventTypeFilter() *BACnetEventTypeTagged
	// GetPriorityFilter returns PriorityFilter (property field)
	GetPriorityFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	// GetNotificationClassFilter returns NotificationClassFilter (property field)
	GetNotificationClassFilter() *BACnetContextTagUnsignedInteger
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

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetAcknowledgmentFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	return m.AcknowledgmentFilter
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEnrollmentFilter() *BACnetRecipientProcessEnclosed {
	return m.EnrollmentFilter
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventStateFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged {
	return m.EventStateFilter
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventTypeFilter() *BACnetEventTypeTagged {
	return m.EventTypeFilter
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetPriorityFilter() *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	return m.PriorityFilter
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetNotificationClassFilter() *BACnetContextTagUnsignedInteger {
	return m.NotificationClassFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEnrollmentSummary factory function for BACnetConfirmedServiceRequestGetEnrollmentSummary
func NewBACnetConfirmedServiceRequestGetEnrollmentSummary(acknowledgmentFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, enrollmentFilter *BACnetRecipientProcessEnclosed, eventStateFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged, eventTypeFilter *BACnetEventTypeTagged, priorityFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, notificationClassFilter *BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *BACnetConfirmedServiceRequestGetEnrollmentSummary {
	_result := &BACnetConfirmedServiceRequestGetEnrollmentSummary{
		AcknowledgmentFilter:          acknowledgmentFilter,
		EnrollmentFilter:              enrollmentFilter,
		EventStateFilter:              eventStateFilter,
		EventTypeFilter:               eventTypeFilter,
		PriorityFilter:                priorityFilter,
		NotificationClassFilter:       notificationClassFilter,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestGetEnrollmentSummary(structType interface{}) *BACnetConfirmedServiceRequestGetEnrollmentSummary {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestGetEnrollmentSummary(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestGetEnrollmentSummary(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummary"
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (acknowledgmentFilter)
	lengthInBits += m.AcknowledgmentFilter.GetLengthInBits()

	// Optional Field (enrollmentFilter)
	if m.EnrollmentFilter != nil {
		lengthInBits += (*m.EnrollmentFilter).GetLengthInBits()
	}

	// Optional Field (eventStateFilter)
	if m.EventStateFilter != nil {
		lengthInBits += (*m.EventStateFilter).GetLengthInBits()
	}

	// Optional Field (eventTypeFilter)
	if m.EventTypeFilter != nil {
		lengthInBits += (*m.EventTypeFilter).GetLengthInBits()
	}

	// Optional Field (priorityFilter)
	if m.PriorityFilter != nil {
		lengthInBits += (*m.PriorityFilter).GetLengthInBits()
	}

	// Optional Field (notificationClassFilter)
	if m.NotificationClassFilter != nil {
		lengthInBits += (*m.NotificationClassFilter).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (acknowledgmentFilter)
	if pullErr := readBuffer.PullContext("acknowledgmentFilter"); pullErr != nil {
		return nil, pullErr
	}
	_acknowledgmentFilter, _acknowledgmentFilterErr := BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParse(readBuffer, uint8(uint8(0)))
	if _acknowledgmentFilterErr != nil {
		return nil, errors.Wrap(_acknowledgmentFilterErr, "Error parsing 'acknowledgmentFilter' field")
	}
	acknowledgmentFilter := CastBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged(_acknowledgmentFilter)
	if closeErr := readBuffer.CloseContext("acknowledgmentFilter"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (enrollmentFilter) (Can be skipped, if a given expression evaluates to false)
	var enrollmentFilter *BACnetRecipientProcessEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("enrollmentFilter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetRecipientProcessEnclosedParse(readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'enrollmentFilter' field")
		default:
			enrollmentFilter = CastBACnetRecipientProcessEnclosed(_val)
			if closeErr := readBuffer.CloseContext("enrollmentFilter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (eventStateFilter) (Can be skipped, if a given expression evaluates to false)
	var eventStateFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("eventStateFilter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParse(readBuffer, uint8(2))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventStateFilter' field")
		default:
			eventStateFilter = CastBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged(_val)
			if closeErr := readBuffer.CloseContext("eventStateFilter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (eventTypeFilter) (Can be skipped, if a given expression evaluates to false)
	var eventTypeFilter *BACnetEventTypeTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("eventTypeFilter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetEventTypeTaggedParse(readBuffer, uint8(3))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventTypeFilter' field")
		default:
			eventTypeFilter = CastBACnetEventTypeTagged(_val)
			if closeErr := readBuffer.CloseContext("eventTypeFilter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (priorityFilter) (Can be skipped, if a given expression evaluates to false)
	var priorityFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priorityFilter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParse(readBuffer, uint8(4))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priorityFilter' field")
		default:
			priorityFilter = CastBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(_val)
			if closeErr := readBuffer.CloseContext("priorityFilter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (notificationClassFilter) (Can be skipped, if a given expression evaluates to false)
	var notificationClassFilter *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("notificationClassFilter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(5), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'notificationClassFilter' field")
		default:
			notificationClassFilter = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("notificationClassFilter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestGetEnrollmentSummary{
		AcknowledgmentFilter:          CastBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged(acknowledgmentFilter),
		EnrollmentFilter:              CastBACnetRecipientProcessEnclosed(enrollmentFilter),
		EventStateFilter:              CastBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged(eventStateFilter),
		EventTypeFilter:               CastBACnetEventTypeTagged(eventTypeFilter),
		PriorityFilter:                CastBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(priorityFilter),
		NotificationClassFilter:       CastBACnetContextTagUnsignedInteger(notificationClassFilter),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pushErr != nil {
			return pushErr
		}

		// Simple Field (acknowledgmentFilter)
		if pushErr := writeBuffer.PushContext("acknowledgmentFilter"); pushErr != nil {
			return pushErr
		}
		_acknowledgmentFilterErr := m.AcknowledgmentFilter.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("acknowledgmentFilter"); popErr != nil {
			return popErr
		}
		if _acknowledgmentFilterErr != nil {
			return errors.Wrap(_acknowledgmentFilterErr, "Error serializing 'acknowledgmentFilter' field")
		}

		// Optional Field (enrollmentFilter) (Can be skipped, if the value is null)
		var enrollmentFilter *BACnetRecipientProcessEnclosed = nil
		if m.EnrollmentFilter != nil {
			if pushErr := writeBuffer.PushContext("enrollmentFilter"); pushErr != nil {
				return pushErr
			}
			enrollmentFilter = m.EnrollmentFilter
			_enrollmentFilterErr := enrollmentFilter.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("enrollmentFilter"); popErr != nil {
				return popErr
			}
			if _enrollmentFilterErr != nil {
				return errors.Wrap(_enrollmentFilterErr, "Error serializing 'enrollmentFilter' field")
			}
		}

		// Optional Field (eventStateFilter) (Can be skipped, if the value is null)
		var eventStateFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged = nil
		if m.EventStateFilter != nil {
			if pushErr := writeBuffer.PushContext("eventStateFilter"); pushErr != nil {
				return pushErr
			}
			eventStateFilter = m.EventStateFilter
			_eventStateFilterErr := eventStateFilter.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("eventStateFilter"); popErr != nil {
				return popErr
			}
			if _eventStateFilterErr != nil {
				return errors.Wrap(_eventStateFilterErr, "Error serializing 'eventStateFilter' field")
			}
		}

		// Optional Field (eventTypeFilter) (Can be skipped, if the value is null)
		var eventTypeFilter *BACnetEventTypeTagged = nil
		if m.EventTypeFilter != nil {
			if pushErr := writeBuffer.PushContext("eventTypeFilter"); pushErr != nil {
				return pushErr
			}
			eventTypeFilter = m.EventTypeFilter
			_eventTypeFilterErr := eventTypeFilter.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("eventTypeFilter"); popErr != nil {
				return popErr
			}
			if _eventTypeFilterErr != nil {
				return errors.Wrap(_eventTypeFilterErr, "Error serializing 'eventTypeFilter' field")
			}
		}

		// Optional Field (priorityFilter) (Can be skipped, if the value is null)
		var priorityFilter *BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter = nil
		if m.PriorityFilter != nil {
			if pushErr := writeBuffer.PushContext("priorityFilter"); pushErr != nil {
				return pushErr
			}
			priorityFilter = m.PriorityFilter
			_priorityFilterErr := priorityFilter.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("priorityFilter"); popErr != nil {
				return popErr
			}
			if _priorityFilterErr != nil {
				return errors.Wrap(_priorityFilterErr, "Error serializing 'priorityFilter' field")
			}
		}

		// Optional Field (notificationClassFilter) (Can be skipped, if the value is null)
		var notificationClassFilter *BACnetContextTagUnsignedInteger = nil
		if m.NotificationClassFilter != nil {
			if pushErr := writeBuffer.PushContext("notificationClassFilter"); pushErr != nil {
				return pushErr
			}
			notificationClassFilter = m.NotificationClassFilter
			_notificationClassFilterErr := notificationClassFilter.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("notificationClassFilter"); popErr != nil {
				return popErr
			}
			if _notificationClassFilterErr != nil {
				return errors.Wrap(_notificationClassFilterErr, "Error serializing 'notificationClassFilter' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
