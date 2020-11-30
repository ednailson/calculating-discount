import {Info, Discount} from "../../proto/discount_pb";
import {client} from "./grpc";
import {Metadata} from "grpc";

export default function calculateDiscount(user_id: string, product_id: string) {
    return new Promise<Discount>(((resolve, reject) => {
        const request = new Info()
        request.setUserId(user_id)
        request.setProductId(product_id)
        client.calculateDiscount(request, new Metadata(), {deadline: new Date(Date.now() + 100)}, (err, discount) => {
            if (err) reject(err)
            else resolve(discount)
        })
    }))
}