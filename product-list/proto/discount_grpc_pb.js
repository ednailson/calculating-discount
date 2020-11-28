// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var discount_pb = require('./discount_pb.js');

function serialize_discount_Discount(arg) {
  if (!(arg instanceof discount_pb.Discount)) {
    throw new Error('Expected argument of type discount.Discount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_discount_Discount(buffer_arg) {
  return discount_pb.Discount.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_discount_Info(arg) {
  if (!(arg instanceof discount_pb.Info)) {
    throw new Error('Expected argument of type discount.Info');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_discount_Info(buffer_arg) {
  return discount_pb.Info.deserializeBinary(new Uint8Array(buffer_arg));
}


var DiscountServiceService = exports.DiscountServiceService = {
  calculateDiscount: {
    path: '/discount.DiscountService/CalculateDiscount',
    requestStream: false,
    responseStream: false,
    requestType: discount_pb.Info,
    responseType: discount_pb.Discount,
    requestSerialize: serialize_discount_Info,
    requestDeserialize: deserialize_discount_Info,
    responseSerialize: serialize_discount_Discount,
    responseDeserialize: deserialize_discount_Discount,
  },
};

exports.DiscountServiceClient = grpc.makeGenericClientConstructor(DiscountServiceService);
