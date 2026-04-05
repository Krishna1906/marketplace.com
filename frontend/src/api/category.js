import API from "./auth";

// Get all categories
export const getCategories = () => {
  return API.get("/categories");
};

// Get products by category ID
export const getCategoryProducts = (id) => {
  return API.get(`/categories/${id}`);
};