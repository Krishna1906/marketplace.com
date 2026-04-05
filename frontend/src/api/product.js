import API from "./auth";

export const getProducts = () => {
  return API.get("/products");
};

export const getProductById = (id) => {
  return API.get(`/products/${id}`);
}