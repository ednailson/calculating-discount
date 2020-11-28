import {Request, Response} from "express"
import db from '../database/connection'
import Product from "./interfaces/product";

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
        products.forEach((product => {
            const commandProduct: Product = {
                description: product.description,
                title: product.title,
                price_in_cents: product.price_in_cents,
                id: product._key,
                discount: {
                    percentage: 0,
                    value_in_cents: 0
                }
            }
            commandProducts.push(commandProduct)
        }))
        return response.status(200).json(commandProducts)
    }
}