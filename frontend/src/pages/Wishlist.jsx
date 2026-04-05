import { useEffect, useState } from "react";
import api from "../utils/axios";
import { useNavigate } from "react-router-dom";
import { useCart } from "../context/CartContext";

const Wishlist = () => {
  const [wishlist, setWishlist] = useState([]);
  const navigate = useNavigate();
  const { fetchCart } = useCart(); // ✅ FIXED

  useEffect(() => {
    fetchWishlist();
  }, []);

  const fetchWishlist = async () => {
    try {
      const res = await api.get("/wishlist");
      setWishlist(res.data || []);
    } catch (err) {
      console.error("Error fetching wishlist", err);
    }
  };

  // ❌ Remove
  const handleRemove = async (productId) => {
    try {
      await api.delete(`/wishlist/remove?product_id=${productId}`);
      window.dispatchEvent(new Event("wishlistUpdated"));
      fetchWishlist();
    } catch (err) {
      console.error(err);
    }
  };

  // 🛒 Move to cart
  const handleMoveToCart = async (productId) => {
    try {
      await api.post("/wishlist/move-to-cart", {
        product_id: productId,
      });

      // 🔥 IMPORTANT FIX
      await fetchCart(); // ✅ Now this works

      window.dispatchEvent(new Event("wishlistUpdated"));
      fetchWishlist();

      alert("Moved to cart ✅");
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="wishlist p-6 mt-20">
      <h2 className="text-2xl font-bold mb-6">My Wishlist ❤️</h2>

      {wishlist.length === 0 ? (
        <p className="text-gray-500">Your wishlist is empty</p>
      ) : (
        <div className="grid md:grid-cols-4 gap-6">
          {wishlist.map((item) => (
            <div
              key={item.product_id}
              className="bg-white p-4 rounded shadow hover:shadow-md"
            >
              <div className="productimage">
                <img
                  src={
                    item.image
                      ? "http://localhost:8080" + item.image
                      : "https://via.placeholder.com/200"
                  }
                  alt={item.name}
                  className="productimage h-40 w-full object-cover rounded"
                  onClick={() => navigate(`/product/${item.product_id}`)}
                />
              </div>

              <h3 className="wishlist_details mt-2 font-semibold">{item.name}</h3>
              <p className="text-green-600">
                ₹{Number(item.price).toLocaleString("en-IN")}
              </p>

              <div className="flex gap-2 mt-3">
                <button
                  onClick={() => handleMoveToCart(item.product_id)}
                  className="flex-1 bg-blue-600 text-white py-1 rounded hover:bg-blue-700"
                >
                  Move to Cart
                </button>

                <button
                  onClick={() => handleRemove(item.product_id)}
                  className="flex-1 bg-red-500 text-white py-1 rounded hover:bg-red-600"
                >
                  Remove
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Wishlist;