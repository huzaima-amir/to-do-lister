import SignupForm from "../components/SignUpForm";
import { Link } from "react-router-dom";

export default function SignupPage() {
  return (
    <div>
      <SignupForm />
      <p style={{ textAlign: "center", marginTop: 20 }}>
        Already have an account? <Link to="/login">Log in</Link>
      </p>
    </div>
  );
}
