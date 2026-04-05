import { useEffect, useState } from "react";
import axios from "axios";

const Payment = () => {
  const [address, setAddress] = useState(null);
  const [method, setMethod] = useState("COD");

  const [card, setCard] = useState({
    card_number: "",
    expiry: "",
    cvv: "",
  });

  const token = localStorage.getItem("token");

  useEffect(() => {
    const saved = localStorage.getItem("selectedAddress");
    if (saved) {
      setAddress(JSON.parse(saved));
    }
  }, []);

  const handleCardChange = (e) => {
    setCard({ ...card, [e.target.name]: e.target.value });
  };

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
      if (method === "CARD") {
        await saveCard();
      }

      await axios.post(
        "http://localhost:8080/api/orders/place",
        {
          payment_method: method,
          address_id: address.id, // 🔥 IMPORTANT
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      alert("Order placed successfully!");

    } catch (err) {
      console.error(err);
      alert("Payment failed");
    }
  };

  if (!address) return <p className="p-6">No address selected</p>;

  return (
    <div className="p-6 max-w-xl mx-auto space-y-4">

      <h2 className="text-xl font-semibold">Payment</h2>

      {/* ADDRESS */}
      <div className="bg-white p-4 rounded shadow">
        <p className="font-medium">
          {address.full_name} ({address.type})
        </p>
        <p className="text-sm text-gray-600">
          {address.address_line}, {address.city}, {address.state} - {address.pincode}
        </p>
        <p>{address.phone}</p>
      </div>

      {/* PAYMENT METHODS */}
      <div className="bg-white p-4 rounded shadow space-y-3">

        {/* COD */}
        <label className="flex items-center gap-2">
          <input
            type="radio"
            checked={method === "COD"}
            onChange={() => setMethod("COD")}
          />
          Cash on Delivery
        </label>

        {/* CARD */}
        <label className="flex items-center gap-2">
          <input
            type="radio"
            checked={method === "CARD"}
            onChange={() => setMethod("CARD")}
          />
          Pay with Card
        </label>

        {/* CARD FORM */}
        {method === "CARD" && (
          <div className="space-y-2 mt-3">

            <input
              name="card_number"
              placeholder="Card Number"
              value={card.card_number}
              onChange={handleCardChange}
              className="w-full border p-2"
            />

            <input
              name="expiry"
              placeholder="MM/YY"
              value={card.expiry}
              onChange={handleCardChange}
              className="w-full border p-2"
            />

            <input
              name="cvv"
              placeholder="CVV"
              value={card.cvv}
              onChange={handleCardChange}
              className="w-full border p-2"
            />

          </div>
        )}
      </div>

      {/* PLACE ORDER */}
      <button
        onClick={placeOrder}
        className="w-full bg-green-500 text-white py-3 rounded"
      >
        Place Order
      </button>

    </div>
  );
};

export default Payment;