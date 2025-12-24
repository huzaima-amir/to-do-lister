import { Routes, Route, Navigate } from "react-router-dom";

import LoginPage from "./pages/LogIn";
import SignupPage from "./pages/SignUp";
import Dashboard from "./pages/Dashboard";
import Settings from "./pages/UserSettings";

function App() {
  return (
    <Routes>
      {/* Redirect root to login */}
      <Route path="/" element={<Navigate to="/login" replace />} />

      <Route path="/login" element={<LoginPage />} />
      <Route path="/signup" element={<SignupPage />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/settings" element={<Settings />} />
    </Routes>
  );
}

export default App;

