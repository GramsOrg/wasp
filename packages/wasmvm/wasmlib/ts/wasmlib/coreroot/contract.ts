// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmlib from '../index';
import * as sc from './index';

export class DeployContractCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableDeployContractParams = new sc.MutableDeployContractParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncCallContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncDeployContract);
    }
}

export class GrantDeployPermissionCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableGrantDeployPermissionParams = new sc.MutableGrantDeployPermissionParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncCallContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncGrantDeployPermission);
    }
}

export class RequireDeployPermissionsCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableRequireDeployPermissionsParams = new sc.MutableRequireDeployPermissionsParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncCallContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncRequireDeployPermissions);
    }
}

export class RevokeDeployPermissionCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableRevokeDeployPermissionParams = new sc.MutableRevokeDeployPermissionParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncCallContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncRevokeDeployPermission);
    }
}

export class SubscribeBlockContextCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableSubscribeBlockContextParams = new sc.MutableSubscribeBlockContextParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncCallContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncSubscribeBlockContext);
    }
}

export class FindContractCall {
    func:    wasmlib.ScView;
    params:  sc.MutableFindContractParams = new sc.MutableFindContractParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableFindContractResults = new sc.ImmutableFindContractResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewFindContract);
    }
}

export class GetContractRecordsCall {
    func:    wasmlib.ScView;
    results: sc.ImmutableGetContractRecordsResults = new sc.ImmutableGetContractRecordsResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetContractRecords);
    }
}

export class ScFuncs {
    static deployContract(ctx: wasmlib.ScFuncCallContext): DeployContractCall {
        const f = new DeployContractCall(ctx);
        f.params = new sc.MutableDeployContractParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    static grantDeployPermission(ctx: wasmlib.ScFuncCallContext): GrantDeployPermissionCall {
        const f = new GrantDeployPermissionCall(ctx);
        f.params = new sc.MutableGrantDeployPermissionParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    static requireDeployPermissions(ctx: wasmlib.ScFuncCallContext): RequireDeployPermissionsCall {
        const f = new RequireDeployPermissionsCall(ctx);
        f.params = new sc.MutableRequireDeployPermissionsParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    static revokeDeployPermission(ctx: wasmlib.ScFuncCallContext): RevokeDeployPermissionCall {
        const f = new RevokeDeployPermissionCall(ctx);
        f.params = new sc.MutableRevokeDeployPermissionParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    static subscribeBlockContext(ctx: wasmlib.ScFuncCallContext): SubscribeBlockContextCall {
        const f = new SubscribeBlockContextCall(ctx);
        f.params = new sc.MutableSubscribeBlockContextParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    static findContract(ctx: wasmlib.ScViewCallContext): FindContractCall {
        const f = new FindContractCall(ctx);
        f.params = new sc.MutableFindContractParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableFindContractResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getContractRecords(ctx: wasmlib.ScViewCallContext): GetContractRecordsCall {
        const f = new GetContractRecordsCall(ctx);
        f.results = new sc.ImmutableGetContractRecordsResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }
}
