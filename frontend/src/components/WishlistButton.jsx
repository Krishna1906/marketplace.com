import { FaHeart, FaRegHeart } from "react-icons/fa";

const WishlistButton = ({ productId, isInWishlist, toggleWishlist }) => {
  const isWishlisted = isInWishlist(productId);

  return (
    <div
      onClick={() => toggleWishlist(productId)}
      style={{
        position: "absolute",
        top: "10px",
        right: "10px",
        cursor: "pointer",
        zIndex: 10,
      }}
    >
      {isWishlisted ? (
        <FaHeart size={22} color="red" />
      ) : (
        <FaRegHeart size={22} color="gray" />
      )}
    </div>
  );
};

export default WishlistButton;