import ChangeNameForm from "../components/ChangeNameForm";
import ChangeUsernameForm from "../components/ChangeUsernameForm";
import ChangePasswordForm from "../components/ChangePasswordForm";

export default function Settings() {
  return (
    <div style={{ maxWidth: 400, margin: "40px auto" }}>
      <h2>Account Settings</h2>

      <ChangeNameForm />
      <hr />

      <ChangeUsernameForm />
      <hr />

      <ChangePasswordForm />
    </div>
  );
}
