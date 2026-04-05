import { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";

const steps = [
  "PLACED",
  "SHIPPED",
  "OUT_FOR_DELIVERY",
  "DELIVERED",
];

const TrackOrder = () => {
  const { order_id } = useParams();
  const [order, setOrder] = useState(null);

  const token = localStorage.getItem("token");

  useEffect(() => {
    fetchOrder();
  }, []);

  const fetchOrder = async () => {
    try {
      const res = await axios.get(
        `http://localhost:8080/api/orders/track/${order_id}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setOrder(res.data);
    } catch (err) {
      console.error(err);
    }
  };

  const getStepIndex = (status) => {
    return steps.indexOf(status);
  };

  if (!order) return <p className="p-6">Loading...</p>;

  const currentStep = getStepIndex(order.status);

  return (
    <div className="bg-gray-100 min-h-screen p-6">

      <div className="bg-white p-6 rounded shadow max-w-2xl mx-auto">

        <h2 className="text-xl font-bold mb-4">
          Track Order #{order.order_id}
        </h2>

        {/* 🔹 Timeline */}
        <div className="flex flex-col gap-6">

          {steps.map((step, index) => (
            <div key={step} className="flex items-center gap-4">

              {/* Circle */}
              <div
                className={`w-6 h-6 rounded-full flex items-center justify-center
                  ${index <= currentStep ? "bg-green-500" : "bg-gray-300"}
                `}
              >
                {index <= currentStep && "✔"}
              </div>

              {/* Text */}
              <div>
                <p
                  className={`font-medium ${
                    index <= currentStep ? "text-green-600" : "text-gray-400"
                  }`}
                >
                  {step.replaceAll("_", " ")}
                </p>
              </div>

            </div>
          ))}

        </div>

        {/* 🔹 Status Message */}
        <div className="mt-6 text-center">
          <p className="text-lg font-semibold text-blue-600">
            Current Status: {order.status.replaceAll("_", " ")}
          </p>
        </div>

      </div>

    </div>
  );
};

export default TrackOrder;