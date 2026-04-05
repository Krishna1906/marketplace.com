import { useEffect, useState } from "react";
import { getProductImages } from "../api/image";
import { useNavigate } from "react-router-dom";

const ProductCard = ({ product }) => {
  const [image, setImage] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    let isMounted = true; // prevents state update after unmount

    const fetchImage = async () => {
      try {
        console.log("Fetching images for product:", product.id);

        const res = await getProductImages(product.id);

        // Ensure it's always an array
        const images = res?.data ?? [];

        console.log("API RESPONSE:", images);

        if (isMounted) {
          if (Array.isArray(images) && images.length > 0) {
            setImage("http://localhost:8080" + images[0].image_url);
          } else {
            setImage(null); // fallback
            console.log("No images found for product:", product.id);
          }
        }
      } catch (err) {
        console.error("FULL ERROR:", err?.response || err.message);

        if (isMounted) {
          setImage(null); // fallback on error
        }
      }
    };

    fetchImage();

    return () => {
      isMounted = false;
    };
  }, [product.id]);

  return (
    <div
      onClick={() => navigate(`/product/${product.id}`)}
      className="bg-white p-3 rounded-lg shadow hover:shadow-lg cursor-pointer"
    >
      <div className="product_image">
      <img
        src={image || "https://via.placeholder.com/200"}
        alt={product.name}
        className="h-40 w-full object-cover rounded"
      />
</div>
      <h3 className="productname mt-2 font-semibold text-sm">
        {product.name}
      </h3>

      <p className="text-green-600 font-bold">
  ₹{product.price.toLocaleString("en-IN")}
</p>
    </div>
  );
};

export default ProductCard;