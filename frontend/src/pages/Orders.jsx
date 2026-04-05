import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Orders = () => {
  const [orders, setOrders] = useState([]);
  const navigate = useNavigate();

  const token = localStorage.getItem("token");

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const res = await axios.get(
        "http://localhost:8080/api/orders",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setOrders(res.data || []);
    } catch (err) {
      console.error("Error fetching orders:", err);
    }
  };

  // ✅ Format Date (IST Fix)
  const formatDate = (date) => {
    if (!date) return "N/A";

    try {
      return new Date(
        date.replace(" ", "T").split(".")[0] + "+05:30"
      ).toLocaleString("en-IN");
    } catch {
      return "Invalid Date";
    }
  };

  // ✅ Cancel Order
  const cancelOrder = async (orderId) => {
    try {
      await axios.put(
        `http://localhost:8080/api/orders/cancel/${orderId}`,
        {},
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      alert("Order cancelled");
      fetchOrders();
    } catch (err) {
      alert("Cannot cancel this order");
    }
  };

  // ✅ STATUS STEP CHECK
  const isCompleted = (orderStatus, step) => {
    const steps = ["PLACED", "SHIPPED", "OUT_FOR_DELIVERY", "DELIVERED"];
    return steps.indexOf(orderStatus) >= steps.indexOf(step);
  };

  return (
    <div className="bg-gray-100 min-h-screen p-6">

      <h1 className="text-2xl font-bold mb-6">My Orders</h1>

      {/* ✅ EMPTY STATE */}
      {orders.length === 0 ? (
        <div className="flex flex-col items-center justify-center mt-20">
          <p className="text-5xl">📦</p>
          <p className="text-xl font-semibold text-gray-600 mt-4">
            No Orders Found
          </p>
          <button
            onClick={() => navigate("/")}
            className="mt-6 bg-blue-600 text-white px-5 py-2 rounded"
          >
            Start Shopping
          </button>
        </div>
      ) : (
        orders.map((order, index) => (
          <div
            key={order.order_id}
            className="bg-white p-4 rounded shadow mb-6"
          >

            {/* 🔹 HEADER */}
            <div className="flex justify-between items-center mb-3">

              <div>
                <p className="font-semibold">
                  Order #{orders.length - index}
                </p>

                <p className="text-sm text-gray-500">
                  {formatDate(order.created_at)}
                </p>
              </div>

              <div className="text-right">

                <p className="text-green-600 font-bold text-lg">
                  ₹{order.total_amount.toLocaleString("en-IN")}
                </p>

                <p className="text-sm text-gray-500">
                  {order.payment_method}
                </p>

                {/* STATUS BADGE */}
                <span
                  className={`text-xs px-2 py-1 rounded ${
                    order.status === "DELIVERED"
                      ? "bg-green-100 text-green-700"
                      : order.status === "CANCELLED"
                      ? "bg-red-100 text-red-700"
                      : "bg-yellow-100 text-yellow-700"
                  }`}
                >
                  {order.status.replaceAll("_", " ")}
                </span>

              </div>
            </div>

            {/* ✅ TIMELINE (IMPROVED) */}
            {order.status !== "CANCELLED" && (
              <div className="flex gap-6 text-xs mt-2 mb-3">

                <span className={isCompleted(order.status, "PLACED") ? "text-green-600" : ""}>
                  ● Placed
                </span>

                <span className={isCompleted(order.status, "SHIPPED") ? "text-green-600" : ""}>
                  ● Shipped
                </span>

                <span className={isCompleted(order.status, "OUT_FOR_DELIVERY") ? "text-green-600" : ""}>
                  ● Out for Delivery
                </span>

                <span className={isCompleted(order.status, "DELIVERED") ? "text-green-600" : ""}>
                  ● Delivered
                </span>

              </div>
            )}

            {order.status === "CANCELLED" && (
              <p className="text-red-600 text-sm mb-2">
                Order Cancelled ❌
              </p>
            )}

            {/* 🔹 ITEMS */}
            {order.items?.map((item) => (
              <div
                key={item.product_id}
                className="flex gap-4 border-t pt-3 mt-3"
              >

                <img
                  src={
                    item.image
                      ? "http://localhost:8080" + item.image
                      : "https://via.placeholder.com/100"
                  }
                  alt={item.name}
                  className="w-20 h-20 object-cover rounded"
                />

                <div className="flex-1">
                  <p
                    className="font-medium cursor-pointer"
                    onClick={() =>
                      navigate(`/product/${item.product_id}`)
                    }
                  >
                    {item.name}
                  </p>

                  <p className="text-sm text-gray-500">
                    Qty: {item.quantity}
                  </p>

                  <p className="text-green-600 font-semibold">
                    ₹{item.price.toLocaleString("en-IN")}
                  </p>
                </div>

              </div>
            ))}

            {/* 🔹 ACTIONS */}
            <div className="flex gap-4 mt-4">

              {/* 🚚 TRACK BUTTON */}
              <button
                onClick={() => navigate(`/track/${order.order_id}`)}
                className="text-blue-600 text-sm"
              >
                Track Order
              </button>

              {/* ❌ CANCEL BUTTON */}
              {order.status === "PLACED" && (
                <button
                  onClick={() => cancelOrder(order.order_id)}
                  className="text-red-500 text-sm"
                >
                  Cancel Order
                </button>
              )}

            </div>

          </div>
        ))
      )}
    </div>
  );
};

export default Orders;