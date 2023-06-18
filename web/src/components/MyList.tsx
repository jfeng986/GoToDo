import { UserContext } from "../services/UserContext";
import { useContext, useState } from "react";
import { Task } from "../Types";
import { httpClient } from "../services/HttpClient";

export default function MyList() {
  const { tasks, setTasks } = useContext(UserContext);
  const [selectedTask, setSelectedTask] = useState<string | null>(null);

  const todoTasks = tasks?.filter((task) => task.status === "Todo") || [];
  const inProgressTasks =
    tasks?.filter((task) => task.status === "In Progress") || [];
  const doneTasks = tasks?.filter((task) => task.status === "Done") || [];

  const handleTaskClick = (task: Task) => {
    setSelectedTask(task.id === selectedTask ? null : task.id);
  };

  const markAs = async (task: Task, status: string) => {
    try {
      const token = localStorage.getItem("token");
      await httpClient.post(
        `/api/v1/update_task/${task.id}`,
        { status, title: task.title, content: task.content, id: task.id },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      if (!tasks) return;
      setTasks(tasks.map((t) => (t.id === task.id ? { ...t, status } : t)));
    } catch (err) {
      console.error(err);
    }
  };

  const deleteTask = async (task: Task) => {
    try {
      const token = localStorage.getItem("token");
      console.log("token", token);
      const response = await httpClient.post(
        `/api/v1/delete_task/${task.id}`,
        {},
        { headers: { Authorization: `Bearer ${token}` } }
      );
      console.log(response);
      if (!tasks) return;
      setTasks(tasks.filter((t) => t.id !== task.id));
    } catch (err) {
      console.error(err);
    }
  };

  const renderTask = (task: Task, availableStatuses: string[]) => (
    <div
      key={task.id}
      onClick={() => handleTaskClick(task)}
      className="border-2 border-gray-200 p-2 my-2 rounded-md shadow-sm cursor-pointer"
    >
      <div className="flex justify-between">
        <p className="text-base font-semibold">{task.title}</p>
        {task.id === selectedTask ? (
          <button
            className="bg-gray-400 text-white rounded px-2 focus:outline-none"
            onClick={(e) => {
              e.stopPropagation();
              deleteTask(task);
            }}
          >
            Delete
          </button>
        ) : null}
      </div>
      {task.id === selectedTask && (
        <div className="mt-2 text-gray-600">
          <p className="italic pb-2">Description: {task.content}</p>
          <div className="flex justify-end items-center">
            <p className="pr-1">Mark as:</p>
            {availableStatuses.map((status, index) => (
              <button
                key={index}
                onClick={(e) => {
                  e.stopPropagation();
                  markAs(task, status);
                }}
                className="bg-gray-200 px-2 py-1 rounded-md mr-2"
              >
                {status}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );

  return (
    <div className="flex flex-col items-center">
      <div className="grid grid-cols-3 gap-4 w-full p-5">
        <div className="border-2 border-gray-200 p-4 rounded-md shadow-sm overflow-y-auto h-96">
          <h2 className="text-xl font-semibold mb-2">Todo</h2>
          {todoTasks.map((task) => renderTask(task, ["In Progress", "Done"]))}
        </div>
        <div className="border-2 border-gray-200 p-4 rounded-md shadow-sm overflow-y-auto h-96">
          <h2 className="text-xl font-semibold mb-2">In Progress</h2>
          {inProgressTasks.map((task) => renderTask(task, ["Todo", "Done"]))}
        </div>
        <div className="border-2 border-gray-200 p-4 rounded-md shadow-sm overflow-y-auto h-96">
          <h2 className="text-xl font-semibold mb-2">Done</h2>
          {doneTasks.map((task) => renderTask(task, ["Todo", "In Progress"]))}
        </div>
      </div>
    </div>
  );
}
