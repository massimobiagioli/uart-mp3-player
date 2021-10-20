import { useEffect, useState } from "react";
import { readSDContent } from "../api/Backend";

const SDContent = () => {
  const [data, setData] = useState()

  useEffect(() => {
    readSDContent()
      .then(data => setData(data))
      .catch(err => console.error(err.message));
  }, [])  
  
  return (
    <div>
      {data && data.folders.map(folder => 
        <div>
          {folder.id}
        </div>
      )}
    </div>
  );
};
  
export default SDContent;