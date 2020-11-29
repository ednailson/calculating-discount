import {Request, Response} from "express"
import db from '../database/connection'
import Product from "./interfaces/product";
import calculateDiscount from "../client/calculate-discount";
import {Discount} from "../../proto/discount_pb";

export default class ProductController {
    async read(request: Request, response: Response) {
        const userId = request.query["user_id"]
        if (!userId) {
            return response.status(400).json({user_id: "user_id is required"})
        }

        const userColl = await db.collection("user-collection")
        const productColl = await db.collection("product-collection")

        const cursor = await productColl.all()
        const products = await cursor.all()

        const users = await userColl.lookupByKeys([userId.toString()])
        if (!users[0]) {
            return response.status(404).json({user_id: "user not found"})
        }

        const user = users[0]
        let commandProducts: Product[] = []
        for (let i = 0; i < products.length; i++) {
            const product = products[i]
            let discount: Discount = new Discount()
            await calculateDiscount(user._key, product._key).then((result) => {
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