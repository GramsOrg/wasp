// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from '../wasmtypes';
import * as sc from './index';

export class ImmutableStoreBlobResults extends wasmtypes.ScProxy {
    // calculated hash of blob set
    hash(): wasmtypes.ScImmutableHash {
        return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ResultHash));
    }
}

export class MutableStoreBlobResults extends wasmtypes.ScProxy {
    // calculated hash of blob set
    hash(): wasmtypes.ScMutableHash {
        return new wasmtypes.ScMutableHash(this.proxy.root(sc.ResultHash));
    }
}

export class ImmutableGetBlobFieldResults extends wasmtypes.ScProxy {
    // blob data
    bytes(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultBytes));
    }
}

export class MutableGetBlobFieldResults extends wasmtypes.ScProxy {
    // blob data
    bytes(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultBytes));
    }
}

export class MapStringToImmutableInt32 extends wasmtypes.ScProxy {

    getInt32(key: string): wasmtypes.ScImmutableInt32 {
        return new wasmtypes.ScImmutableInt32(this.proxy.key(wasmtypes.stringToBytes(key)));
    }
}

export class ImmutableGetBlobInfoResults extends wasmtypes.ScProxy {
    // size for each named blob
    blobSizes(): sc.MapStringToImmutableInt32 {
        return new sc.MapStringToImmutableInt32(this.proxy);
    }
}

export class MapStringToMutableInt32 extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getInt32(key: string): wasmtypes.ScMutableInt32 {
        return new wasmtypes.ScMutableInt32(this.proxy.key(wasmtypes.stringToBytes(key)));
    }
}

export class MutableGetBlobInfoResults extends wasmtypes.ScProxy {
    // size for each named blob
    blobSizes(): sc.MapStringToMutableInt32 {
        return new sc.MapStringToMutableInt32(this.proxy);
    }
}

export class MapHashToImmutableInt32 extends wasmtypes.ScProxy {

    getInt32(key: wasmtypes.ScHash): wasmtypes.ScImmutableInt32 {
        return new wasmtypes.ScImmutableInt32(this.proxy.key(wasmtypes.hashToBytes(key)));
    }
}

export class ImmutableListBlobsResults extends wasmtypes.ScProxy {
    // total size for each blob set
    blobSizes(): sc.MapHashToImmutableInt32 {
        return new sc.MapHashToImmutableInt32(this.proxy);
    }
}

export class MapHashToMutableInt32 extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getInt32(key: wasmtypes.ScHash): wasmtypes.ScMutableInt32 {
        return new wasmtypes.ScMutableInt32(this.proxy.key(wasmtypes.hashToBytes(key)));
    }
}

export class MutableListBlobsResults extends wasmtypes.ScProxy {
    // total size for each blob set
    blobSizes(): sc.MapHashToMutableInt32 {
        return new sc.MapHashToMutableInt32(this.proxy);
    }
}
