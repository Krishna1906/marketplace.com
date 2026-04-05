import axios from "axios";

const BASE_URL = "http://localhost:8080/api";
const getToken = () => localStorage.getItem("token");

// ✅ GET CART
export const getCart = async () => {
  return await axios.get(`${BASE_URL}/cart`, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
};

// ✅ ADD TO CART (POST)
export const addToCart = async (product_id, quantity = 1) => {
  return await axios.post(
    `${BASE_URL}/cart`,
    { product_id, quantity },
    {
      headers: { Authorization: `Bearer ${getToken()}` },
    }
  );
};

// ✅ UPDATE CART (SET FINAL QUANTITY)
export const updateCartItem = async (product_id, quantity) => {
  return await axios.put(
    `${BASE_URL}/cart/update/${product_id}`,
    { quantity },
    {
      headers: { Authorization: `Bearer ${getToken()}` },
    }
  );
};

// ✅ DELETE ITEM
export const deleteCartItem = async (product_id) => {
  return await axios.delete(`${BASE_URL}/cart/${product_id}`, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
};