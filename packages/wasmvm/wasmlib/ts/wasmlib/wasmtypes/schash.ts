// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import {panic} from "../sandbox";
import {base58Encode, WasmDecoder, WasmEncoder, zeroes} from "./codec";
import {Proxy} from "./proxy";
import {bytesCompare} from "./scbytes";

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

export const ScHashLength = 32;

export class ScHash {
    id: u8[] = zeroes(ScHashLength);

    public equals(other: ScHash): bool {
        return bytesCompare(this.id, other.id) == 0;
    }

    // convert to byte array representation
    public toBytes(): u8[] {
        return hashToBytes(this);
    }

    // human-readable string representation
    public toString(): string {
        return hashToString(this);
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

export function hashDecode(dec: WasmDecoder): ScHash {
    return hashFromBytesUnchecked(dec.fixedBytes(ScHashLength));
}

export function hashEncode(enc: WasmEncoder, value: ScHash): void {
    enc.fixedBytes(value.toBytes(), ScHashLength);
}

export function hashFromBytes(buf: u8[]): ScHash {
    if (buf.length == 0) {
        return new ScHash();
    }
    if (buf.length != ScHashLength) {
        panic("invalid Hash length");
    }
    return hashFromBytesUnchecked(buf);
}

export function hashToBytes(value: ScHash): u8[] {
    return value.id;
}

export function hashToString(value: ScHash): string {
    // TODO standardize human readable string
    return base58Encode(value.id);
}

function hashFromBytesUnchecked(buf: u8[]): ScHash {
    let o = new ScHash();
    o.id = buf.slice(0);
    return o;
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

export class ScImmutableHash {
    proxy: Proxy;

    constructor(proxy: Proxy) {
        this.proxy = proxy;
    }

    exists(): bool {
        return this.proxy.exists();
    }

    toString(): string {
        return hashToString(this.value());
    }

    value(): ScHash {
        return hashFromBytes(this.proxy.get());
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

export class ScMutableHash extends ScImmutableHash {
    delete(): void {
        this.proxy.delete();
    }

    setValue(value: ScHash): void {
        this.proxy.set(hashToBytes(value));
    }
}
