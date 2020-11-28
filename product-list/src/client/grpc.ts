import { DiscountServiceClient } from "../../proto/discount_grpc_pb"
import { credentials } from "grpc"

const port = 9000

export const client = new DiscountServiceClient(
    `localhost:${port}`,
    credentials.createInsecure()
)