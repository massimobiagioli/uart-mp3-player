import { useEffect, useState } from "react";
import Container from "@material-ui/core/Container";
import { readSDContent } from "../api/Backend";
import Folder from "./Folder";

const SDContent = () => {
  const [data, setData] = useState();

  useEffect(() => {
    readSDContent()
      .then((data) => setData(data))
      .catch((err) => console.error(err.message));
  }, []);

  if (!data) {
    return <div>Loading...</div>;
  }

  return (
    <Container>
      {data.folders.map((folder) => (
        <Folder key={folder.id} folder={folder} />
      ))}
    </Container>
  );
};

export default SDContent;
