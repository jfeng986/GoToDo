import { createContext, useEffect, useState } from "react";
import { httpClient } from "./HttpClient";
import { Task, UserContextType } from "../Types";

export const UserContext = createContext<UserContextType>({
  username: null,
  setUsername: () => {},
  tasks: [],
  setTasks: () => {},
  loading: true,
});

export function UserContextProvider({ children }: any) {
  const [loading, setLoading] = useState(true);
  const [username, setUsername] = useState<string | null>(null);
  const [tasks, setTasks] = useState<Task[] | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setLoading(false);
      return;
    }
    const getPofile = async () => {
      const response = await httpClient.get("/api/v1/user/profile", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      if (response.data.code !== 200) {
        setLoading(false);
        localStorage.removeItem("token");
        return;
      }
      console.log("response.data.profile: ", response.data.profile);
      setUsername(response.data.profile.user.username);

      const tasksData = Array.isArray(response.data.profile.tasks) // Check if tasks is an array
        ? response.data.profile.tasks.map((task: any) => ({
            id: task.id,
            title: task.title,
            content: task.content,
            status: task.status,
          }))
        : [];

      setTasks(tasksData);
      setLoading(false);
      console.log("response: ", response);
      console.log("tasksData: ", tasksData);
      console.log("username: ", username);
    };
    getPofile();
  }, []);

  return (
    <UserContext.Provider
      value={{ username, setUsername, tasks, setTasks, loading }}
    >
      {children}
    </UserContext.Provider>
  );
}
