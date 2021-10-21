import { Card, CardHeader, CardContent } from "@mui/material";
import Song from "./Song";

const Folder = ({ folder }) => {
  const title = `${folder.id} - ${folder.name} - ${folder.author} (${folder.year})`;
  return (
    <Card>
      <CardHeader title={title} />
      <CardContent>
        {folder.songs.map((song) => (
          <Song key={song.id} song={song} folderId={folder.id} />
        ))}
      </CardContent>
    </Card>
  );
};

export default Folder;
