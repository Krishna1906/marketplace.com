import { useEffect, useState } from "react";
import { getCategories } from "../api/category";
import { getProducts } from "../api/product";
import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../context/CartContext";
import api from "../utils/axios";
import { FaHeart } from "react-icons/fa";


const Navbar = () => {
  const [categories, setCategories] = useState([]);
  const navigate = useNavigate();
  const [wishlistCount, setWishlistCount] = useState(0);
  const { cart } = useCart();

  // ✅ Check login
  const token = localStorage.getItem("token");
  const isLoggedIn = !!token;

  useEffect(() => {
  const fetchData = async () => {
    try {
      const [catRes, prodRes] = await Promise.all([
        getCategories(),
        getProducts(),
      ]);

      const allCategories = catRes.data;
      const products = prodRes.data;

      const categoryIdsWithProducts = new Set(
        products.map((p) => p.category_id)
      );

      const filteredCategories = allCategories.filter((cat) =>
        categoryIdsWithProducts.has(cat.id)
      );

      setCategories(filteredCategories);
    } catch (err) {
      console.error("Error loading categories/products", err);
    }
  };

  fetchData();

  // ✅ Fetch wishlist
  const fetchWishlist = async () => {
    try {
      if (!token) return;
      const res = await api.get("/wishlist").catch(() => ({ data: [] }));
      setWishlistCount(res.data?.length || 0);
    } catch (err) {
      console.error(err);
    }
  };

  fetchWishlist();

  // ✅ LISTEN FOR GLOBAL UPDATE
  const handleWishlistUpdate = () => {
    fetchWishlist();
  };

  window.addEventListener("wishlistUpdated", handleWishlistUpdate);

  return () => {
    window.removeEventListener("wishlistUpdated", handleWishlistUpdate);
  };

}, []);

  const totalQty =
    cart?.reduce((sum, item) => sum + (item.quantity || 0), 0) || 0;

  // ✅ Logout function
  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
    window.location.reload(); // refresh navbar
  };

  return (
    <div className="nav fixed top-0 left-0 w-full bg-blue-600 text-white shadow z-50">

      {/* Top */}
      <div className="mainnav flex items-center justify-between px-6 py-3">

        {/* Logo */}
        <Link to="/" className="text-xl font-bold">
          <i>MARKETPLACE.COM</i>
        </Link>

        {/* Search */}
        <input
          type="text"
          placeholder="Search..."
          className="w-1/2 px-4 py-1 rounded text-black"
        />

        {/* Right Section */}
        <div className="flex items-center gap-6">

          {/* <Link to="/wishlist">
          Wishlist ({wishlist.length})
        </Link> */}

        {isLoggedIn && (
  <div
    onClick={() => navigate("/wishlist")}
    className="cursor-pointer relative flex items-center"
  >
    <FaHeart className="text-white text-lg" />

    {wishlistCount > 0 && (
      <span className="absolute -top-2 -right-3 bg-red-500 text-white text-xs px-2 py-0.5 rounded-full">
        {wishlistCount}
      </span>
    )}
  </div>
)}

          {/* 🛒 Cart */}
          <div
            onClick={() => navigate("/cart")}
            className="cursor-pointer relative"
          >
            🛒
            {totalQty > 0 && (
              <span className="absolute -top-2 -right-3 bg-red-500 text-white text-xs px-2 py-0.5 rounded-full">
                {totalQty}
              </span>
            )}
          </div>

          {/* ✅ Auth Section */}
          {!isLoggedIn ? (
            <button
              onClick={() => navigate("/login")}
              className="bg-white text-blue-600 px-3 py-1 rounded"
            >
              Login
            </button>
          ) : (
            <>
              <button
                onClick={() => navigate("/orders")}
                className="hover:text-yellow-300"
              >
                My Orders
              </button>

              <button
                onClick={handleLogout}
                className="bg-red-500 px-3 py-1 rounded"
              >
                Logout
              </button>
            </>
          )}
        </div>
      </div>

      {/* Categories */}
      <div className="categories flex gap-6 px-6 pb-2 overflow-x-auto text-sm">
        {categories.map((cat) => (
          <div
            key={cat.id}
            onClick={() => navigate(`/?category=${cat.id}`)}
            className="cursor-pointer hover:text-yellow-300 whitespace-nowrap"
          >
            {cat.name}
          </div>
        ))}
      </div>

    </div>
  );
};

export default Navbar;