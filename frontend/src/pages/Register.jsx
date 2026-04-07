import { useState, useEffect } from "react"; // <-- add useEffect here
import { registerUser } from "../api/auth";
import { FaEye, FaEyeSlash } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";


const Register = () => {
  const [form, setForm] = useState({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("authToken"); // or wherever you store it
    if (token) {
      // User already logged in → redirect to Login page or Dashboard
      navigate("/login");  // or "/dashboard"
    }
  }, [navigate]);

  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");
  const [strength, setStrength] = useState("");

  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  // 🔹 Password strength checker
  const checkPasswordStrength = (password) => {
    if (password.length < 6) return "Weak";
    if (password.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/)) return "Strong";
    return "Medium";
  };

  // 🔹 Handle input change
  const handleChange = (e) => {
    const { name, value } = e.target;

    setForm({
      ...form,
      [name]: value,
    });

    if (name === "password") {
      setStrength(checkPasswordStrength(value));
    }
  };

  // 🔹 Submit
  const handleSubmit = async (e) => {
    e.preventDefault();

    if (form.password !== form.confirmPassword) {
      setMessage("Passwords do not match ❌");
      return;
    }

    setLoading(true);
    setMessage("");

    try {
      const res = await registerUser({
        name: form.name,
        email: form.email,
        password: form.password,
      });

      setMessage("Registration successful ✅");
      console.log(res.data);

      setForm({
        name: "",
        email: "",
        password: "",
        confirmPassword: "",
      });
      setStrength("");

    } catch (error) {
      setMessage(
        error.response?.data?.message || "Registration failed ❌"
      );
    } finally {
      setLoading(false);
    }
  };

  // 🔹 Strength color
  const getStrengthColor = () => {
    if (strength === "Strong") return "green";
    if (strength === "Medium") return "orange";
    return "red";
  };

  return (
    <div style={{ maxWidth: "400px", margin: "50px auto" }}>
      <h2>Register</h2>

      <form onSubmit={handleSubmit} autoComplete="off">
        {/* Name */}
        <input
  type="text"
  name="name"
  placeholder="Name"
  value={form.name}
  onChange={handleChange}
  autoComplete="off"
/>
        <br /><br />

        {/* Email */}
        <input
  type="email"
  name="email"
  placeholder="Email"
  value={form.email}
  onChange={handleChange}
  autoComplete="new-email"
/>
        <br /><br />

        {/* Password */}
        <div style={{ position: "relative" }}>
          <input
  type={showPassword ? "text" : "password"}
  name="password"
  placeholder="Password"
  value={form.password}
  onChange={handleChange}
  autoComplete="new-password"
/>

          <span
            onClick={() => setShowPassword(!showPassword)}
            style={{
              position: "absolute",
              right: "10px",
              top: "50%",
              transform: "translateY(-50%)",
              cursor: "pointer"
            }}
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </span>
        </div>

        {/* Strength */}
        {form.password && (
          <p style={{ color: getStrengthColor() }}>
            Strength: {strength}
          </p>
        )}

        <br />

        {/* Confirm Password */}
        <div style={{ position: "relative" }}>
          <input
  type={showConfirmPassword ? "text" : "password"}
  name="confirmPassword"
  placeholder="Confirm Password"
  value={form.confirmPassword}
  onChange={handleChange}
  autoComplete="new-password"
/>

          <span
            onClick={() =>
              setShowConfirmPassword(!showConfirmPassword)
            }
            style={{
              position: "absolute",
              right: "10px",
              top: "50%",
              transform: "translateY(-50%)",
              cursor: "pointer"
            }}
          >
            {showConfirmPassword ? <FaEyeSlash /> : <FaEye />}
          </span>
        </div>

        {/* Match indicator */}
        {form.confirmPassword && (
          <p style={{
            color:
              form.password === form.confirmPassword
                ? "green"
                : "red"
          }}>
            {form.password === form.confirmPassword
              ? "Passwords match ✅"
              : "Passwords do not match ❌"}
          </p>
        )}

        <br />

        {/* Submit */}
        <button
          type="submit"
          disabled={
            loading ||
            form.password !== form.confirmPassword
          }
        >
          {loading ? "Registering..." : "Register"}
        </button>
      </form>

      {/* Message */}
      {message && <p>{message}</p>}

      <p style={{ marginTop: "10px" }}>
  Already a user?{" "}
  <Link to="/login" style={{ color: "blue", textDecoration: "underline" }}>
    Login here
  </Link>
</p>
    </div>
  );
};

export default Register;