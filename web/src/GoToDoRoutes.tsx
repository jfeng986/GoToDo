import { useContext } from "react";
import { UserContext } from "./services/UserContext.tsx";
import RegisterLogin from "./components/RegisterLogin";
import Home from "./components/Home.tsx";

export default function GoToDoRoutes() {
  const { username, loading } = useContext(UserContext);

  if (loading) {
    return <div>Loading...</div>;
  }
  console.log(username);
  if (username) {
    return <Home />;
  }
  return <RegisterLogin />;
}
