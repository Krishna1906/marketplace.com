import { useState } from "react";
import { loginUser } from "../api/auth";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const navigate = useNavigate(); // 👈 ADD THIS

  const [form, setForm] = useState({
    email: "",
    password: "",
  });

  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage("");

    try {
      const res = await loginUser(form);

      // Save token
      const token = res.data.token;
      if (token) {
        localStorage.setItem("token", token);
      }

      setMessage("Login successful ✅");

      // 🔥 REDIRECT TO HOME
      navigate("/");

    } catch (error) {
      setMessage("Login failed ❌");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "100px auto" }}>
      <h2>Login</h2>

      <form onSubmit={handleSubmit}>
        <input name="email" onChange={handleChange} placeholder="Email" />
        <br /><br />
        <input name="password" type="password" onChange={handleChange} placeholder="Password" />
        <br /><br />
        <button type="submit">
          {loading ? "Logging in..." : "Login"}
        </button>
      </form>

      {message && <p>{message}</p>}
    </div>
  );
};

export default Login;