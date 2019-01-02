// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

type TransactionID struct {
	ObjectID
}

func (p TransactionID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *TransactionID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeTransaction) << 48) | instance)
	return nil
}

type TransactionIDs []TransactionID

func (p TransactionIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode TransactionID")
		}
	}

	return nil
}

//NewTransactionID creates an new TransactionID object
func NewTransactionID(id string) *TransactionID {
	gid := new(TransactionID)
	if err := gid.Parse(id); err != nil {
		panic(errors.Annotate(err, "Parse"))
	}

	return gid
}
