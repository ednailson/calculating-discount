import express from 'express'
import ProductController from "./controllers/ProductController"

const routes = express.Router()
const productController = new ProductController()

routes.get("/product", productController.read)

export default routes