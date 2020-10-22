
import {
  Grid,
  MenuItem,
  TextField,
  Button,
  FormControl,
  InputLabel,
  Select,
} from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Menu from '@material-ui/core/Menu';
import { DefaultApi } from '../../api/apis';
import { Link as RouterLink } from 'react-router-dom';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPhysician } from '../../api/models/EntPhysician';
import { EntPatientroom } from '../../api/models/EntPatientroom';
import { Alert } from '@material-ui/lab';






const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);

export default function MenuAppBar() {
  const classes = useStyles();
  const [auth, setAuth] = React.useState(true);

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);


  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const api = new DefaultApi();
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [physician, setPhysicians] = useState<EntPhysician[]>([]);
  const [patientrooms, setPatientrooms] = useState<EntPatientroom[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [physicianid, setPhysicianid] = useState(Number);
  const [ailment, setAilment] = useState(String);
  const [patientid, setPatientid] = useState(Number);
  const [roomid, setRoomid] = useState(Number);

  useEffect(() => {
    const getPatients = async () => {
      const p = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(p);
    };
    getPatients();


    const getPhysicians = async () => {
      const ph = await api.listPhysician({ limit: 10, offset: 0 });
      setLoading(false);
      setPhysicians(ph);
    };
    getPhysicians();

    const getPatientrooms = async () => {
      const pa = await api.listPatientroom({ limit: 10, offset: 0 });
      setLoading(false);
      setPatientrooms(pa);
    };
    getPatientrooms();

  }, [loading]);

  const handlesymptomChange = (event: any) => {
    setAilment(event.target.value as string);
  };

  const CreatePatientofphysician = async () => {
    const patientofphysician = {
      ailment: ailment,
      patientID: patientid,
      roomID: roomid,
      physicianID: physicianid,



    }
    console.log(patientofphysician);

    const res: any = await api.createPatientofphysician({ patientofphysician: patientofphysician });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const patient_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientid(event.target.value as number);
  };
  const physician_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPhysicianid(event.target.value as number);
  };

  const room_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoomid(event.target.value as number);
  };

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h5" className={classes.title}>
            ระบบผู้ป่วยในการดูแลของแพทย์
          </Typography>
          {auth && (
            <div>
              <IconButton
                aria-label="account of current patientofphysician"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={open}
                onClose={handleClose}
              >
                <Button className={classes.button} variant="outlined" color="inherit"
                  size="small" component={RouterLink}
                  to="/">
                  logout
              </Button>

              </Menu>


            </div>
          )}
        </Toolbar>
      </AppBar>

      <AppBar position="static" color='inherit'>
        <Grid container alignItems="center" spacing={4}>
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>

          <Grid item xs={6}>   <center>
            <h2 align='right'>แพทย์ (Physician)</h2></center>
          </Grid>
          <Grid item xs={6}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>แพทย์</InputLabel>
              <Select
                labelId="physiciant_id-label"
                label="Physician"
                id="physician_id"
                value={physicianid}
                onChange={physician_id_handleChange}
                style={{ width: 300 }}
              >
                {physician.map((item: EntPhysician) =>
                  <MenuItem value={item.id}>{item.physicianname}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}> <center>
            <h2 align='right'>ห้องผู้ป่วย (Patientroom)</h2></center>

          </Grid>
          <Grid item xs={6}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>ห้องผู้ป่วย</InputLabel>
              <Select
                labelId="roomt_id-label"
                label="Patientroom"
                id="room_id"
                value={roomid}
                onChange={room_id_handleChange}
                style={{ width: 300 }}
              >
                {patientrooms.map((item: EntPatientroom) =>
                  <MenuItem value={item.id}>{item.typeroom}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>



          <Grid item xs={6}><center>
            <h2 align='right'>ผู้ป่วย (Patient)</h2></center>
          </Grid>
          <Grid item xs={6}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"
            ><InputLabel>ผู้ป่วย</InputLabel>
              <Select
                labelId="patient_id-label"
                label="patient"
                id="patient_id"
                value={patientid}
                onChange={patient_id_handleChange}
                style={{ width: 300 }}
              >
                {patients.map((item: EntPatient) =>
                  <MenuItem value={item.id}>{item.patientname}</MenuItem>)}
              </Select>
            </FormControl>
          </Grid>



          <Grid item xs={6}><center>
            <h2 align='right'>อาการโรค (Ailment)</h2></center>
          </Grid>
          <Grid item xs={6}>
            <FormControl
              fullWidth
              className={classes.margin}
              variant="outlined"

            ><TextField
                id="ailment"
                label="อาการโรค"
                variant="outlined"
                type="string"
                size="medium"
                value={ailment}
                onChange={handlesymptomChange}
                style={{ width: 300 }}
              />
            </FormControl>
          </Grid>







          <Grid item xs={4}>
          </Grid>
          <Grid item xs={4}>   <center>

            <Button
              onClick={() => {
                CreatePatientofphysician();
              }}
              variant="contained"
              color="primary"
            >
              Save
             </Button>
            <Button
              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/login"
              variant="contained"
            >
              Show
             </Button>

             
                  {status ? (
                    <div>
                      {alert ? (
                        <Alert severity="success">
                          This is a success alert — check it out!
                        </Alert>
                      ) : (
                          <Alert severity="warning" style={{ marginTop: 20 }}>
                            This is a warning alert — check it out!
                          </Alert>
                        )}
                    </div>
                  ) : null}
               
          </center>
          </Grid>
          <Grid item xs={4}></Grid>
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>
        </Grid>
      </AppBar>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
          </Typography>
        </Toolbar>
      </AppBar>
    </div>
  );
}
