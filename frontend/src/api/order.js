import axios from "axios";

const BASE_URL = "http://localhost:8080/api";

// 🧾 Get order summary
export const getOrderSummary = () => {
  return axios.get(`${BASE_URL}/ordersummary`);
};

// 🛒 Place order (if backend supports)
export const placeOrder = () => {
  return axios.post(`${BASE_URL}/order`);
};