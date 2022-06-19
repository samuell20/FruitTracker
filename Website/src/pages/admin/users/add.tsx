import type { NextPage } from 'next'
import Alert from '../../../Common/Alert'
import {AddEdit} from 'Common/Users/AddEdit'

const Add: NextPage = () => {
  return (
     <>
       <h1>Usuarios</h1>
        <Alert />
        <AddEdit />
     </> 
  )
}

export default Add 