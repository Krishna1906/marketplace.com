import { useState, useEffect } from "react"; // <-- add useEffect here
import { useNavigate } from "react-router-dom";
import api from "../utils/axios";
import { Link } from "react-router-dom";

const Login = () => {
  const navigate = useNavigate();
  const [form, setForm] = useState({ email: "", password: "" });
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const handleChange = (e) =>
    setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage("");

    try {
      const res = await api.post("/auth/login", form);
      const token = res.data.token;
      if (token) localStorage.setItem("token", token);
      setMessage("Login successful ✅");
      navigate("/");
    } catch (err) {
      setMessage("Login failed ❌");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "100px auto" }}>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <input name="email" onChange={handleChange} placeholder="Email" autoComplete="new-email" />
        <br />
        <br />
        <input
          name="password"
          type="password"
          onChange={handleChange}
          placeholder="Password"
          autoComplete="new-password"
        />
        <br />
        <br />
        <button type="submit">{loading ? "Logging in..." : "Login"}</button>
      </form>
      {message && <p>{message}</p>}

      <p style={{ marginTop: "10px" }}>
  Already a user?{" "}
  <Link to="/register" style={{ color: "blue", textDecoration: "underline" }}>
    Register here
  </Link>
</p>
      
    </div>
  );
};

export default Login;