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
    exports.addFunc(sc.FuncDonate,       funcDonateThunk);
    exports.addFunc(sc.FuncWithdraw,     funcWithdrawThunk);
    exports.addView(sc.ViewDonation,     viewDonationThunk);
    exports.addView(sc.ViewDonationInfo, viewDonationInfoThunk);

    for (let i = 0; i < sc.keyMap.length; i++) {
        sc.idxMap[i] = wasmlib.Key32.fromString(sc.keyMap[i]);
    }
}

function funcDonateThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("donatewithfeedback.funcDonate");
	let f = new sc.DonateContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcDonate(ctx, f);
	ctx.log("donatewithfeedback.funcDonate ok");
}

function funcWithdrawThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("donatewithfeedback.funcWithdraw");
    // only SC creator can withdraw donated funds
	ctx.require(ctx.caller().equals(ctx.contractCreator()), "no permission");

	let f = new sc.WithdrawContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.funcWithdraw(ctx, f);
	ctx.log("donatewithfeedback.funcWithdraw ok");
}

function viewDonationThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("donatewithfeedback.viewDonation");
	let f = new sc.DonationContext();
    f.params.mapID = wasmlib.OBJ_ID_PARAMS;
    f.results.mapID = wasmlib.OBJ_ID_RESULTS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	ctx.require(f.params.nr().exists(), "missing mandatory nr");
	sc.viewDonation(ctx, f);
	ctx.log("donatewithfeedback.viewDonation ok");
}

function viewDonationInfoThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("donatewithfeedback.viewDonationInfo");
	let f = new sc.DonationInfoContext();
    f.results.mapID = wasmlib.OBJ_ID_RESULTS;
    f.state.mapID = wasmlib.OBJ_ID_STATE;
	sc.viewDonationInfo(ctx, f);
	ctx.log("donatewithfeedback.viewDonationInfo ok");
}
