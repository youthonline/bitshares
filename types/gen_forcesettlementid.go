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

type ForceSettlementID struct {
	ObjectID
}

func (p ForceSettlementID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *ForceSettlementID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeForceSettlement) << 48) | instance)
	return nil
}

type ForceSettlementIDs []ForceSettlementID

func (p ForceSettlementIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode ForceSettlementID")
		}
	}

	return nil
}

func ForceSettlementIDFromObject(ob GrapheneObject) ForceSettlementID {
	id, ok := ob.(*ForceSettlementID)
	if ok {
		return *id
	}

	p := ForceSettlementID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeForceSettlement {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeForceSettlement'", p.ID()))
	}

	return p
}

//NewForceSettlementID creates an new ForceSettlementID object
func NewForceSettlementID(id string) GrapheneObject {
	gid := new(ForceSettlementID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"ForceSettlementID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeForceSettlement {
		logging.Errorf(
			"ForceSettlementID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeForceSettlement'", id),
		)
		return nil
	}

	return gid
}
