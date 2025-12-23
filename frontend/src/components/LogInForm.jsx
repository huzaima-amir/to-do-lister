import { useState } from "react";
import { login } from "../api/auth";

export default function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();
    setMessage("");

    try {
      const data = await login(username, password);
      localStorage.setItem("token", data.token);
      setMessage("Login successful!");
    } catch (err) {
      setMessage(err.message);
    }
  }

  return (
    <div style={{ maxWidth: 300, margin: "50px auto" }}>
      <h2>Login</h2>

      <form onSubmit={handleSubmit}>
        <div>
          <label>Username</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>

        <div style={{ marginTop: 10 }}>
          <label>Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        <button style={{ marginTop: 20 }} type="submit">
          Login
        </button>
      </form>

      {message && (
        <p style={{ marginTop: 20, color: message.includes("successful") ? "green" : "red" }}>
          {message}
        </p>
      )}
    </div>
  );
}
