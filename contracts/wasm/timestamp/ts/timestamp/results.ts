// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export class ImmutableGetTimestampResults extends wasmlib.ScMapID {

    timestamp(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultTimestamp]);
    }
}

export class MutableGetTimestampResults extends wasmlib.ScMapID {

    timestamp(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultTimestamp]);
    }
}
