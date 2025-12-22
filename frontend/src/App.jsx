// App.jsx
import React, { useEffect, useState } from "react";
import TaskBoard from "./components/TaskBoard";
import { fetchTasks } from "./services/taskService";

function App() {
  const [tasks, setTasks] = useState([]);

  // load tasks from backend when app starts
  useEffect(() => {
    const loadTasks = async () => {
      const data = await fetchTasks();
      setTasks(data);
    };
    loadTasks();
  }, []);

  return (
    <div className="app">
      <h1>Task Manager</h1>
      <TaskBoard tasks={tasks} />
    </div>
  );
}

export default App;
