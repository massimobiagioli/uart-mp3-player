import { Grid, Box, Button, Stack } from "@mui/material";
import { PlayCircle, StopCircle } from "@mui/icons-material";
import { play, stop } from "../api/Backend";

const handlePlay = (folderId, songId) => async () => {
  try {
    await play(folderId, songId);
  } catch (error) {
    console.log(error.message);
  }
};

const handleStop = (folderId, songId) => async () => {
  try {
    await stop(folderId, songId);
  } catch (error) {
    console.log(error.message);
  }
};

const Song = ({ folderId, song }) => {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <Grid container spacing={2} marginBottom={2}>
        <Grid item xs={1}>
          <div>
            <strong>{song.id}</strong>
          </div>
        </Grid>
        <Grid item xs={7}>
          <div>
            {song.name} - {song.author}
          </div>
          <div>
            <small>{song.filename}</small>
          </div>
        </Grid>
        <Grid item xs={4}>
          <Stack spacing={2} direction="row">
            <Button
              variant="outlined"
              onClick={handlePlay(folderId, song.id)}
              startIcon={<PlayCircle />}
            >
              Play
            </Button>
            <Button
              variant="outlined"
              onClick={handleStop(folderId, song.id)}
              startIcon={<StopCircle />}
            >
              Stop
            </Button>
          </Stack>
        </Grid>
      </Grid>
    </Box>
  );
};

export default Song;
