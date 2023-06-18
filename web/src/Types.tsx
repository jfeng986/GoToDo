export type Task = {
  id: string;
  title: string;
  content: string;
  status: string;
};

export type UserContextType = {
  username: string | null;
  setUsername: (username: string | null) => void;
  tasks: Task[] | null;
  setTasks: (tasks: Task[] | null) => void;
  loading: boolean;
};
