//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAccountUpdate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeAccountUpdate] =
		sampleDataAccountUpdateOperation
}

var sampleDataAccountUpdateOperation = `{
  "account": "1.2.115296",
  "active": {
    "account_auths": [
      [
        "1.2.96086",
         1
      ],
      [
        "1.2.123456",
         1
      ]
    ],
    "address_auths": [],
    "key_auths": [
      [
        "BTS7jp4WPJaVerGJbNC4scaoHSdY7wSemAcdg6HikhNQqvegxYyyM",
        1
      ],
      [
        "BTS7wPKvv8zNBFWZmBQqPfW5Df9sfBvyS2c24MmJ1mkHjZCaowHNx",
        1
      ],
      [
        "BTS5zzvbDtkbUVU1gFFsKqCE55U7JbjTp6mTh1usFv7KGgXL7HDQk",
        1
      ]
    ],
    "weight_threshold": 1
  },
  "extensions": {
    "active_special_authority": [
      1,
      {
        "asset": "1.3.333",
        "num_top_holders": 15
      }
    ],
    "null_ext": {},
    "owner_special_authority": [
      1,
      {
        "asset": "1.3.555",
        "num_top_holders": 16
      }
    ]
  },
  "fee": {
    "amount": 28922,
    "asset_id": "1.3.0"
  },
  "new_options": {
    "extensions": [],
    "memo_key": "BTS7wPKvv8zNBFWZmBQqPfW5Df9sfBvyS2c24MmJ1mkHjZCaowHNx",
    "num_committee": 0,
    "num_witness": 0,
    "votes": [],
    "voting_account": "1.2.21106"
  },
  "owner": {
    "account_auths": [],
    "address_auths": [],
    "key_auths": [
      [
        "BTS7jp4WPJaVerGJbNC4scaoHSdY7wSemAcdg6HikhNQqvegxYyyM",
        1
      ]
    ],
    "weight_threshold": 1
  }
}`

//end of file
