import { useEffect, useState } from "react";
import { getProducts } from "../api/product";
import { getCategoryProducts, getCategories } from "../api/category";
import { useSearchParams } from "react-router-dom";
import ProductCard from "../components/ProductCard";
import axios from "axios";

const Home = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  const [searchParams] = useSearchParams();
  const categoryId = searchParams.get("category");
  const [categoryName, setCategoryName] = useState("");

  // ✅ BANNER STATE
  const [banners, setBanners] = useState([]);
  const [currentBanner, setCurrentBanner] = useState(0);

  // ---------------- PRODUCTS ----------------
  useEffect(() => {
    if (categoryId) {
      fetchCategoryProducts(categoryId);
      fetchCategoryName(categoryId);
    } else {
      fetchAllProducts();
      setCategoryName("");
    }
  }, [categoryId]);

  const fetchAllProducts = async () => {
    setLoading(true);
    try {
      const res = await getProducts();
      setProducts(res.data);
    } catch (err) {
      console.error("Error fetching products");
    } finally {
      setLoading(false);
    }
  };

  const fetchCategoryProducts = async (id) => {
    setLoading(true);
    try {
      const res = await getCategoryProducts(id);
      setProducts(res.data);
    } catch (err) {
      console.error("Error fetching category products");
    } finally {
      setLoading(false);
    }
  };

  const fetchCategoryName = async (id) => {
    try {
      const res = await getCategories();
      const category = res.data.find((c) => c.id == id);
      if (category) setCategoryName(category.name);
    } catch (err) {
      console.error("Error fetching category name");
    }
  };

  // ---------------- BANNERS ----------------
  useEffect(() => {
    fetchBanners();
  }, []);

  const fetchBanners = async () => {
    try {
      const res = await axios.get("http://localhost:8080/api/banner");
      setBanners(res.data || []);
    } catch (err) {
      console.error("Error fetching banners", err);
    }
  };

  // ✅ AUTO SLIDER (5 sec loop)
  useEffect(() => {
    if (banners.length === 0) return;

    const interval = setInterval(() => {
      setCurrentBanner((prev) => (prev + 1) % banners.length);
    }, 5000);

    return () => clearInterval(interval);
  }, [banners]);

  // ---------------- UI ----------------
  return (
    <div className="bg-gray-100 min-h-screen">

      {/* 🔥 DYNAMIC BANNER */}
      <div className="homebanner m-4">
        {banners.length > 0 ? (
          <img
            src={"http://localhost:8080" + banners[currentBanner].image_url}
            className="banner w-full h-[300px] object-cover rounded-lg transition-all duration-500"
          />
        ) : (
          <img
            src="https://via.placeholder.com/1200x300"
            className="w-full rounded-lg"
          />
        )}
      </div>

      {/* PRODUCTS */}
      <div className="m-4 bg-white p-4 rounded-lg">
        <h2 className="text-lg font-bold mb-4">
          {categoryId ? `Category: ${categoryName}` : "All Products"}
        </h2>

        {loading ? (
          <p>Loading...</p>
        ) : (
          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            {products.map((p) => (
              <ProductCard key={p.id} product={p} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default Home;