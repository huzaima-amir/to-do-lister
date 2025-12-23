import { useState } from "react";
import { signup } from "../api/auth";

export default function SignupForm() {
  const [name, setName] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();
    setMessage("");

    try {
      const data = await signup(name, username, password);
      setMessage("Signup successful! Your user ID is " + data.id);
    } catch (err) {
      setMessage(err.message);
    }
  }

  return (
    <div style={{ maxWidth: 300, margin: "50px auto" }}>
      <h2>Sign Up</h2>

      <form onSubmit={handleSubmit}>
        <div>
          <label>Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>

        <div style={{ marginTop: 10 }}>
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
          Sign Up
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
