// package: discount
// file: discount.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Info extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): Info;

    getProductId(): string;
    setProductId(value: string): Info;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Info.AsObject;
    static toObject(includeInstance: boolean, msg: Info): Info.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Info, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Info;
    static deserializeBinaryFromReader(message: Info, reader: jspb.BinaryReader): Info;
}

export namespace Info {
    export type AsObject = {
        userId: string,
        productId: string,
    }
}

export class Discount extends jspb.Message { 
    getPercentage(): number;
    setPercentage(value: number): Discount;

    getValueInCents(): number;
    setValueInCents(value: number): Discount;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Discount.AsObject;
    static toObject(includeInstance: boolean, msg: Discount): Discount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Discount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Discount;
    static deserializeBinaryFromReader(message: Discount, reader: jspb.BinaryReader): Discount;
}

export namespace Discount {
    export type AsObject = {
        percentage: number,
        valueInCents: number,
    }
}
