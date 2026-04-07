import { createContext, useContext, useState, useEffect, useRef } from "react";
import api from "../utils/axios";

export const CartContext = createContext();
export const useCart = () => useContext(CartContext);

const CartProvider = ({ children }) => {
  const actionLock = useRef(false);
  const [cart, setCart] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchCart = async () => {
    try {
      const res = await api.get("/cart");
      setCart(res.data || []);
    } catch (err) {
      console.error("Fetch cart error:", err);
      setCart([]);
    }
  };

  const increaseQty = async (product_id) => {
    if (actionLock.current) return;
    actionLock.current = true;
    try {
      const item = cart.find((i) => i.product_id === product_id);
      await api.put("/cart", { product_id, quantity: item.quantity + 1 });
      await fetchCart();
    } catch (err) {
      console.error(err);
    } finally {
      actionLock.current = false;
    }
  };

  const decreaseQty = async (product_id) => {
    if (actionLock.current) return;
    actionLock.current = true;
    try {
      const item = cart.find((i) => i.product_id === product_id);
      if (!item) return;
      if (item.quantity <= 1) await api.delete(`/cart/${product_id}`);
      else await api.put("/cart", { product_id, quantity: item.quantity - 1 });
      await fetchCart();
    } catch (err) {
      console.error(err);
    } finally {
      actionLock.current = false;
    }
  };

  const removeItem = async (product_id) => {
    if (loading) return;
    setLoading(true);
    try {
      await api.delete(`/cart/${product_id}`);
      await fetchCart();
    } catch (err) {
      console.error(err.response?.data || err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (localStorage.getItem("token")) fetchCart();
  }, []);

  return (
    <CartContext.Provider
      value={{ cart, setCart, increaseQty, decreaseQty, removeItem, fetchCart, loading }}
    >
      {children}
    </CartContext.Provider>
  );
};

export default CartProvider;