import { BrowserRouter, Routes, Route } from "react-router-dom";

import Layout from "./components/Layout";
import SessionManager from "./components/SessionManager";
import CartProvider from "./context/CartContext";

import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import ProductDetails from "./pages/ProductDetails";
import Cart from "./pages/Cart";
import OrderSummary from "./pages/OrderSummary";
import Payment from "./pages/Payment";
import Orders from "./pages/Orders";
import TrackOrder from "./pages/TrackOrder";
import Wishlist from "./pages/Wishlist";

import Dashboard from "./admin/Dashboard";
import AdminRoute from "./components/AdminRoute";

function App() {
  return (
    <BrowserRouter>
      <SessionManager />
      <CartProvider>
        <Layout>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/product/:id" element={<ProductDetails />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/cart" element={<Cart />} />
            <Route path="/order-summary" element={<OrderSummary />} />
            <Route path="/payment" element={<Payment />} />
            <Route path="/orders" element={<Orders />} />
            <Route path="/track/:order_id" element={<TrackOrder />} />
            <Route path="/wishlist" element={<Wishlist />} />
            <Route
              path="/admin"
              element={
                <AdminRoute>
                  <Dashboard />
                </AdminRoute>
              }
            />
          </Routes>
        </Layout>
      </CartProvider>
    </BrowserRouter>
  );
}

export default App;