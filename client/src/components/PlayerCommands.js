import { Button, Stack } from "@mui/material";
import { reset } from "../api/Backend";
import { RotateLeft } from "@mui/icons-material";

const handleReset = async () => {
  try {
    await reset();
  } catch (error) {
    console.log(error.message);
  }
};

const PlayerCommands = () => {
  return (
    <Stack spacing={2} margin={2} direction="row">
      <Button
        variant="outlined"
        color="error"
        onClick={handleReset}
        startIcon={<RotateLeft />}
      >
        Reset
      </Button>
    </Stack>
  );
};

export default PlayerCommands;
