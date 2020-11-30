// package: discount
// file: discount.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as discount_pb from "./discount_pb";

interface IDiscountServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    calculateDiscount: IDiscountServiceService_ICalculateDiscount;
}

interface IDiscountServiceService_ICalculateDiscount extends grpc.MethodDefinition<discount_pb.Info, discount_pb.Discount> {
    path: "/discount.DiscountService/CalculateDiscount";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<discount_pb.Info>;
    requestDeserialize: grpc.deserialize<discount_pb.Info>;
    responseSerialize: grpc.serialize<discount_pb.Discount>;
    responseDeserialize: grpc.deserialize<discount_pb.Discount>;
}

export const DiscountServiceService: IDiscountServiceService;

export interface IDiscountServiceServer {
    calculateDiscount: grpc.handleUnaryCall<discount_pb.Info, discount_pb.Discount>;
}

export interface IDiscountServiceClient {
    calculateDiscount(request: discount_pb.Info, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
    calculateDiscount(request: discount_pb.Info, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
    calculateDiscount(request: discount_pb.Info, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
}

export class DiscountServiceClient extends grpc.Client implements IDiscountServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public calculateDiscount(request: discount_pb.Info, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
    public calculateDiscount(request: discount_pb.Info, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
    public calculateDiscount(request: discount_pb.Info, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_pb.Discount) => void): grpc.ClientUnaryCall;
}
