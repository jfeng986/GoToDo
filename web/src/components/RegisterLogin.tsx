import { useContext, useState } from "react";
import { UserContext } from "../services/UserContext.tsx";
import { httpClient } from "../services/HttpClient.tsx";

export default function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoginOrRegister, setIsLoginOrRegister] = useState("login");
  const { setUsername: setUsernameContext, setTasks } = useContext(UserContext);

  async function handleSubmit(event: any) {
    event.preventDefault();
    const url = isLoginOrRegister === "register" ? "/register" : "/login";
    const response = await httpClient.post(`/api/v1/user${url}`, {
      username,
      password,
    });

    if (response.data.code === 200) {
      localStorage.setItem("token", response.data.token);
      setUsernameContext(username);

      const taskResponse = await httpClient.get("/api/v1/user/profile", {
        headers: {
          Authorization: `Bearer ${response.data.token}`,
        },
      });

      if (taskResponse.status === 200) {
        const tasksData = Array.isArray(taskResponse.data.profile.tasks)
          ? taskResponse.data.profile.tasks.map((task: any) => ({
              id: task.id,
              title: task.title,
              content: task.content,
              status: task.status,
            }))
          : [];

        setTasks(tasksData);
      }
    } else {
      alert(response.data.message);
    }
  }

  return (
    <div className="bg-blue-50 h-screen flex items-center justify-center">
      <form
        className="w-full max-w-sm bg-white p-6 rounded-lg shadow-md"
        onSubmit={handleSubmit}
      >
        <h1 className="text-3xl text-center text-blue-500 mb-5 font-bold">
          GoToDo
        </h1>
        <input
          value={username}
          onChange={(event) => setUsername(event.target.value)}
          type="text"
          placeholder="Username"
          className="w-full px-3 py-2 mb-3 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
        <input
          value={password}
          onChange={(event) => setPassword(event.target.value)}
          type="password"
          placeholder="Password"
          className="w-full px-3 py-2 mb-3 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
        <button className="w-full px-4 py-2 mb-4 text-sm font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline">
          {isLoginOrRegister === "register" ? "Register" : "Login"}
        </button>
        <div className="text-center mt-2 text-sm text-gray-600">
          {isLoginOrRegister === "login" && (
            <div>
              Don't have an account?{" "}
              <button
                onClick={() => setIsLoginOrRegister("register")}
                className="text-blue-500 underline hover:text-blue-700"
              >
                Register here
              </button>
            </div>
          )}
          {isLoginOrRegister === "register" && (
            <div>
              Already a member?{" "}
              <button
                onClick={() => setIsLoginOrRegister("login")}
                className="text-blue-500 underline hover:text-blue-700"
              >
                Login here
              </button>
            </div>
          )}
        </div>
      </form>
    </div>
  );
}
