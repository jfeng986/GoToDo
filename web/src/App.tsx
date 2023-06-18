import { UserContextProvider } from "./services/UserContext";
import GoToDoRoutes from "./GoToDoRoutes";

function App() {
  return (
    <UserContextProvider>
      <GoToDoRoutes />
    </UserContextProvider>
  );
}

export default App;
