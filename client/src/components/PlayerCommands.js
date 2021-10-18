import Button from '@mui/material/Button';
import { makeStyles } from '@material-ui/core/styles';
import { reset } from '../api/Backend'

const useStyles = makeStyles(theme => ({
    playerCommadsContainer: {
      marginTop: '5px',
      marginLeft: '5px'
    },
}));

const handleReset = async () => {
    try {
      await reset();
    } catch (error) {
      console.log(error.message);
    }
}

const PlayerCommands = () => {
    const classes = useStyles();

    return (
      <div className={classes.playerCommadsContainer}>
        <Button 
            variant="outlined" 
            color="error"
            onClick={handleReset}
        >
            Reset
        </Button>
      </div>
    );
  };
  
  export default PlayerCommands;