import { useEffect, useState } from "react";
import api from "../utils/axios";

const useWishlist = () => {
  const [wishlist, setWishlist] = useState([]);

  // Fetch wishlist
  const fetchWishlist = async () => {
    try {
      const res = await api.get("/wishlist");
      setWishlist(res.data);
    } catch (err) {
      console.error("Wishlist fetch error", err);
    }
  };

  useEffect(() => {
    fetchWishlist();
  }, []);

  // Check if product is in wishlist
  const isInWishlist = (productId) => {
    return wishlist.some((item) => item.product_id === productId);
  };

  // Toggle wishlist
  const toggleWishlist = async (productId) => {
    try {
      if (isInWishlist(productId)) {
        await api.delete(`/wishlist/remove?product_id=${productId}`);
      } else {
        await api.post("/wishlist", { product_id: productId });
      }
      fetchWishlist(); // refresh
    } catch (err) {
      console.error("Wishlist toggle error", err);
    }
  };

  return { wishlist, isInWishlist, toggleWishlist };
};

export default useWishlist;