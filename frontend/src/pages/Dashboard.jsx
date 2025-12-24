import UserMenu from "../components/UserMenu";

export default function Dashboard() {
  return (
    <div>
      <div style={{
        display: "flex",
        justifyContent: "flex-end",
        padding: 20
      }}>
        <UserMenu />
      </div>

      <h1 style={{ textAlign: "center" }}>Your Tasks</h1>

      {/* kanban task boards will go here */}
    </div>
  );
}
