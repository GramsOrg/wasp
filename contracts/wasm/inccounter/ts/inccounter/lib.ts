// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export function on_call(index: i32): void {
    return wasmlib.onCall(index);
}

export function on_load(): void {
    let exports = new wasmlib.ScExports();
    exports.addFunc(sc.FuncCallIncrement,          funcCallIncrementThunk);
    exports.addFunc(sc.FuncCallIncrementRecurse5x, funcCallIncrementRecurse5xThunk);
    exports.addFunc(sc.FuncEndlessLoop,            funcEndlessLoopThunk);
    exports.addFunc(sc.FuncIncrement,              funcIncrementThunk);
    exports.addFunc(sc.FuncIncrementWithDelay,     funcIncrementWithDelayThunk);
    exports.addFunc(sc.FuncInit,                   funcInitThunk);
    exports.addFunc(sc.FuncLocalStateInternalCall, funcLocalStateInternalCallThunk);
    exports.addFunc(sc.FuncLocalStatePost,         funcLocalStatePostThunk);
    exports.addFunc(sc.FuncLocalStateSandboxCall,  funcLocalStateSandboxCallThunk);
    exports.addFunc(sc.FuncPostIncrement,          funcPostIncrementThunk);
    exports.addFunc(sc.FuncRepeatMany,             funcRepeatManyThunk);
    exports.addFunc(sc.FuncTestLeb128,             funcTestLeb128Thunk);
    exports.addFunc(sc.FuncWhenMustIncrement,      funcWhenMustIncrementThunk);
    exports.addView(sc.ViewGetCounter,             viewGetCounterThunk);

    for (let i = 0; i < sc.keyMap.length; i++) {
        sc.idxMap[i] = wasmlib.Key32.fromString(sc.keyMap[i]);
    }
}

function funcCallIncrementThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcCallIncrement");
	let f = new sc.CallIncrementContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcCallIncrement(ctx, f);
	ctx.log("inccounter.funcCallIncrement ok");
}

function funcCallIncrementRecurse5xThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcCallIncrementRecurse5x");
	let f = new sc.CallIncrementRecurse5xContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcCallIncrementRecurse5x(ctx, f);
	ctx.log("inccounter.funcCallIncrementRecurse5x ok");
}

function funcEndlessLoopThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcEndlessLoop");
	let f = new sc.EndlessLoopContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcEndlessLoop(ctx, f);
	ctx.log("inccounter.funcEndlessLoop ok");
}

function funcIncrementThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcIncrement");
	let f = new sc.IncrementContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcIncrement(ctx, f);
	ctx.log("inccounter.funcIncrement ok");
}

function funcIncrementWithDelayThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcIncrementWithDelay");
	let f = new sc.IncrementWithDelayContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	ctx.require(f.params.delay().exists(), "missing mandatory delay");
	sc.funcIncrementWithDelay(ctx, f);
	ctx.log("inccounter.funcIncrementWithDelay ok");
}

function funcInitThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcInit");
	let f = new sc.InitContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcInit(ctx, f);
	ctx.log("inccounter.funcInit ok");
}

function funcLocalStateInternalCallThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcLocalStateInternalCall");
	let f = new sc.LocalStateInternalCallContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcLocalStateInternalCall(ctx, f);
	ctx.log("inccounter.funcLocalStateInternalCall ok");
}

function funcLocalStatePostThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcLocalStatePost");
	let f = new sc.LocalStatePostContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcLocalStatePost(ctx, f);
	ctx.log("inccounter.funcLocalStatePost ok");
}

function funcLocalStateSandboxCallThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcLocalStateSandboxCall");
	let f = new sc.LocalStateSandboxCallContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcLocalStateSandboxCall(ctx, f);
	ctx.log("inccounter.funcLocalStateSandboxCall ok");
}

function funcPostIncrementThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcPostIncrement");
	let f = new sc.PostIncrementContext();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcPostIncrement(ctx, f);
	ctx.log("inccounter.funcPostIncrement ok");
}

function funcRepeatManyThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcRepeatMany");
	let f = new sc.RepeatManyContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcRepeatMany(ctx, f);
	ctx.log("inccounter.funcRepeatMany ok");
}

function funcTestLeb128Thunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcTestLeb128");
	let f = new sc.TestLeb128Context();
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcTestLeb128(ctx, f);
	ctx.log("inccounter.funcTestLeb128 ok");
}

function funcWhenMustIncrementThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("inccounter.funcWhenMustIncrement");
	let f = new sc.WhenMustIncrementContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcWhenMustIncrement(ctx, f);
	ctx.log("inccounter.funcWhenMustIncrement ok");
}

function viewGetCounterThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("inccounter.viewGetCounter");
	let f = new sc.GetCounterContext();
    f.results.mapID = wasmlib.OBJ_ID_RESULTS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.viewGetCounter(ctx, f);
	ctx.log("inccounter.viewGetCounter ok");
}
