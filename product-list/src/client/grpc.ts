import { DiscountServiceClient } from "../../proto/discount_grpc_pb"
import { credentials } from "grpc"
import config from "../../config.json"

export const client = new DiscountServiceClient(
    config.discount_list_host,
    credentials.createInsecure()
)