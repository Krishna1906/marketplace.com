import { useEffect, useState } from "react";
import { useCart } from "../context/CartContext";
import { getProductById } from "../api/product";
import { getProductImages } from "../api/image";
import { useNavigate } from "react-router-dom";

const Cart = () => {
  const { cart, increaseQty, decreaseQty, removeItem, loading } = useCart();
  const [cartItems, setCartItems] = useState([]);
  const navigate = useNavigate();

  // 🔄 Load product + image details
  useEffect(() => {
    const loadCartDetails = async () => {
      try {
        const detailedItems = await Promise.all(
          (cart || []).map(async (item) => {
            const productRes = await getProductById(item.product_id);
            const imageRes = await getProductImages(item.product_id);

            const image =
              imageRes.data?.length > 0
                ? "http://localhost:8080" + imageRes.data[0].image_url
                : "https://via.placeholder.com/150";

            return {
              ...productRes.data,
              quantity: item.quantity,
              product_id: item.product_id,
              image,
            };
          })
        );

        setCartItems(detailedItems);
      } catch (err) {
        console.error("Error loading cart:", err);
      }
    };

    if (cart.length > 0) loadCartDetails();
    else setCartItems([]);
  }, [cart]);

  // 💰 Total price
  const totalPrice = cartItems.reduce(
    (sum, item) => sum + item.price * item.quantity,
    0
  );

  // ✅ Place Order → Redirect
  const handlePlaceOrder = () => {
    navigate("/order-summary");
  };

  // 🛒 Empty cart
  if (cartItems.length === 0) {
    return <p className="p-6 text-center">Your cart is empty 🛒</p>;
  }

  return (
    <div className="bg-gray-100 min-h-screen p-6">
      <div className="grid md:grid-cols-3 gap-6">

        {/* LEFT */}
        <div className="md:col-span-2 space-y-4">
          {cartItems.map((item) => (
            <div key={item.product_id} className="bg-white p-4 rounded shadow flex gap-4">
              
              <img
                src={item.image}
                alt={item.name}
                onClick={() => navigate(`/product/${item.product_id}`)}
                className="w-32 h-32 object-cover rounded cursor-pointer"
              />

              <div className="flex-1">
                <h2
                  onClick={() => navigate(`/product/${item.product_id}`)}
                  className="font-semibold text-lg cursor-pointer"
                >
                  {item.name}
                </h2>

                <p className="text-gray-500 text-sm mt-1">
                  {item.description?.slice(0, 80)}...
                </p>

                <p className="text-green-600 font-bold mt-2">
                  ₹{item.price.toLocaleString("en-IN")}
                </p>

                <div className="flex items-center gap-3 mt-3">
                  <button onClick={() => decreaseQty(item.product_id)}>-</button>
                  <span>{item.quantity}</span>
                  <button onClick={() => increaseQty(item.product_id)}>+</button>
                </div>

                <button
                  onClick={() => removeItem(item.product_id)}
                  className="text-red-500 mt-3 text-sm"
                >
                  REMOVE
                </button>
              </div>
            </div>
          ))}
        </div>

        {/* RIGHT */}
        <div className="bg-white p-4 rounded shadow h-fit">
          <h2 className="font-bold text-lg mb-4">Price Details</h2>

          <div className="flex justify-between">
            <span>Total Items</span>
            <span>{cartItems.length}</span>
          </div>

          <div className="flex justify-between mt-2">
            <span>Total Price</span>
            <span>₹{totalPrice.toLocaleString("en-IN")}</span>
          </div>

          <hr className="my-3" />

          <div className="flex justify-between font-bold text-lg">
            <span>Amount Payable</span>
            <span>₹{totalPrice.toLocaleString("en-IN")}</span>
          </div>

          <button
            onClick={handlePlaceOrder}
            className="mt-4 orderbutton w-full bg-orange-500 text-white py-2 rounded"
          >
            Place Order
          </button>
        </div>

      </div>
    </div>
  );
};

export default Cart;