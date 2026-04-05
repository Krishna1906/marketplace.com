import Navbar from "./Navbar";

const Layout = ({ children }) => {
  return (
    <>
      <Navbar />

      {/* 🔥 ADD TOP PADDING */}
      <div className="home pt-24">
        {children}
      </div>
    </>
  );
};

export default Layout;