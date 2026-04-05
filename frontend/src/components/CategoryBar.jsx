import { useEffect, useState } from "react";
import { getCategories } from "../api/category";

const CategoryBar = () => {
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const res = await getCategories();
        setCategories(res.data); // adjust if needed
      } catch (err) {
        console.error("Failed to load categories");
      } finally {
        setLoading(false);
      }
    };

    fetchCategories();
  }, []);

  return (
    <div className="flex justify-around bg-white py-3 shadow overflow-x-auto">
      
      {loading ? (
        <p>Loading...</p>
      ) : (
        categories.map((cat) => (
          <div
            key={cat.id}
            className="text-center text-sm cursor-pointer px-4"
          >
            <div className="font-medium">
              {cat.name}
            </div>
          </div>
        ))
      )}
      
    </div>
  );
};

export default CategoryBar;