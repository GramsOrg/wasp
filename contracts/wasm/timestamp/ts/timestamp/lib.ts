// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export function on_call(index: i32): void {
    return wasmlib.onCall(index);
}

export function on_load(): void {
    let exports = new wasmlib.ScExports();
    exports.addFunc(sc.FuncNow, funcNowThunk);
    exports.addView(sc.ViewGetTimestamp, viewGetTimestampThunk);

    for (let i = 0; i < sc.keyMap.length; i++) {
        sc.idxMap[i] = wasmlib.Key32.fromString(sc.keyMap[i]);
    }
}

function funcNowThunk(ctx: wasmlib.ScFuncContext): void {
    ctx.log("timestamp.funcNow");
    let f = new sc.NowContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
    sc.funcNow(ctx, f);
    ctx.log("timestamp.funcNow ok");
}

function viewGetTimestampThunk(ctx: wasmlib.ScViewContext): void {
    ctx.log("timestamp.viewGetTimestamp");
    let f = new sc.GetTimestampContext();
    f.results.mapID = wasmlib.OBJ_ID_RESULTS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
    sc.viewGetTimestamp(ctx, f);
    ctx.log("timestamp.viewGetTimestamp ok");
}
