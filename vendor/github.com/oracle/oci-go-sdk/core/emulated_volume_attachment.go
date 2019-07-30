// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// EmulatedVolumeAttachment An Emulated volume attachment.
type EmulatedVolumeAttachment struct {

	// The availability domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the volume attachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the instance the volume is attached to.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The date and time the volume was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// The device name.
	Device *string `mandatory:"false" json:"device"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My volume attachment`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `mandatory:"false" json:"isShareable"`

	// Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// The current state of the volume attachment.
	LifecycleState VolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m EmulatedVolumeAttachment) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

//GetCompartmentId returns CompartmentId
func (m EmulatedVolumeAttachment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDevice returns Device
func (m EmulatedVolumeAttachment) GetDevice() *string {
	return m.Device
}

//GetDisplayName returns DisplayName
func (m EmulatedVolumeAttachment) GetDisplayName() *string {
	return m.DisplayName
}

//GetId returns Id
func (m EmulatedVolumeAttachment) GetId() *string {
	return m.Id
}

//GetInstanceId returns InstanceId
func (m EmulatedVolumeAttachment) GetInstanceId() *string {
	return m.InstanceId
}

//GetIsReadOnly returns IsReadOnly
func (m EmulatedVolumeAttachment) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m EmulatedVolumeAttachment) GetIsShareable() *bool {
	return m.IsShareable
}

//GetLifecycleState returns LifecycleState
func (m EmulatedVolumeAttachment) GetLifecycleState() VolumeAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m EmulatedVolumeAttachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetVolumeId returns VolumeId
func (m EmulatedVolumeAttachment) GetVolumeId() *string {
	return m.VolumeId
}

//GetIsPvEncryptionInTransitEnabled returns IsPvEncryptionInTransitEnabled
func (m EmulatedVolumeAttachment) GetIsPvEncryptionInTransitEnabled() *bool {
	return m.IsPvEncryptionInTransitEnabled
}

func (m EmulatedVolumeAttachment) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EmulatedVolumeAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmulatedVolumeAttachment EmulatedVolumeAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeEmulatedVolumeAttachment
	}{
		"emulated",
		(MarshalTypeEmulatedVolumeAttachment)(m),
	}

	return json.Marshal(&s)
}