 // Login and signup API helpers

export async function signup(name, username, password) {
  const response = await fetch("http://localhost:8080/users/signup", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name, username, password }),
  });

  if (!response.ok) {
    const text = await response.text();
    throw new Error(text || "Signup failed");
  }

  return response.json(); // { id: 123 }
}


export async function login(username, password) {
  const response = await fetch("http://localhost:8080/users/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (!response.ok) {
    const text = await response.text();
    throw new Error(text || "Login failed");
  }

  return response.json(); // { token: "..." }
}
