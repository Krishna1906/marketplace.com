import { Navigate } from "react-router-dom";

const AdminRoute = ({ children }) => {
  const role = localStorage.getItem("role");

  if (role !== "ADMIN") {
    return <Navigate to="/" />;
  }

  return children;
};

export default AdminRoute;