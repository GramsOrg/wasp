// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";

export const ScName        = "testcore";
export const ScDescription = "Core test for ISCP wasmlib Rust/Wasm library";
export const HScName       = new wasmlib.ScHname(0x370d33ad);

export const ParamAddress         = "address";
export const ParamAgentID         = "agentID";
export const ParamCaller          = "caller";
export const ParamChainID         = "chainID";
export const ParamChainOwnerID    = "chainOwnerID";
export const ParamContractCreator = "contractCreator";
export const ParamContractID      = "contractID";
export const ParamCounter         = "counter";
export const ParamFail            = "initFailParam";
export const ParamHash            = "Hash";
export const ParamHname           = "Hname";
export const ParamHnameContract   = "hnameContract";
export const ParamHnameEP         = "hnameEP";
export const ParamHnameZero       = "Hname-0";
export const ParamInt64           = "int64";
export const ParamInt64Zero       = "int64-0";
export const ParamIntValue        = "intParamValue";
export const ParamName            = "intParamName";
export const ParamProgHash        = "progHash";
export const ParamString          = "string";
export const ParamStringZero      = "string-0";
export const ParamVarName         = "varName";

export const ResultChainOwnerID = "chainOwnerID";
export const ResultCounter      = "counter";
export const ResultIntValue     = "intParamValue";
export const ResultMintedColor  = "mintedColor";
export const ResultMintedSupply = "mintedSupply";
export const ResultSandboxCall  = "sandboxCall";
export const ResultValues       = "this";
export const ResultVars         = "this";

export const StateCounter      = "counter";
export const StateHnameEP      = "hnameEP";
export const StateInts         = "ints";
export const StateMintedColor  = "mintedColor";
export const StateMintedSupply = "mintedSupply";

export const FuncCallOnChain                 = "callOnChain";
export const FuncCheckContextFromFullEP      = "checkContextFromFullEP";
export const FuncDoNothing                   = "doNothing";
export const FuncGetMintedSupply             = "getMintedSupply";
export const FuncIncCounter                  = "incCounter";
export const FuncInit                        = "init";
export const FuncPassTypesFull               = "passTypesFull";
export const FuncRunRecursion                = "runRecursion";
export const FuncSendToAddress               = "sendToAddress";
export const FuncSetInt                      = "setInt";
export const FuncSpawn                       = "spawn";
export const FuncTestBlockContext1           = "testBlockContext1";
export const FuncTestBlockContext2           = "testBlockContext2";
export const FuncTestCallPanicFullEP         = "testCallPanicFullEP";
export const FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull";
export const FuncTestChainOwnerIDFull        = "testChainOwnerIDFull";
export const FuncTestEventLogDeploy          = "testEventLogDeploy";
export const FuncTestEventLogEventData       = "testEventLogEventData";
export const FuncTestEventLogGenericData     = "testEventLogGenericData";
export const FuncTestPanicFullEP             = "testPanicFullEP";
export const FuncWithdrawToChain             = "withdrawToChain";
export const ViewCheckContextFromViewEP      = "checkContextFromViewEP";
export const ViewFibonacci                   = "fibonacci";
export const ViewGetCounter                  = "getCounter";
export const ViewGetInt                      = "getInt";
export const ViewGetStringValue              = "getStringValue";
export const ViewJustView                    = "justView";
export const ViewPassTypesView               = "passTypesView";
export const ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView";
export const ViewTestChainOwnerIDView        = "testChainOwnerIDView";
export const ViewTestPanicViewEP             = "testPanicViewEP";
export const ViewTestSandboxCall             = "testSandboxCall";

export const HFuncCallOnChain                 = new wasmlib.ScHname(0x95a3d123);
export const HFuncCheckContextFromFullEP      = new wasmlib.ScHname(0xa56c24ba);
export const HFuncDoNothing                   = new wasmlib.ScHname(0xdda4a6de);
export const HFuncGetMintedSupply             = new wasmlib.ScHname(0x0c2d113c);
export const HFuncIncCounter                  = new wasmlib.ScHname(0x7b287419);
export const HFuncInit                        = new wasmlib.ScHname(0x1f44d644);
export const HFuncPassTypesFull               = new wasmlib.ScHname(0x733ea0ea);
export const HFuncRunRecursion                = new wasmlib.ScHname(0x833425fd);
export const HFuncSendToAddress               = new wasmlib.ScHname(0x63ce4634);
export const HFuncSetInt                      = new wasmlib.ScHname(0x62056f74);
export const HFuncSpawn                       = new wasmlib.ScHname(0xec929d12);
export const HFuncTestBlockContext1           = new wasmlib.ScHname(0x796d4136);
export const HFuncTestBlockContext2           = new wasmlib.ScHname(0x758b0452);
export const HFuncTestCallPanicFullEP         = new wasmlib.ScHname(0x4c878834);
export const HFuncTestCallPanicViewEPFromFull = new wasmlib.ScHname(0xfd7e8c1d);
export const HFuncTestChainOwnerIDFull        = new wasmlib.ScHname(0x2aff1167);
export const HFuncTestEventLogDeploy          = new wasmlib.ScHname(0x96ff760a);
export const HFuncTestEventLogEventData       = new wasmlib.ScHname(0x0efcf939);
export const HFuncTestEventLogGenericData     = new wasmlib.ScHname(0x6a16629d);
export const HFuncTestPanicFullEP             = new wasmlib.ScHname(0x24fdef07);
export const HFuncWithdrawToChain             = new wasmlib.ScHname(0x437bc026);
export const HViewCheckContextFromViewEP      = new wasmlib.ScHname(0x88ff0167);
export const HViewFibonacci                   = new wasmlib.ScHname(0x7940873c);
export const HViewGetCounter                  = new wasmlib.ScHname(0xb423e607);
export const HViewGetInt                      = new wasmlib.ScHname(0x1887e5ef);
export const HViewGetStringValue              = new wasmlib.ScHname(0xcf0a4d32);
export const HViewJustView                    = new wasmlib.ScHname(0x33b8972e);
export const HViewPassTypesView               = new wasmlib.ScHname(0x1a5b87ea);
export const HViewTestCallPanicViewEPFromView = new wasmlib.ScHname(0x91b10c99);
export const HViewTestChainOwnerIDView        = new wasmlib.ScHname(0x26586c33);
export const HViewTestPanicViewEP             = new wasmlib.ScHname(0x22bc4d72);
export const HViewTestSandboxCall             = new wasmlib.ScHname(0x42d72b63);
