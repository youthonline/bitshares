// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

type GlobalPropertyID struct {
	ObjectID
}

func (p GlobalPropertyID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *GlobalPropertyID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeGlobalProperty) << 48) | instance)
	return nil
}

type GlobalPropertyIDs []GlobalPropertyID

func (p GlobalPropertyIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode GlobalPropertyID")
		}
	}

	return nil
}

//NewGlobalPropertyID creates an new GlobalPropertyID object
func NewGlobalPropertyID(id string) *GlobalPropertyID {
	gid := new(GlobalPropertyID)
	if err := gid.Parse(id); err != nil {
		panic(errors.Annotate(err, "Parse"))
	}

	return gid
}
