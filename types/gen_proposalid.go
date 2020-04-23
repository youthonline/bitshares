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

type ProposalID struct {
	ObjectID
}

func (p ProposalID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *ProposalID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeProposal) << 48) | instance)
	return nil
}

type ProposalIDs []ProposalID

func (p ProposalIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode ProposalID")
		}
	}

	return nil
}

func ProposalIDFromObject(ob GrapheneObject) ProposalID {
	id, ok := ob.(*ProposalID)
	if ok {
		return *id
	}

	p := ProposalID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeProposal {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeProposal'", p.ID()))
	}

	return p
}

//NewProposalID creates an new ProposalID object
func NewProposalID(id string) GrapheneObject {
	gid := new(ProposalID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"ProposalID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeProposal {
		logging.Errorf(
			"ProposalID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeProposal'", id),
		)
		return nil
	}

	return gid
}
