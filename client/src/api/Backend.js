import axios from "axios";

const reset = async () => axios.post("/api/reset");

const play = async (folderId, songId) =>
  axios.post("/api/play", { folderId, songId });

const stop = async (folderId, songId) =>
  axios.post("/api/stop", { folderId, songId });

const readSDContent = async () => {
  return axios
    .get("/api/sd-content")
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
};

export { reset, play, stop, readSDContent };
