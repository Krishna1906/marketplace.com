import { useEffect, useState } from "react";
import axios from "axios";

const AddressModal = ({ show, onClose, existingAddress, refresh }) => {
  const token = localStorage.getItem("token");

  const initialState = {
    type: "Home",
    full_name: "",
    phone: "",
    address_line: "",
    city: "",
    state: "",
    pincode: "",
  };

  const [form, setForm] = useState(initialState);

  // ✅ Handle Add vs Edit
  useEffect(() => {
    if (existingAddress) {
      // 🔥 EDIT MODE
      setForm(existingAddress);
    } else {
      // 🔥 ADD MODE (RESET)
      setForm(initialState);
    }
  }, [existingAddress, show]);

  if (!show) return null;

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  // ✅ SAVE (POST / PUT)
  const saveAddress = async () => {
    try {
      if (existingAddress?.id) {
        // 🔥 UPDATE
        await axios.put(
          `http://localhost:8080/api/ordersummary/address/${existingAddress.id}`,
          form,
          {
            headers: { Authorization: `Bearer ${token}` },
          }
        );
      } else {
        // 🔥 CREATE
        await axios.post(
          "http://localhost:8080/api/ordersummary/address",
          form,
          {
            headers: { Authorization: `Bearer ${token}` },
          }
        );
      }

      refresh();
      handleClose();
    } catch (err) {
      console.error("Save Error:", err);
    }
  };

  // ✅ DELETE
  const deleteAddress = async () => {
    if (!existingAddress?.id) return;

    try {
      await axios.delete(
        `http://localhost:8080/api/ordersummary/address/${existingAddress.id}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      refresh();
      handleClose();
    } catch (err) {
      console.error("Delete Error:", err);
    }
  };

  // ✅ CLOSE + RESET
  const handleClose = () => {
    setForm(initialState);
    onClose();
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-40 flex justify-center items-center">
      <div className="bg-white p-6 rounded w-[400px] editaddress space-y-3">

        <h2 className="text-lg font-semibold">
          {existingAddress ? "Edit Address" : "Add Address"}
        </h2>

        {/* TYPE */}
        <select
          name="type"
          value={form.type}
          onChange={handleChange}
          className="border p-2"
        >
          <option>Home</option>
          <option>Work</option>
          <option>Office</option>
        </select><br />

        {/* INPUTS */}
        <input
          name="full_name"
          placeholder="Full Name"
          value={form.full_name}
          onChange={handleChange}
          className="border p-2"
        /><br />

        <input
          name="phone"
          placeholder="Phone"
          value={form.phone}
          onChange={handleChange}
          className="border p-2"
        /><br />

        <input
          name="address_line"
          placeholder="Address"
          value={form.address_line}
          onChange={handleChange}
          className="border p-2"
        /><br />

        <input
          name="city"
          placeholder="City"
          value={form.city}
          onChange={handleChange}
          className="border p-2"
        /><br />

        <input
          name="state"
          placeholder="State"
          value={form.state}
          onChange={handleChange}
          className="border p-2"
        /><br />

        <input
          name="pincode"
          placeholder="Pincode"
          value={form.pincode}
          onChange={handleChange}
          className="border p-2"
        /><br />

        {/* ACTIONS */}
        <div className="flex justify-between mt-3">

          {/* DELETE only in edit */}
          {existingAddress?.id && (
            <button onClick={deleteAddress} className="text-red-500">
              Delete
            </button>
          )}

          <div className="flex gap-2 ml-auto">
            <button onClick={handleClose}>Cancel</button>

            <button
              onClick={saveAddress}
              className="bg-blue-500 text-white px-4 py-1 rounded"
            >
              Save
            </button>
          </div>

        </div>

      </div>
    </div>
  );
};

export default AddressModal;