import { useEffect, useRef } from "react";
import { useNavigate, useLocation } from "react-router-dom";

const SessionManager = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const timeoutRef = useRef(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return; // do nothing if not logged in
    if (location.pathname === "/login" || location.pathname === "/register") return;

    const logout = () => {
      localStorage.removeItem("token");
      window.dispatchEvent(new Event("logout"));
      alert("Session expired. Please login again.");
      navigate("/login");
    };

    const resetTimer = () => {
      if (timeoutRef.current) clearTimeout(timeoutRef.current);
      timeoutRef.current = setTimeout(logout, 5 * 60 * 1000); // 5 min
    };

    const events = ["mousemove", "keydown", "click", "scroll"];
    events.forEach((event) => window.addEventListener(event, resetTimer));

    resetTimer(); // start timer

    // sync logout across tabs
    const handleStorage = (e) => {
      if (e.key === "token" && !e.newValue) navigate("/login");
    };
    window.addEventListener("storage", handleStorage);

    return () => {
      if (timeoutRef.current) clearTimeout(timeoutRef.current);
      events.forEach((event) => window.removeEventListener(event, resetTimer));
      window.removeEventListener("storage", handleStorage);
    };
  }, [navigate, location.pathname]);

  return null;
};

export default SessionManager;