import { DiscountServiceClient } from "../../proto/discount_grpc_pb"
import { credentials } from "grpc"

const port = 9000

export const client = new DiscountServiceClient(
    `discount-calculator.service.com.br:${port}`,
    credentials.createInsecure()
)