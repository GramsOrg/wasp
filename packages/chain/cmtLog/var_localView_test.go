// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package cmtLog_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/chain/cmtLog"
	"github.com/iotaledger/wasp/packages/isc"
)

func TestLocalView(t *testing.T) {
	j := cmtLog.NewVarLocalView()
	require.Nil(t, j.GetBaseAliasOutput())
	require.True(t, j.AliasOutputConfirmed(isc.NewAliasOutputWithID(&iotago.AliasOutput{}, iotago.OutputID{})))
	require.NotNil(t, j.GetBaseAliasOutput())
}
