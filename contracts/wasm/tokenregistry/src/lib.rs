// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use tokenregistry::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod params;
mod state;
mod structs;
mod tokenregistry;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_MINT_SUPPLY,        func_mint_supply_thunk);
    exports.add_func(FUNC_TRANSFER_OWNERSHIP, func_transfer_ownership_thunk);
    exports.add_func(FUNC_UPDATE_METADATA,    func_update_metadata_thunk);
    exports.add_view(VIEW_GET_INFO,           view_get_info_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct MintSupplyContext {
	params: ImmutableMintSupplyParams,
	state: MutableTokenRegistryState,
}

fn func_mint_supply_thunk(ctx: &ScFuncContext) {
	ctx.log("tokenregistry.funcMintSupply");
	let f = MintSupplyContext {
		params: ImmutableMintSupplyParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableTokenRegistryState {
			id: OBJ_ID_STATE,
		},
	};
	func_mint_supply(ctx, &f);
	ctx.log("tokenregistry.funcMintSupply ok");
}

pub struct TransferOwnershipContext {
	params: ImmutableTransferOwnershipParams,
	state: MutableTokenRegistryState,
}

fn func_transfer_ownership_thunk(ctx: &ScFuncContext) {
	ctx.log("tokenregistry.funcTransferOwnership");
    // TODO the one who can transfer token ownership
		ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

	let f = TransferOwnershipContext {
		params: ImmutableTransferOwnershipParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableTokenRegistryState {
			id: OBJ_ID_STATE,
		},
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	func_transfer_ownership(ctx, &f);
	ctx.log("tokenregistry.funcTransferOwnership ok");
}

pub struct UpdateMetadataContext {
	params: ImmutableUpdateMetadataParams,
	state: MutableTokenRegistryState,
}

fn func_update_metadata_thunk(ctx: &ScFuncContext) {
	ctx.log("tokenregistry.funcUpdateMetadata");
    // TODO the one who can change the token info
		ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

	let f = UpdateMetadataContext {
		params: ImmutableUpdateMetadataParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableTokenRegistryState {
			id: OBJ_ID_STATE,
		},
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	func_update_metadata(ctx, &f);
	ctx.log("tokenregistry.funcUpdateMetadata ok");
}

pub struct GetInfoContext {
	params: ImmutableGetInfoParams,
	state: ImmutableTokenRegistryState,
}

fn view_get_info_thunk(ctx: &ScViewContext) {
	ctx.log("tokenregistry.viewGetInfo");
	let f = GetInfoContext {
		params: ImmutableGetInfoParams {
			id: OBJ_ID_PARAMS,
		},
		state: ImmutableTokenRegistryState {
			id: OBJ_ID_STATE,
		},
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	view_get_info(ctx, &f);
	ctx.log("tokenregistry.viewGetInfo ok");
}
