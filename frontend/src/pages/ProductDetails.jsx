import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { getProductById } from "../api/product";
import { getProductImages } from "../api/image";
import { addToCart } from "../api/cart";
import { useCart } from "../context/CartContext";
import api from "../utils/axios";
import { FaHeart, FaRegHeart } from "react-icons/fa";

const ProductDetails = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { fetchCart } = useCart();

  const [product, setProduct] = useState(null);
  const [images, setImages] = useState([]);
  const [mainImage, setMainImage] = useState("");
  const [loading, setLoading] = useState(true);
  const [isWishlisted, setIsWishlisted] = useState(false);

  useEffect(() => {
    fetchData();
  }, [id]);

  const fetchData = async () => {
    try {
      setLoading(true);

      const token = localStorage.getItem("token");

      const [productRes, imageRes, wishlistRes] = await Promise.all([
        getProductById(id),
        getProductImages(id),
        token ? api.get("/wishlist") : Promise.resolve({ data: [] }),
      ]);

      setProduct(productRes.data);

      const imgs = imageRes.data.map(
        (img) => "http://localhost:8080" + img.image_url
      );

      setImages(imgs);
      if (imgs.length > 0) setMainImage(imgs[0]);

      // ✅ Check wishlist
      const wishlistData = Array.isArray(wishlistRes.data)
        ? wishlistRes.data
        : [];

      const exists = wishlistData.some(
        (item) => item.product_id === Number(id)
      );
      setIsWishlisted(exists);

    } catch (err) {
      console.error("Error fetching product:", err);
    } finally {
      setLoading(false);
    }
  };

  // ❤️ TOGGLE WISHLIST
  const handleWishlistToggle = async () => {
    const token = localStorage.getItem("token");

    if (!token) {
      alert("Please login first");
      navigate("/login");
      return;
    }

    try {
            if (isWishlisted) {
        await api.delete(`/wishlist/remove?product_id=${product.id}`);
        setIsWishlisted(false);
      } else {
        await api.post("/wishlist", { product_id: product.id });
        setIsWishlisted(true);
      }

      // ✅ TRIGGER NAVBAR UPDATE
      window.dispatchEvent(new Event("wishlistUpdated"));
    } catch (err) {
      console.error(err.response?.data || err.message);
    }
  };

  // 🛒 ADD TO CART
  const handleAddToCart = async () => {
    if (!product || product.stock === 0) return;

    const token = localStorage.getItem("token");

    if (!token) {
      alert("Please login first");
      navigate("/login");
      return;
    }

    try {
      await addToCart(product.id, 1);
      await fetchCart();
      alert("Added to cart ✅");
    } catch (err) {
      console.error(err.response?.data || err.message);
    }
  };

  // ⏳ LOADING
  if (loading) {
    return <div className="p-6 text-center">Loading product...</div>;
  }

  // ❌ NOT FOUND
  if (!product) {
    return <div className="p-6 text-center">Product not found</div>;
  }

  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <div className="bg-white p-6 rounded-lg grid md:grid-cols-2 gap-6">

        {/* LEFT SIDE */}
        <div>

          {/* MAIN IMAGE WITH ❤️ */}
          <div className="h-96 mb-4 relative group">

            {/* ❤️ Wishlist */}
            <div
              onClick={handleWishlistToggle}
              className="absolute top-3 right-3 bg-white rounded-full p-2 shadow-md cursor-pointer transition-transform duration-200 hover:scale-110 active:scale-95"
            >
              {isWishlisted ? (
                <FaHeart className="text-red-500 text-lg" />
              ) : (
                <FaRegHeart className="text-gray-600 text-lg group-hover:text-red-400" />
              )}
            </div>
              <div className="product_image">
            <img
              src={mainImage || "https://via.placeholder.com/400"}
              alt="product"
              className="h-96 productimage object-cover rounded"
            /></div>
          </div>

          {/* THUMBNAILS */}
          <div className="flex gap-2 mt-4">
            {images.map((img, i) => (
              <img
                key={i}
                src={img}
                alt={`thumb-${i}`}
                onClick={() => setMainImage(img)}
                className={`h-20 w-20 productimage cursor-pointer object-cover rounded border-2 ${
                  mainImage === img ? "border-blue-600" : "border-gray-300"
                }`}
              />
            ))}
          </div>

        </div>

        {/* RIGHT SIDE */}
        <div className="flex flex-col justify-center">
          <h1 className="text-2xl font-bold">{product.name}</h1>

          <p className="text-green-600 text-xl mt-2">
            ₹{Number(product.price).toLocaleString("en-IN")}
          </p>

          <p className="mt-3 text-gray-700">{product.description}</p>

          {/* STOCK */}
          {product.stock === 0 ? (
            <p className="text-red-600 font-semibold mt-3">
              Out of Stock
            </p>
          ) : (
            <p className="text-green-600 mt-3">
              {product.stock <= 5 ? (
                <span className="font-semibold">
                  Only {product.stock} left! Hurry up!
                </span>
              ) : (
                "In Stock"
              )}
            </p>
          )}

          {/* BUTTON */}
          <div className="flex gap-4 mt-6">
            <button
              onClick={handleAddToCart}
              disabled={product.stock === 0}
              className={`px-6 py-2 rounded text-white ${
                product.stock === 0
                  ? "bg-gray-400 cursor-not-allowed"
                  : "bg-blue-600 hover:bg-blue-700"
              }`}
            >
              {product.stock === 0 ? "Out of Stock" : "Add to Cart"}
            </button>
          </div>
        </div>

      </div>
    </div>
  );
};

export default ProductDetails;