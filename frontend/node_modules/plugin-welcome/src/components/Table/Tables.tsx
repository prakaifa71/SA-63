import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import { DefaultApi } from '../../api/apis';
import { EntPatientofphysician } from '../../api/models/EntPatientofphysician';
import { AppBar, IconButton, Menu, Toolbar, Typography } from '@material-ui/core';
import {
  MenuItem, 
  Button, 
} from '@material-ui/core';
import AccountCircle from '@material-ui/icons/AccountCircle';

 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [patientofphysicians, setPatientofphysicians] = useState<EntPatientofphysician[]>();
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getPatientofphysicians = async () => {
     const res = await api.listPatientofphysician({ limit: 10, offset: 0 });
     setLoading(false);
     setPatientofphysicians(res);
   };
   getPatientofphysicians();
 }, [loading]);
 
 const deletePatientofphysicians = async (id: number) => {
   const res = await api.deletePatientofphysician({ id: id });
   setLoading(true);
 };
 
 return (
 
   
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Physician</TableCell>
           <TableCell align="center">Room</TableCell>
           <TableCell align="center">Patient</TableCell>
           <TableCell align="center">Ailment</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
       {patientofphysicians === undefined 
          ? null
          : patientofphysicians.map((item :any)=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.idphysicianid?.physicianname}</TableCell>
             <TableCell align="center">{item.edges?.roomid?.typeroom}</TableCell>
             <TableCell align="center">{item.edges?.idpatient?.patientname}</TableCell>
             <TableCell align="center">{item.ailment}</TableCell>
             
             
             

             <TableCell align="center">
               <Button
                 onClick={() => {
                   deletePatientofphysicians(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
             
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
