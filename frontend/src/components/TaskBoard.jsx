// TaskBoard.js
import React from "react";
import "./TaskBoard.css";

function TaskBoard() {
  return (
    <div className="task-board">
      <div className="task-column">
        <h2>Upcoming</h2>
        {/* upcoming tasks */}
      </div>

      <div className="task-column">
        <h2>In Progress</h2>
        {/* in progress tasks */}
      </div>

      <div className="task-column">
        <h2>Finished</h2>
        {/* finished tasks */}
      </div>
    </div>
  );
}

export default TaskBoard;
