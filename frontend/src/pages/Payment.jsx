import { useEffect, useState } from "react";
import axios from "axios";
import api from "../utils/axios";
import { useNavigate } from "react-router-dom";

const Payment = () => {
const navigate = useNavigate();
  const [methods, setMethods] = useState([]);
  const [selectedMethod, setSelectedMethod] = useState("CASH_ON_DELIVERY");
  const [totalAmount, setTotalAmount] = useState(0);

  const [card, setCard] = useState({
    card_number: "",
    expiry: "",
    cvv: "",
  });

  const token = localStorage.getItem("token");

  useEffect(() => {
  fetchMethods();
  fetchSummary();
}, []);

const fetchSummary = async () => {
  try {
    const res = await api.get("/ordersummary");
    setTotalAmount(res.data.total_amount);
  } catch (err) {
    console.error(err);
  }
};

  const fetchMethods = async () => {
  const res = await api.get("/payment");
  setMethods(res.data.methods);
};

  const handleCardChange = (e) => {
    setCard({ ...card, [e.target.name]: e.target.value });
  };

  const isCardValid =
    card.card_number.length >= 12 &&
    card.expiry.length >= 4 &&
    card.cvv.length >= 3;

  // ✅ SAVE CARD
  const saveCard = async () => {
    await axios.post(
      "http://localhost:8080/api/payment/card",
      card,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );
  };

  // ✅ PLACE ORDER
  const placeOrder = async () => {
  try {
    let paymentType = selectedMethod;

    if (selectedMethod === "CASH_ON_DELIVERY") {
      paymentType = "COD";
    }

    if (selectedMethod === "CARD") {
      await saveCard();
    }

    await axios.post(
      "http://localhost:8080/api/orders/place",
      {
        payment_type: paymentType,
      },
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    alert("Order placed successfully!");

    // ✅ HARD REFRESH (simple way)
    window.location.href = "/orders";

  } catch (err) {
    console.error(err);
    alert("Something went wrong");
  }
};

  return (
    <div className="flex p-6 gap-6">

      {/* LEFT - METHODS */}
      <div className="w-1/3 bg-white shadow rounded">

        {methods.map((method) => (
          <div
            key={method}
            onClick={() => setSelectedMethod(method)}
            className={`p-4 border-b cursor-pointer ${
              selectedMethod === method ? "bg-gray-100" : ""
            }`}
          >
            {method === "CARD" && "Credit / Debit / ATM Card"}
            {method === "CASH_ON_DELIVERY" && "Cash on Delivery"}
            {method === "UPI" && "UPI"}
          </div>
        ))}

      </div>

      {/* RIGHT - CONTENT */}
      <div className="w-2/3 bg-white p-6 shadow rounded">

        {/* CARD */}
        {selectedMethod === "CARD" && (
          <div className="space-y-4">

            <h3 className="font-semibold">Card Payment</h3>

            <input
              name="card_number"
              placeholder="Card Number"
              value={card.card_number}
              onChange={handleCardChange}
              className="border p-2"
            />

            <div className="flex gap-3">
              <input
                name="expiry"
                placeholder="MM/YY"
                value={card.expiry}
                onChange={handleCardChange}
                className="w-1/2 border p-2"
              />

              <input
                name="cvv"
                placeholder="CVV"
                value={card.cvv}
                onChange={handleCardChange}
                className="w-1/2 border p-2"
              />
            </div>

            {isCardValid && (
              <button
                onClick={placeOrder}
                className="bg-yellow-400 py-2 rounded"
              >
                 Pay ₹{totalAmount.toLocaleString("en-IN")} & Place Order
              </button>
            )}

          </div>
        )}

        {/* COD */}
        {selectedMethod === "CASH_ON_DELIVERY" && (
          <div className="space-y-4">

            <h3 className="font-semibold">Cash on Delivery</h3>

            <button
              onClick={placeOrder}
              className="bg-yellow-400 py-2 rounded"
            >
              Pay ₹{totalAmount.toLocaleString("en-IN")} & Place Order
            </button>

          </div>
        )}

        {/* UPI */}
        {selectedMethod === "UPI" && (
          <div className="space-y-4">

            <h3 className="font-semibold">UPI</h3>

            <input
              placeholder="Enter UPI ID"
              className="border p-2"
            />

            <button
              onClick={placeOrder}
              className="bg-yellow-400 py-2 rounded"
            >
              Pay ₹{totalAmount.toLocaleString("en-IN")} & Place Order
            </button>

          </div>
        )}

      </div>
    </div>
  );
};

export default Payment;