import LoginForm from "../components/LogInForm";
import { Link } from "react-router-dom";

export default function LoginPage() {
  return (
    <div>
      <LoginForm />
      <p style={{ textAlign: "center", marginTop: 20 }}>
        Don't have an account? <Link to="/signup">Sign up</Link>
      </p>
    </div>
  );
}
