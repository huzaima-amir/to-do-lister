import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function UserMenu() {
  const [open, setOpen] = useState(false);
  const navigate = useNavigate();

  function logout() {
    localStorage.removeItem("token");
    navigate("/login");
  }

  return (
    <div style={{ position: "relative" }}>
      <div
        onClick={() => setOpen(!open)}
        style={{
          width: 40,
          height: 40,
          borderRadius: "50%",
          background: "#ddd",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          cursor: "pointer",
          fontWeight: "bold"
        }}
      >
        U
      </div>

      {open && (
        <div
          style={{
            position: "absolute",
            right: 0,
            top: 50,
            background: "white",
            border: "1px solid #ccc",
            borderRadius: 6,
            padding: 10,
            width: 150
          }}
        >
          <p
            style={{ cursor: "pointer", margin: 0 }}
            onClick={() => navigate("/settings")}
          >
            Account Settings
          </p>

          <p
            style={{ cursor: "pointer", marginTop: 10, color: "red" }}
            onClick={logout}
          >
            Logout
          </p>
        </div>
      )}
    </div>
  );
}
