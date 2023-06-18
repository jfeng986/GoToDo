import { useState, useContext } from "react";
import { httpClient } from "../services/HttpClient";
import { UserContext } from "../services/UserContext";

const Status = {
  TODO: "Todo",
  IN_PROGRESS: "In Progress",
  DONE: "Done",
};

export default function CreateTask() {
  const [taskName, setTaskName] = useState("");
  const [content, setContent] = useState("");
  const [status, setStatus] = useState(Status.TODO);
  const [showInputs, setShowInputs] = useState(false); // State variable to control the visibility of inputs
  const { tasks, setTasks } = useContext(UserContext);

  const handleCreateTask = async () => {
    const token = localStorage.getItem("token");
    const taskData = {
      title: taskName,
      content: content,
      status: status,
    };
    console.log("Sending task data:", taskData);
    const response = await httpClient.post(
      "/api/v1/create_task",
      {
        title: taskName,
        content: content,
        status: status,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log("response: ", response);
    if (response.data.code === 200) {
      const newTask = {
        id: response.data.task.id,
        title: taskName,
        content: content,
        status: status,
      };

      if (tasks !== null) {
        setTasks([...tasks, newTask]);
      } else {
        setTasks([newTask]);
      }
    }

    setTaskName("");
    setContent("");
    setStatus(Status.TODO);
    setShowInputs(false);
  };

  function handleCancel() {
    setTaskName("");
    setContent("");
    setStatus(Status.TODO);
    setShowInputs(false);
  }

  function showInputFields() {
    setShowInputs(true);
  }

  return (
    <div>
      <div className="flex justify-center">
        <div className="column">
          {!showInputs && (
            <button
              className="bg-blue-400 hover:bg-blue-700 text-white w-full rounded p-2"
              onClick={showInputFields}
            >
              Create Task
            </button>
          )}
          {showInputs && (
            <>
              <input
                value={taskName}
                onChange={(event) => setTaskName(event.target.value)}
                type="text"
                placeholder="Task Name"
                className="block w-full rounded-sm p-2 mb-2"
              />
              <input
                value={content}
                onChange={(event) => setContent(event.target.value)}
                type="text"
                placeholder="Content"
                className="block w-full rounded-sm p-2 mb-2"
              />
              <select
                value={status}
                onChange={(event) => setStatus(event.target.value)}
                className="block w-full rounded-sm p-2 mb-2"
              >
                <option value={Status.TODO}>{Status.TODO}</option>
                <option value={Status.IN_PROGRESS}>{Status.IN_PROGRESS}</option>
                <option value={Status.DONE}>{Status.DONE}</option>
              </select>
              <div className="flex justify-center">
                <div className="px-2 w-full">
                  <button
                    className="bg-blue-400 hover:bg-blue-700 text-white w-full rounded p-2"
                    onClick={handleCreateTask}
                  >
                    Create
                  </button>
                </div>
                <div className="px-2 w-full">
                  <button
                    className="bg-blue-400 hover:bg-blue-700 text-white w-full rounded p-2"
                    onClick={handleCancel}
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
