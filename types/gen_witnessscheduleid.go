// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/youthonline/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type WitnessScheduleID struct {
	ObjectID
}

func (p WitnessScheduleID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *WitnessScheduleID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeWitnessSchedule) << 48) | instance)
	return nil
}

type WitnessScheduleIDs []WitnessScheduleID

func (p WitnessScheduleIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode WitnessScheduleID")
		}
	}

	return nil
}

func WitnessScheduleIDFromObject(ob GrapheneObject) WitnessScheduleID {
	id, ok := ob.(*WitnessScheduleID)
	if ok {
		return *id
	}

	p := WitnessScheduleID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeWitnessSchedule {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeWitnessSchedule'", p.ID()))
	}

	return p
}

//NewWitnessScheduleID creates an new WitnessScheduleID object
func NewWitnessScheduleID(id string) GrapheneObject {
	gid := new(WitnessScheduleID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"WitnessScheduleID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeWitnessSchedule {
		logging.Errorf(
			"WitnessScheduleID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeWitnessSchedule'", id),
		)
		return nil
	}

	return gid
}
