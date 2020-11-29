import {Request, Response} from "express"
import db from '../database/connection'
import Product from "./interfaces/product";
import calculateDiscount from "../client/calculate-discount";
import {Discount} from "../../proto/discount_pb";

export default class ProductController {
    async read(request: Request, response: Response) {
        const receivedUser = request.query["user_id"]
        let userId = ""
        if (receivedUser) {
            userId = receivedUser.toString()
        }

        const productColl = await db.collection("product-collection")

        const cursor = await productColl.all()
        const products = await cursor.all()


        let commandProducts: Product[] = []
        for (let i = 0; i < products.length; i++) {
            const product = products[i]
            let discount: Discount = new Discount()
            await calculateDiscount(userId, product._key).then((result) => {
                discount = result
            }).catch(() => {})
            const commandProduct: Product = {
                description: product.description,
                title: product.title,
                price_in_cents: product.price_in_cents,
                id: product._key,
                discount: {
                    percentage: discount.getPercentage(),
                    value_in_cents: discount.getValueInCents()
                }
            }
            commandProducts.push(commandProduct)
        }
        return response.status(200).json(commandProducts)
    }
}