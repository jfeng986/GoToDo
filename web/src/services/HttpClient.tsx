import axios from "axios";

const serverIP = import.meta.env.VITE_HOST;
const serverPort = import.meta.env.VITE_PORT;

const serverUrl = `http://${serverIP}:${serverPort}`;
//axios.defaults.withCredentials = true;
export const httpClient = axios.create({
  baseURL: serverUrl,
});
