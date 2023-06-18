import { UserContext } from "../services/UserContext";
import { useContext } from "react";
import MyList from "./MyList";
import CreateTask from "./CreateTask";

export function Home() {
  const { username, setUsername } = useContext(UserContext);

  function logout() {
    setUsername(null);
    localStorage.removeItem("token");
  }

  return (
    <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 pt-8">
      <div className="flex justify-between items-center">
        <div className="flex">
          <p className="text-lg font-semibold">Username:</p>
          <p className="ml-2 text-lg">{username}</p>
        </div>
        <button
          onClick={logout}
          className="bg-red-400 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Logout
        </button>
      </div>
      <div className="text-center mt-2">
        <h2 className="text-base font-semibold text-indigo-600 tracking-wide uppercase">
          Welcome to
        </h2>
        <p className="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight lg:text-6xl">
          Golang Todo App
        </p>
        <p className="max-w-xl mt-5 mx-auto text-xl text-gray-500">
          Start managing your tasks!
        </p>
      </div>
      <div className="mt-10">
        <CreateTask />
      </div>
      <div>
        <MyList />
      </div>
    </div>
  );
}

export default Home;
