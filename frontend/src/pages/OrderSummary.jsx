import { useEffect, useState } from "react";
import axios from "axios";
import AddressModal from "../components/AddressModal";
import { useNavigate } from "react-router-dom";

const OrderSummary = () => {
  const [order, setOrder] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [selectedAddress, setSelectedAddress] = useState(null);
  const navigate = useNavigate();

  // ✅ NEW STATE
  const [selectedAddressId, setSelectedAddressId] = useState(null);

  const token = localStorage.getItem("token");

  useEffect(() => {
    fetchOrderSummary();
  }, []);

  const fetchOrderSummary = async () => {
    try {
      const res = await axios.get(
        "http://localhost:8080/api/ordersummary",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setOrder(res.data);

      // ✅ AUTO SELECT FIRST ADDRESS (optional UX)
      if (res.data.addresses?.length > 0) {
        setSelectedAddressId(res.data.addresses[0].id);
      }

    } catch (err) {
      console.error(err);
    }
  };

  const handleContinue = () => {
  const selected = order.addresses.find(
    (a) => a.id === selectedAddressId
  );

  // ✅ Save selected address
  localStorage.setItem("selectedAddress", JSON.stringify(selected));

  // ✅ Navigate to payment page
  navigate("/payment");
};

  if (!order) return <p className="p-6">Loading...</p>;

  return (
    <div className="bg-gray-100 min-h-screen p-6">

      {/* STEP HEADER */}
      <div className="flex items-center gap-6 mb-6 text-gray-600">
        <div className="flex items-center gap-2">
          <span className="bg-green-500 text-white rounded-full px-2">✓</span>
          Address
        </div>
        <div className="flex items-center gap-2 font-semibold text-blue-600">
          <span className="bg-blue-500 text-white rounded-full px-2">2</span>
          Order Summary
        </div>
        <div className="flex items-center gap-2">
          <span className="bg-gray-300 rounded-full px-2">3</span>
          Payment
        </div>
      </div>

      <div className="grid md:grid-cols-3 gap-6">

        {/* LEFT SECTION */}
        <div className="md:col-span-2 space-y-4">

          {/* 🔥 ADDRESS LIST */}
          <div className="bg-white p-4 rounded shadow">
            <div className="flex justify-between mb-3">
              <h2 className="font-semibold">Select Address</h2>

              <button
                onClick={() => {
                  setSelectedAddress(null);
                  setShowModal(true);
                }}
                className="text-blue-600 text-sm"
              >
                + Add New
              </button>
            </div>

            {order.addresses?.length === 0 && (
              <p className="text-sm text-gray-500">
                No address found. Please add one.
              </p>
            )}

            {order.addresses?.map((addr) => (
              <div
                key={addr.id}
                onClick={() => setSelectedAddressId(addr.id)}
                className={`border p-3 mb-2 rounded flex justify-between items-center cursor-pointer ${
                  selectedAddressId === addr.id
                    ? "border-blue-500 bg-blue-50"
                    : ""
                }`}
              >
                <div className="flex gap-3 items-start">

                  {/* ✅ RADIO BUTTON */}
                  <input
                    type="radio"
                    checked={selectedAddressId === addr.id}
                    onChange={() => setSelectedAddressId(addr.id)}
                  />

                  <div>
                    <p className="font-medium">
                      {addr.full_name}
                      <span className="ml-2 text-xs bg-gray-200 px-2 py-1">
                        {addr.type}
                      </span>
                    </p>

                    <p className="text-sm text-gray-600">
                      {addr.address_line}, {addr.city}, {addr.state} -{" "}
                      {addr.pincode}
                    </p>

                    <p className="text-sm">{addr.phone}</p>
                  </div>
                </div>

                {/* EDIT BUTTON */}
                <button
                  onClick={(e) => {
                    e.stopPropagation(); // 🔥 IMPORTANT
                    setSelectedAddress(addr);
                    setShowModal(true);
                  }}
                  className="text-blue-600 text-sm"
                >
                  Edit
                </button>
              </div>
            ))}
          </div>

          {/* 🔥 ITEMS */}
          {order.items.map((item) => (
            <div
              key={item.product_id}
              className="bg-white p-4 rounded shadow flex gap-4"
            >
              <img
                className="summaryimage w-24 h-24 object-cover"
                src={`http://localhost:8080${item.image}`}
                alt={item.name}
              />

              <div className="flex-1">
                <h3 className="font-medium">{item.name}</h3>

                <p className="text-sm text-gray-500 mt-1">
                  Qty: {item.quantity}
                </p>

                <p className="text-green-600 font-semibold mt-2">
                  ₹{item.price.toLocaleString("en-IN")}
                </p>

                <p className="text-sm text-gray-500">
                  Delivery in 2-3 days
                </p>
              </div>
            </div>
          ))}
        </div>

        {/* RIGHT SECTION */}
        <div className="space-y-4">

          {/* PRICE DETAILS */}
          <div className="bg-white p-4 rounded shadow">
            <h2 className="font-semibold mb-3">Price Details</h2>

            <div className="flex justify-between text-sm mb-2">
              <span>Price ({order.items.length} items)</span>
              <span>₹{order.total_amount.toLocaleString("en-IN")}</span>
            </div>

            <div className="flex justify-between text-sm mb-2">
              <span>Delivery Charges</span>
              <span className="text-green-600">FREE</span>
            </div>

            <hr className="my-3" />

            <div className="flex justify-between font-bold text-lg">
              <span>Total Amount</span>
              <span>₹{order.total_amount.toLocaleString("en-IN")}</span>
            </div>

            <p className="text-green-600 text-sm mt-2">
              You'll save extra on this order!
            </p>
          </div>

          {/* ✅ CONTINUE BUTTON */}
          <button
  onClick={handleContinue}
  disabled={!selectedAddressId}
  className={`orderbutton py-3 font-semibold rounded ${
    selectedAddressId
      ? "bg-yellow-400 hover:bg-yellow-500"
      : "bg-gray-300 cursor-not-allowed"
  }`}
>
  Continue
</button>
        </div>

        {/* 🔥 MODAL */}
        <AddressModal
          show={showModal}
          onClose={() => setShowModal(false)}
          existingAddress={selectedAddress}
          refresh={fetchOrderSummary}
        />
      </div>
    </div>
  );
};

export default OrderSummary;