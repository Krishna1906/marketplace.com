import { createContext, useContext, useEffect, useState } from "react";
import {
  getCart,
  updateCartItem,
  deleteCartItem,
} from "../api/cart";
import { useRef } from "react";

export const CartContext = createContext();
export const useCart = () => useContext(CartContext);

const CartProvider = ({ children }) => {
  const actionLock = useRef(false);
  const [cart, setCart] = useState([]);
  const [loading, setLoading] = useState(false); // ✅ ADD THIS

  const fetchCart = async () => {
    try {
      const res = await getCart();
      setCart(res.data || []);
    } catch (err) {
      console.error("Fetch cart error:", err);
      setCart([]);
    }
  };

  // ➕ Increase
  const increaseQty = async (product_id) => {
  if (actionLock.current) return;
  actionLock.current = true;

  try {
    const item = cart.find(i => i.product_id === product_id);

    // console.log("CURRENT:", item.quantity);
    // console.log("SENDING:", item.quantity + 1);

    await updateCartItem(product_id, item.quantity + 1);
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
    const item = cart.find(i => i.product_id === product_id);
    if (!item) return;

    if (item.quantity <= 1) {
      await deleteCartItem(product_id);
    } else {
      await updateCartItem(product_id, item.quantity - 1);
    }

    await fetchCart();
  } catch (err) {
    console.error(err);
  } finally {
    actionLock.current = false;
  }
};

  // ❌ Remove
  const removeItem = async (product_id) => {
    if (loading) return; // 🚫 PREVENT DOUBLE CALL
    setLoading(true);

    try {
      await deleteCartItem(product_id);
      await fetchCart();
    } catch (err) {
      console.error("Remove error:", err.response?.data || err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (localStorage.getItem("token")) {
      fetchCart();
    }
  }, []);

  return (
    <CartContext.Provider
      value={{
        cart,
        setCart,
        increaseQty,
        decreaseQty,
        removeItem,
        fetchCart,
        loading, // ✅ expose loading for UI (optional)
      }}
    >
      {children}
    </CartContext.Provider>
  );
};

export default CartProvider;