// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test_data

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/vat_move"
)

var EthVatMoveLog = types.Log{
	Address: common.HexToAddress(constants.VatContractAddress),
	Topics: []common.Hash{
		common.HexToHash("0x78f1947000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x000000000000000000000000a730d1ff8b6bc74a26d54c20a9dda539909bab0e"),
		common.HexToHash("0x000000000000000000000000b730d1ff8b6bc74a26d54c20a9dda539909bab0e"),
		common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002a"),
	},
	Data:        hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000006478f19470a730d1ff8b6bc74a26d54c20a9dda539909bab0e000000000000000000000000b730d1ff8b6bc74a26d54c20a9dda539909bab0e000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a"),
	BlockNumber: 10,
	TxHash:      common.HexToHash("0xe8f39fbb7fea3621f543868f19b1114e305aff6a063a30d32835ff1012526f91"),
	TxIndex:     7,
	BlockHash:   fakes.FakeHash,
	Index:       8,
	Removed:     false,
}

var rawVatMoveLog, _ = json.Marshal(EthVatMoveLog)
var VatMoveModel = vat_move.VatMoveModel{
	Src:              "0xA730d1FF8B6Bc74a26d54c20a9dda539909BaB0e",
	Dst:              "0xB730D1fF8b6BC74a26D54c20a9ddA539909BAb0e",
	Rad:              "42",
	LogIndex:         EthVatMoveLog.Index,
	TransactionIndex: EthVatMoveLog.TxIndex,
	Raw:              rawVatMoveLog,
}
