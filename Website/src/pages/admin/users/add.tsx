import type { NextPage } from 'next'
import Alert from '../../../Common/Alert'
import {AddEdit} from 'Common/Users/AddEdit'

const Add: NextPage = (params) => {
  return (
     <>
        <Alert />
        <AddEdit params={params} />
     </> 
  )
}

export default Add 