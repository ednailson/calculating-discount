import Discount from "./discount";

export default interface Product {
    id: string
    price_in_cents: number
    title: string
    description: string
    discount: Discount
}