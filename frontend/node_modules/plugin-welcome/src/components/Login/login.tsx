import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import {
 Content,

 Page,
 pageTheme,
 ContentHeader, 
 Link,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {
 
 return (
   
   <Page theme={pageTheme.home}>
    
     <Content>
       <ContentHeader title="ข้อมูลผู้ป่วยในการดูแลของแพทย์">
       
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add Data           
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
export default WelcomePage;