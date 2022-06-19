import type { NextPage } from 'next'
import Alert from '../../../../Common/Alert'
import {AddEdit} from 'Common/Users/AddEdit'
import { userService } from 'Services'

const Edit: NextPage = (params) => {
  return (
     <>
       <h1>Usuarios</h1>
        <Alert />
        <AddEdit params={params}/>
     </> 
  )
}

export async function getServerSideProps({ params }:any) {
    //const user = await userService.getById(params.id);

    return {
        props: {
          user:{
            id:1
          }
        }
    }
}

export default Edit 