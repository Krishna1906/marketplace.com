import API from "./auth";

// Get all images of a product
export const getProductImages = (productId) => {
  return API.get(`/images/product/${productId}`);
};